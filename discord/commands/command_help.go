package cmd

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandHelp represents a help command
type CommandHelp struct {
	Common CommandCommon
}

// Execute acts upon a given message
func (c CommandHelp) Execute(s *discordgo.Session, m *discordgo.MessageCreate, cmds Commands) {
	delimiter := string(m.Content[0])
	params := strings.Split(m.Content, " ")

	if len(params) < 2 {

	} else {
		switch result := params[1]; result {
		case "hello":
			return
		default:
			if string(result[0]) == delimiter {
				cmdHelp, err := findCommandByName(cmds, strings.TrimLeft(result, delimiter))
				if err != nil {
					log.Printf("error: %s", err)
				} else {
					bot, _ := s.User("@me")
					var embed *discordgo.MessageEmbed
					embed.Author.Name = bot.Username
					embed.Author.IconURL = bot.Avatar
					log.Printf("%s", cmdHelp.GetCommons().Caller)

				}
			}
			return
		}
	}
}

// GetCommons gets the common struct for the command
func (c CommandHelp) GetCommons() CommandCommon {
	return c.Common
}
