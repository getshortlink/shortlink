package cmd

import (
	"github.com/getshortlink/shortlink/context"
	"github.com/getshortlink/shortlink/server"
	"github.com/spf13/cobra"
)

type ServerCommand struct {
	*cobra.Command
}

func NewServerCommand() *ServerCommand {
	c := &ServerCommand{
		Command: &cobra.Command{
			Use:   "server",
			Short: "Run the HTTP server",
			Long:  "Run the HTTP server",
		},
	}

	c.Command.RunE = c.Run

	return c
}

func (c *ServerCommand) Run(cmd *cobra.Command, args []string) error {
	server, err := server.NewServer()
	if err != nil {
		return err
	}

	_ = context.WithSigtermCancel(cmd.Context(), func() {
		server.Stop()
	})

	return server.Start()
}
