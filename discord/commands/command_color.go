package cmd

import "github.com/bwmarrin/discordgo"

// CommandColor represents a struct for changing name colors
type CommandColor struct {
	Caller   string
	Response string
	Cooldown int
}

// Execute reporesents acting upon the color command
func (c CommandColor) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	send := insertPlaceholders(c.Response, m)

	s.ChannelMessage(m.ChannelID, send)
}
