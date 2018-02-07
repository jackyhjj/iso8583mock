package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/nkristianto/tcp_server/commands"
	"github.com/nkristianto/tcp_server/infrastructures"
)

var text = "Starting Application"

func main() {
	// show text //
	fmt.Println(fmt.Sprintf(text))

	command := commands.NewCommandEngine()
	cfg, err := command.GetRoot().PersistentFlags().GetString("config")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	infrastructures.SetConfig(cfg)
	command.Run()
	//dd
}
