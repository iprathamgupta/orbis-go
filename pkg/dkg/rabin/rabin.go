package rabin

import (
	"context"
	"errors"
	"fmt"
	"sync"

	logging "github.com/ipfs/go-log"
	"go.dedis.ch/kyber/v3"
	rabindkg "go.dedis.ch/kyber/v3/share/dkg/rabin"
	"go.dedis.ch/kyber/v3/suites"
	"google.golang.org/protobuf/proto"

	rabinv1alpha1 "github.com/sourcenetwork/orbis-go/gen/proto/orbis/rabin/v1alpha1"
	"github.com/sourcenetwork/orbis-go/pkg/bulletin"
	"github.com/sourcenetwork/orbis-go/pkg/crypto"
	"github.com/sourcenetwork/orbis-go/pkg/db"
	orbisdkg "github.com/sourcenetwork/orbis-go/pkg/dkg"
	"github.com/sourcenetwork/orbis-go/pkg/transport"
	"github.com/sourcenetwork/orbis-go/pkg/types"
)

var log = logging.Logger("orbis/dkg/rabin")

const name = "rabin"

type dkg struct {
	mu sync.Mutex

	ringID types.RingID

	rdkg         *rabindkg.DistKeyGenerator
	participants []orbisdkg.Node

	index     int
	num       int32
	threshold int32

	suite   suites.Suite
	pubKey  kyber.Point
	privKey kyber.Scalar

	dealRepo          db.Repository[*rabinv1alpha1.Deal]
	respRepo          db.Repository[*rabinv1alpha1.Response]
	secretCommitsRepo db.Repository[*rabinv1alpha1.SecretCommits]
	dkgRepo           db.Repository[*rabinv1alpha1.DKG]

	deals     chan dealDispatch
	responses chan responseDispatch
	commits   chan secretCommitsDispatch

	share crypto.PriShare

	db        *db.DB
	transport transport.Transport
	bulletin  bulletin.Bulletin

	state orbisdkg.State
}

func New(repo *db.DB, rkeys []db.RepoKey, t transport.Transport, b bulletin.Bulletin) (*dkg, error) {
	if len(rkeys) != 4 {
		return nil, ErrMissingRepoKeys
	}
	dealsRepo, err := db.GetRepo(repo, rkeys[0], dealPkFunc)
	if err != nil {
		return nil, errors.Join(ErrCouldntGetRepo, err)
	}
	respsRepo, err := db.GetRepo(repo, rkeys[1], responsePkFunc)
	if err != nil {
		return nil, errors.Join(ErrCouldntGetRepo, err)
	}
	secretCommitsRepo, err := db.GetRepo(repo, rkeys[2], secretCommitsPkFunc)
	if err != nil {
		return nil, errors.Join(ErrCouldntGetRepo, err)
	}
	dkgRepo, err := db.GetRepo(repo, rkeys[3], dkgPkFunc)
	if err != nil {
		return nil, errors.Join(ErrCouldntGetRepo, err)
	}

	return &dkg{
		db:                repo,
		dealRepo:          dealsRepo,
		respRepo:          respsRepo,
		secretCommitsRepo: secretCommitsRepo,
		dkgRepo:           dkgRepo,
		transport:         t,
		bulletin:          b,
		index:             -1,
	}, nil
}

// Init initializes the DKG with the target nodes
func (d *dkg) Init(ctx context.Context, pk crypto.PrivateKey, rid types.RingID, nodes []orbisdkg.Node, n int32, threshold int32, fromState bool) error {
	// try load from persisted state
	// otherwise initalize from new
	if fromState {
		return d.initFromState(ctx, pk, rid, nodes, n, threshold)
	}
	return d.initFromNew(ctx, pk, rid, nodes, n, threshold)
}

func (d *dkg) initFromState(ctx context.Context, pk crypto.PrivateKey, rid types.RingID, nodes []orbisdkg.Node, n int32, threshold int32) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.ringID = rid
	err := d.loadUnsafe(ctx)
	if err != nil {
		return err
	}

	// ensure the loaded state hasn't diverged from the initialized state
	// Would imply that either the caller has made a mistake, or the DB
	// has been currupted, either way, bad!
	// TODO!!
	// if !d.Equal(&dkg{...}) { ... }

	return d.initCommon(ctx)
}

