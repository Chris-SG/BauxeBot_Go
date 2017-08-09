package cmd

import "github.com/bwmarrin/discordgo"

// Command works as an interface for other commands
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
}

// CommandCommon represents fields common between all commands
type CommandCommon struct {
	Caller      string
	Response    string
	Description string
}

func init() {

}
