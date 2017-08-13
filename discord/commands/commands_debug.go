package cmd

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandDebug represents a struct for changing name colors
type CommandDebug struct {
	Common CommandCommon
}

// Execute acts upon a given message
func (c CommandDebug) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	c.Common.RequiredPermissions = 8
	log.Print("Debug called")

	if !c.Common.canExecute(s, m) {
		log.Print("Insufficient Permissions.")
		return
	}

	params := strings.Split(m.Content, " ")
	if len(params) < 2 {
		log.Print("No parameters provided.")
		return
	}

	log.Printf("%s", params[1])

	channel, _ := s.State.Channel(m.ChannelID)
	guild, _ := s.State.Guild(channel.GuildID)

	switch content := params[1]; content {
	case "roleinfo":
		roles, _ := s.GuildRoles(guild.ID)
		log.Printf("Displaying roles for server: %s (%s)", guild.Name, guild.ID)
		for _, roleFromList := range roles {
			log.Printf("Name: %s - ID: %s - Perms: %d - Pos: %d", roleFromList.Name, roleFromList.ID, roleFromList.Permissions, roleFromList.Position)
		}
	}

	return
}
