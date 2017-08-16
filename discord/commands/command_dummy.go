package cmd

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// CommandDummy represents a simple response message
type CommandDummy struct {
	Cooldown int
	Common   CommandCommon
}

// Execute acts upon a given message
func (c CommandDummy) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	send := insertPlaceholders(c.Common.Response, m)
	log.Print("trying to send")

	s.ChannelMessageSend(m.ChannelID, send)
}

func (c CommandDummy) GetCommons() CommandCommon {
	return c.Common
}
