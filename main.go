package main

import (
	"github.com/gochrono/clocktower/commands"
	"os"
)

func main() {
	resp := commands.Execute(os.Args[1:])
	if resp.Err != nil {
		if resp.IsUserError() {
			resp.Cmd.Println("")
			resp.Cmd.Println(resp.Cmd.UsageString())
		}
		os.Exit(-1)
	}
}
