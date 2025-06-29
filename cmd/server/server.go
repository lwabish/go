package server

import (
	"github.com/spf13/cobra"
)

// Cmd represents the server command
var Cmd = &cobra.Command{
	Use:   "server",
	Short: "start a http server",
}

func init() {
	Cmd.AddCommand(rawCmd)
	Cmd.AddCommand(homeCmd)
}
