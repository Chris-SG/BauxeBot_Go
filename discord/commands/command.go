package cmd

import "github.com/bwmarrin/discordgo"

// Command works as an interface for other commands
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
}

func init() {

}
