// Copyright © 2018 SMB Lab AS <torkel@smblab.no>

package cmd

import (
	"net"
	"os"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"github.com/tlyng/opasvc/interceptor"
	"github.com/tlyng/opasvc/pb"
	"github.com/tlyng/opasvc/server"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	listenAddress string
	log           grpclog.LoggerV2
)

func init() {
	log = grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(log)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve HelloService",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", listenAddress)
		if err != nil {
			log.Fatalln("Failed to listen:", err)
		}

		s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.PolicyInterceptor()))
		pb.RegisterHelloServer(s, server.New())

		reflection.Register(s)
		log.Info("Serving gRPC on http://", listenAddress)
		log.Fatal(s.Serve(lis))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.PersistentFlags().StringVar(&listenAddress, "listen-address", ":50052", "Listen address for gRPC server")
}
