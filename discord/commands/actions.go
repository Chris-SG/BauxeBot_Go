package cmd

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

// This file will be used for actions that may be required by all commands
// This will include:
// Delete caller message
// Delete respond message after time

// DeleteMessage a
func DeleteMessage() {

}

func orderRolesByPositon(r discordgo.Roles) {
	for i := 0; i < len(r)-1; i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i].Position > r[j].Position {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
}

func findCommandByName(cmds Commands, name string) (c Command, err error) {
	for _, colorCommand := range cmds.ColorCommands {
		if colorCommand.Common.Caller == name {
			c = colorCommand
			return
		}
	}
	for _, dummyCommand := range cmds.DummyCommands {
		if dummyCommand.Common.Caller == name {
			c = dummyCommand
			return
		}
	}
	for _, debugCommand := range cmds.DebugCommands {
		if debugCommand.Common.Caller == name {
			c = debugCommand
			return
		}
	}
	for _, modCommand := range cmds.ModerationCommands {
		if modCommand.Common.Caller == name {
			c = modCommand
			return
		}
	}
	err = errors.New("Cannot find command")
	return
}
