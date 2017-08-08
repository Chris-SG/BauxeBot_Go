package bauxebotdiscordcmd

import "github.com/bwmarrin/discordgo"

type command interface {
	execute(s *discordgo.Session, m *discordgo.MessageCreate)
}

func init() {

}
