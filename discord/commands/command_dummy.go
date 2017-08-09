package cmd

import "github.com/bwmarrin/discordgo"

// CommandDummy represents a simple response message
type CommandDummy struct {
	Cooldown int
	Common   CommandCommon
}

// Execute acts upon a given message
func (c CommandDummy) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	send := insertPlaceholders(c.Common.Response, m)

	s.ChannelMessage(m.ChannelID, send)
}
