package commands

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"

	"github.com/nkristianto/tcp_server/commands/action"
)

type CommandEngine struct {
	rootCmd *cobra.Command
}

// NewCommandEngine the command line boot loader
func NewCommandEngine() *CommandEngine {
	var rootCmd = &cobra.Command{
		Use:   "Tcp Mock Server",
		Short: "Tcp Mock command line",
		Long:  "Tcp Mock command line",
	}

	defer func() {
		r := recover()
		if r != nil {
			logrus.Error(r)
		}
	}()

	rootCmd.PersistentFlags().StringP("config", "c", "configurations", "the config path location")

	return &CommandEngine{
		rootCmd:rootCmd,
	}
}

func(c *CommandEngine) GetRoot() *cobra.Command{
	return c.rootCmd
}

func(c *CommandEngine) Run(){
	var commands = []*cobra.Command{
		{
			Use: "server",
			Short:"ISO 8583 Mock Server",
			Long:"ISO 8583 Mock Server",
			Run:func(cmd *cobra.Command, args []string){
				action.RunTcpServer()
			},
		},
		{
			Use: "client",
			Short:"ISO 8583 Mock client",
			Long:"ISO 8583 Mock client",
			Run:func(cmd *cobra.Command, args []string){
				action.RunTcpClient()
			},
		},
	}
	for _, command := range commands{
		c.rootCmd.AddCommand(command)
	}

	c.rootCmd.Execute()
}


