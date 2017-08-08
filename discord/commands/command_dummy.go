package bauxebotdiscordcmd

import "github.com/bwmarrin/discordgo"

type commandDummy struct {
	response string
	cooldown int
}

func (c commandDummy) execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	send := insertPlaceholders(c.response, m)

	s.ChannelMessage(m.ChannelID, send)
}
