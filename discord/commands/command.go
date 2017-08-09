package cmd

import (
	"container/list"

	"github.com/bwmarrin/discordgo"
)

// Command works as an interface for other commands
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
}

// CommandCommon represents fields common between all commands
type CommandCommon struct {
	Caller              string
	Response            string
	Description         string
	Channels            *list.List
	RequiredPermissions int
	RequiredUsers       *list.List
}

const (
	f0 = iota
	f1 = iota
	f2 = iota
	//etc
)

func canExecute(s *discordgo.Session, m *discordgo.MessageCreate) (canExecute bool) {
	userPerms, _ := s.UserChannelPermissions(m.Author.ID, m.ChannelID)
	for i := 0; i < 32; i++ {
		if RequiredPermissions&(1<<i) != 0 {
			if userPerms & (0 << i) {
				return false

			}
		}

	}

	return true
}
