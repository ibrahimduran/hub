package cmd

import (
	"fmt"

	"github.com/ibrahimduran/hub/internal/auth"
	"github.com/ibrahimduran/hub/internal/hub"
	"github.com/ibrahimduran/hub/internal/protocol"
	local "github.com/ibrahimduran/hub/internal/protocol/local"
	"github.com/spf13/cobra"
)

var hubCmd = &cobra.Command{
	Use: "hub",
	Run: func(cmd *cobra.Command, args []string) {
		server := hub.NewServer()

		localProvider := &local.LocalProvider{}
		server.Use(localProvider)

		localProvider.Register(&local.Help{})
		localProvider.Register(&local.Info{})
		localProvider.Register(&protocol.ModeCommand{})
		localProvider.Register(&auth.RegisterCommand{})

		fmt.Printf("Listening on :%d\n", 3000)
		server.Listen(3000)
	},
}

func init() {
	rootCmd.AddCommand(hubCmd)
}
