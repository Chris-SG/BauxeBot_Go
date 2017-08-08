package bauxebotdiscordcmd

import "github.com/bwmarrin/discordgo"

type commandColor struct {
	response string
	cooldown int
}

func (c commandColor) execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	send := insertPlaceholders(c.response, m)

	s.ChannelMessage(m.ChannelID, send)
}
