// Code generated by protoc-gen-cobra. DO NOT EDIT.

package p2pv1alpha1

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
	io "io"
)

func HostServiceClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("HostService"),
		Short: "HostService service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_HostServiceHostCommand(cfg),
		_HostServicePeersCommand(cfg),
		_HostServiceConnectCommand(cfg),
		_HostServiceSendCommand(cfg),
		_HostServicePublishCommand(cfg),
		_HostServiceSubscribeCommand(cfg),
	)
	return cmd
}

func _HostServiceHostCommand(cfg *client.Config) *cobra.Command {
	req := &HostRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Host"),
		Short: "Host RPC client",
		Long:  "Host returns the information about the host node.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService", "Host"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewHostServiceClient(cc)
				v := &HostRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Host(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	return cmd
}

func _HostServicePeersCommand(cfg *client.Config) *cobra.Command {
	req := &PeersRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Peers"),
		Short: "Peers RPC client",
		Long:  "Peers lists information about connected peer nodes.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService", "Peers"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewHostServiceClient(cc)
				v := &PeersRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Peers(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	return cmd
}

func _HostServiceConnectCommand(cfg *client.Config) *cobra.Command {
	req := &ConnectRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Connect"),
		Short: "Connect RPC client",
		Long:  "Connect connects to a peer node.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService", "Connect"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewHostServiceClient(cc)
				v := &ConnectRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Connect(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.PeerInfo, cfg.FlagNamer("PeerInfo"), "", "")

	return cmd
}

func _HostServiceSendCommand(cfg *client.Config) *cobra.Command {
	req := &SendRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Send"),
		Short: "Send RPC client",
		Long:  "Send sends a message to a peer node.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService", "Send"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewHostServiceClient(cc)
				v := &SendRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Send(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.PeerInfo, cfg.FlagNamer("PeerInfo"), "", "")
	cmd.PersistentFlags().StringVar(&req.Protocol, cfg.FlagNamer("Protocol"), "", "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Data, cfg.FlagNamer("Data"), "")

	return cmd
}

func _HostServicePublishCommand(cfg *client.Config) *cobra.Command {
	req := &PublishRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Publish"),
		Short: "Publish RPC client",
		Long:  "Publish broadcasts a message to a topic.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService", "Publish"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewHostServiceClient(cc)
				v := &PublishRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Publish(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Topic, cfg.FlagNamer("Topic"), "", "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Data, cfg.FlagNamer("Data"), "")

	return cmd
}

func _HostServiceSubscribeCommand(cfg *client.Config) *cobra.Command {
	req := &SubscribeRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Subscribe"),
		Short: "Subscribe RPC client",
		Long:  "Subscribe broadcasts a message to a topic.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "HostService", "Subscribe"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewHostServiceClient(cc)
				v := &SubscribeRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				stm, err := cli.Subscribe(cmd.Context(), v)

				if err != nil {
					return err
				}

				for {
					res, err := stm.Recv()
					if err != nil {
						if err == io.EOF {
							break
						}
						return err
					}
					if err = out(res); err != nil {
						return err
					}
				}
				return nil

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Topic, cfg.FlagNamer("Topic"), "", "")

	return cmd
}