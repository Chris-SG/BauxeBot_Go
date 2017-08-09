package cmd

import "github.com/bwmarrin/discordgo"

// CommandColor represents a struct for changing name colors
type CommandColor struct {
	Cooldown int
	Common   CommandCommon
}

// Execute reporesents acting upon the color command
func (c CommandColor) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	send := insertPlaceholders(c.Common.Response, m)

	s.ChannelMessage(m.ChannelID, send)
}
