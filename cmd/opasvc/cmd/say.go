// Copyright © 2018 SMB Lab AS <torkel@smblab.no>

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/tlyng/opasvc/pb"
)

// sayCmd represents the say command
var sayCmd = &cobra.Command{
	Use:   "say <greeting> <name>",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Usage()
			os.Exit(1)
		}
		dialAddr := fmt.Sprintf("passthrough://localhost/%s", listenAddress)
		ctx := context.Background()
		conn, err := grpc.DialContext(
			ctx,
			dialAddr,
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)
		if err != nil {
			log.Fatalln("Failed to dial server:", err)
		}
		defer conn.Close()

		hello := pb.NewHelloClient(conn)

		greeting := args[0]
		name := args[1]
		resp, err := hello.Say(ctx, &pb.Request{
			Greeting: greeting,
			Name:     name,
		})
		if err != nil {
			log.Fatalln("Failed to Say:", err)
		}
		log.Infoln("Response:", resp)
	},
}

func init() {
	rootCmd.AddCommand(sayCmd)
}