func (d *dkg) initFromNew(ctx context.Context, pk crypto.PrivateKey, rid types.RingID, nodes []orbisdkg.Node, n int32, threshold int32) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if pk == nil {
		return fmt.Errorf("missing private key")
	}

	suite, err := crypto.SuiteForType(pk.Type())
	if err != nil {
		return fmt.Errorf("get suite for type: %w", err)
	}

	d.ringID = rid
	d.suite = suite
	d.privKey = pk.Scalar()
	d.pubKey = suite.Point().Mul(d.privKey, nil) // public point for scalar
	d.num = n
	d.threshold = threshold

	if len(nodes) != int(n) {
		return orbisdkg.ErrBadNodeSet
	}

	d.participants = nodes

	points := make([]kyber.Point, 0, len(d.participants))
	for i, n := range d.participants {
		point := n.PublicKey().Point()
		if point.Equal(d.pubKey) {
			d.index = i
		}
		points = append(points, point)
	}

	// we didn't find ourselves in the list
	if d.index == -1 {
		return orbisdkg.ErrMissingSelf
	}

	rdkg, err := rabindkg.NewDistKeyGenerator(d.suite, d.privKey, points, int(d.threshold))
	if err != nil {
		return fmt.Errorf("create DKG: %w", err)
	}

	d.rdkg = rdkg
	d.state = orbisdkg.INITIALIZED

	err = d.initCommon(ctx)
	if err != nil {
		return err
	}
	return d.save(ctx) // save the initialized state
}

// initCommon does all the none state initialization. Shared
// between initFromNew() and initFromState()
func (d *dkg) initCommon(ctx context.Context) error {
	// setup stream handler for transport
	d.setupHandlers()

	d.deals = make(chan dealDispatch, d.numExpectedDeals())
	d.responses = make(chan responseDispatch, d.numExpectedResponses())
	d.commits = make(chan secretCommitsDispatch, d.numExpectedCommits())

	return nil
}

func (d *dkg) Name() string {
	return name
}

func (d *dkg) PublicKey() crypto.PublicKey {
	pk, _ := crypto.PublicKeyFromPoint(d.pubKey)
	return pk
}

func (d *dkg) Share() crypto.PriShare {
	return d.share
}

func (d *dkg) State() orbisdkg.State {
	return d.state
}

// Start the DKG setup process.
func (d *dkg) Start(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.state == orbisdkg.INITIALIZED {
		log.Debug("Generating and persisting deals")
		deals, err := d.rdkg.Deals()
		if err != nil {
			return fmt.Errorf("generate deals: %w", err)
		}

		// persist deals
		for _, deal := range deals {

			dealproto, err := d.dealToProto(deal)
			if err != nil {
				return fmt.Errorf("convert deal to proto: %w", err)
			}

			dealproto.RingId = string(d.ringID)
			dealproto.NodeId = d.transport.Host().ID()
			log.Debugf("Creating doc: %+v", deal)
			err = d.dealRepo.Create(ctx, dealproto)
			if err != nil {
				return fmt.Errorf("create deal: %w", err)
			}
		}

		d.state = orbisdkg.STARTED
		err = d.save(ctx)
		if err != nil {
			return fmt.Errorf("save dkg state as started")
		}
	}

	if d.state == orbisdkg.STARTED {

		// send deals
		dealProtos, err := d.dealRepo.GetAll(ctx) // todo: add filter
		if err != nil {
			return fmt.Errorf("getting deals to broadcast: %w", err)
		}

		for _, deal := range dealProtos {

			// skip deals that don't match the nodeID and ringID of this DKG
			// This can be removed once proper filtering is working on the repo
			if deal.NodeId != d.transport.Host().ID() &&
				deal.RingId != string(d.ringID) {
				continue
			}

			buf, err := proto.Marshal(deal)
			if err != nil {
				return fmt.Errorf("marshal deal: %w", err)
			}

			log.Infof("Node %d sending deal to participant %d", d.index, deal.TargetIndex)
			err = d.send(ctx, string(ProtocolDeal), buf, d.participants[deal.TargetIndex])
			if err != nil {
				return fmt.Errorf("send deal: %w", err)
			}
		}
		d.state = RECIEVING
		err = d.save(ctx)
		if err != nil {
			return fmt.Errorf("save dkg state as recieving")
		}
	}

	go d.dispatch()

	return nil
}

func (d *dkg) send(ctx context.Context, msgType string, buf []byte, node transport.Node) error {

	cid, err := types.CidFromBytes(buf)
	if err != nil {
		return fmt.Errorf("cid from bytes: %w", err)
	}

	msg, err := d.transport.NewMessage(d.ringID, cid.String(), false, buf, msgType)
	if err != nil {
		return fmt.Errorf("new message: %w", err)
	}

	log.Debugf("send to: %q, addr: %q", node.ID(), node.Address())
	err = d.transport.Send(ctx, node, msg)
	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}

