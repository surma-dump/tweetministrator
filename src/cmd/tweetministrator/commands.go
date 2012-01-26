package main

import (
	"log"
	"os/exec"
)

func ExecuteCommand(cmdname string) {
	cmd, ok := config.Commands[cmdname]
	if !ok {
		log.Printf("Invalid command: %s", cmdname)
		return
	}
	log.Printf("Executing %s...", cmdname)
	e := exec.Command(cmd[0], cmd[1:]...).Start()
	if e != nil {
		log.Printf("Could not execute command %s: %s", cmdname, e)
	}
}