func (d *dkg) Close(_ context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (d *dkg) ProcessMessage(msg *transport.Message) error {
	// TODO maybe?: validate msg.PublicKey matches payload pubkeys

	log.Debugf("process message: id: %q, type: %q", msg.Id, msg.GetType())

	switch msg.GetType() {
	case string(ProtocolDeal):

		var protoDeal rabinv1alpha1.Deal

		err := proto.Unmarshal(msg.Payload, &protoDeal)
		if err != nil {
			return fmt.Errorf("unmarshal deal message: %w", err)
		}
		protoDeal.RingId = string(d.ringID)
		protoDeal.NodeId = msg.NodeId
		err = d.dealRepo.Create(context.TODO(), &protoDeal)
		if err != nil && !errors.Is(err, db.ErrRecordAlreadyExists) {
			// don't need to process if it already exists
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to insert deal into repo: %w", err)
		}

		err = d.dispatchDealProto(&protoDeal)
		if err != nil {
			return fmt.Errorf("dispatch deal message: %w", err)
		}

	case string(ProtocolResponse):

		var protoResponse rabinv1alpha1.Response

		err := proto.Unmarshal(msg.Payload, &protoResponse)
		if err != nil {
			return fmt.Errorf("unmarshal response message: %w", err)
		}

		protoResponse.RingId = string(d.ringID)
		protoResponse.NodeId = msg.NodeId
		err = d.respRepo.Create(context.TODO(), &protoResponse)
		if err != nil && !errors.Is(err, db.ErrRecordAlreadyExists) {
			// don't need to process if it already exists
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to insert deal into repo: %w", err)
		}

		err = d.dispatchResponseProto(&protoResponse)
		if err != nil {
			return fmt.Errorf("dispatch response message: %w", err)
		}

	case string(ProtocolSecretCommits):

		var protoSecretCommits rabinv1alpha1.SecretCommits

		err := proto.Unmarshal(msg.Payload, &protoSecretCommits)
		if err != nil {
			return fmt.Errorf("unmarshal secret commits: %w", err)
		}

		protoSecretCommits.RingId = string(d.ringID)
		protoSecretCommits.NodeId = msg.NodeId
		err = d.secretCommitsRepo.Create(context.TODO(), &protoSecretCommits)
		if err != nil && !errors.Is(err, db.ErrRecordAlreadyExists) {
			// don't need to process if it already exists
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to insert deal into repo: %w", err)
		}

		err = d.dispatchSecretCommitsProto(&protoSecretCommits)
		if err != nil {
			return fmt.Errorf("dispatch secret commits: %w", err)
		}

	default:
		return fmt.Errorf("unknown message type: %q", msg.GetType())
	}

	return nil
}

// dispatch is responsible for handling all the incoming
// messages, and dispatching them to their cooresponding
// handlers, but with ordering via channels. This gurantees
// that we handle all the events at their appropriate
// time.
//
// It is designed to run in a gourinte
func (d *dkg) dispatch() {

	// processDeals
	for i := 0; i < d.numExpectedDeals() && d.state < PROCESSED_DEALS; i++ {
		dd := <-d.deals
		log.Infof("Node %d handling deal for dealer %d (%d/%d)", d.index, dd.deal.Index, i+1, d.numExpectedDeals())
		dd.err <- d.processDeal(dd.deal)
	}

	d.state = PROCESSED_DEALS
	err := d.save(context.TODO())
	if err != nil {
		log.Fatalf("failed to save DKG state: %w", err)
	}

	// processResponses
	for i := 0; i < d.numExpectedResponses() && d.state < PROCESSED_RESPONSES; i++ {
		rd := <-d.responses
		log.Infof("Node %d handling response for dealer %d (%d/%d)", d.index, rd.respone.Index, i+1, d.numExpectedResponses())
		rd.err <- d.processResponse(rd.respone)
	}

	d.state = PROCESSED_RESPONSES
	err = d.save(context.TODO())
	if err != nil {
		log.Fatalf("failed to save DKG state: %w", err)
	}

	// processSecrets
	for i := 0; i < d.numExpectedCommits(); i++ {
		sd := <-d.commits
		log.Infof("Node %d handling secret for dealer %d (%d/%d)", d.index, sd.secretCommits.Index, i+1, d.numExpectedCommits())
		sd.err <- d.processSecretCommits(sd.secretCommits)
	}

	// don't need to update state and save
	// like the above dispatch loops
	// since `processSecreCommits` will do
	// this for us.
}

func (d *dkg) dispatchDealProto(dealproto *rabinv1alpha1.Deal) error {
	deal, err := d.dealFromProto(dealproto)
	if err != nil {
		return fmt.Errorf("deal from proto: %w", err)
	}

	err = d.dispatchDeal(deal)
	if err != nil {
		return fmt.Errorf("process deals: %w", err)
	}

	return nil
}

func (d *dkg) dispatchDeal(deal *rabindkg.Deal) error {
	dealDispatchEvent := dealDispatch{
		err:  make(chan error),
		deal: deal,
	}
	// TODO: proper handle full channel failure
	select {
	case d.deals <- dealDispatchEvent:
		// send
	default:
		log.Warn("can't send on deals dispatch channel")
	}
	return <-dealDispatchEvent.err // recieve
}

func (d *dkg) dispatchResponseProto(respproto *rabinv1alpha1.Response) error {
	resp := d.responseFromProto(respproto)

	err := d.dispatchResponse(resp)
	if err != nil {
		return fmt.Errorf("process secret commits: %w", err)
	}
	return nil
}

func (d *dkg) dispatchResponse(resp *rabindkg.Response) error {
	respDispatchEvent := responseDispatch{
		err:     make(chan error),
		respone: resp,
	}
	// TODO: proper handle full channel failure
	select {
	case d.responses <- respDispatchEvent:
		// send
	default:
		log.Warn("cant send on response dispatch channel")
	}
	return <-respDispatchEvent.err
}

func (d *dkg) dispatchSecretCommitsProto(scproto *rabinv1alpha1.SecretCommits) error {
	sc, err := secretCommitsFromProto(d.suite, scproto)
	if err != nil {
		return fmt.Errorf("secret commits from proto: %w", err)
	}

	err = d.dispatchSecretCommit(sc)
	if err != nil {
		return fmt.Errorf("process secret commits: %w", err)
	}
	return nil
}

func (d *dkg) dispatchSecretCommit(sc *rabindkg.SecretCommits) error {
	scDispatchEvent := secretCommitsDispatch{
		err:           make(chan error),
		secretCommits: sc,
	}
	// TODO: proper handle full channel failure
	select {
	case d.commits <- scDispatchEvent:
		// send
	default:
		log.Warn("can't send on commits dispatch channel")
	}
	return <-scDispatchEvent.err // recieve
}

func (d *dkg) numExpectedDeals() int {
	return len(d.participants) - 1
}

func (d *dkg) numExpectedResponses() int {
	l := len(d.participants)
	return (l - 1) * (l - 1)
}

func (d *dkg) numExpectedCommits() int {
	l := len(d.participants)
	return (l - 1) * (l - 1)
}

// save will persist the current DKG state to the DKG Repo.
// It only saves state from the DKG struct, and not the dynamic
// deals, responses, and secret commmits from the internal
// rabin implementation. Those are saved independantly.
func (d *dkg) save(ctx context.Context) error {
	dkgp, err := dkgToProto(d)
	if err != nil {
		return fmt.Errorf("proto conversion: %w", err)
	}

	err = d.dkgRepo.Save(ctx, dkgp)
	if err != nil {
		return fmt.Errorf("saving dkg: %w", err)
	}
	return nil
}

// load will get the persisted DKG state from the DKG Repo.
// It directly writes the result into the pointer, in place.
func (d *dkg) load(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.loadUnsafe(ctx)
}

// loadUnsafe is the same as load, but without locks.
// It requires the caller to aquire a lock
func (d *dkg) loadUnsafe(ctx context.Context) error {
	dkgp, err := dkgToProto(d)
	if err != nil {
		return fmt.Errorf("dkg to proto: %w", err)
	}
	dkgp, err = d.dkgRepo.Get(ctx, dkgp)
	if err != nil {
		return fmt.Errorf("dkg from repo: %w", err)
	}

	_d, err := dkgFromProto(dkgp)
	if err != nil {
		return fmt.Errorf("dkg from proto: %w", err)
	}

	// inplace mutation of state defined on pointer reciever
	d.index = _d.index
	d.num = _d.num
	d.threshold = _d.threshold
	d.suite = _d.suite
	d.state = _d.state
	d.pubKey = _d.pubKey
	d.share = _d.share
	d.participants = _d.participants
	return nil
}
