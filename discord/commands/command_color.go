package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandColor represents a struct for changing name colors
type CommandColor struct {
	Common CommandCommon
}

// Execute reporesents acting upon the color command
func (c CommandColor) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Replace any placeholder text ({})
	send := insertPlaceholders(c.Common.Response, m)
	log.Printf("Trying to set color %s", m.Content)

	// Check if user has correct permissions
	if c.Common.canExecute(s, m) {
		// What is the command action
		switch action := c.Common.Action; action {
		case "setcolor":
			// Check if user sent 6 characters (hex representation)
			parts := strings.Split(m.Content, " ")
			if len(parts) < 2 || len(parts[1]) != 6 {
				if len(parts[1]) == 7 && parts[1][0] == '#' {
					parts[1] = strings.TrimPrefix(parts[1], "#")
				}
				c.Common.sendErrorResponse(s, m.ChannelID)
				return
			}
			// Convert hex to int for role
			roleColor, err := c.hexToInt(parts[1])
			if err != nil {
				c.Common.sendErrorResponse(s, m.ChannelID)
				return
			}
			channel, _ := s.State.Channel(m.ChannelID)
			c.createRoleWithColor(s, channel.GuildID, m.Author.ID, roleColor)
			s.ChannelMessageSend(m.ChannelID, send)
			log.Print("Done")
		case "removecolor":
			channel, _ := s.State.Channel(m.ChannelID)
			guild, _ := s.State.Guild(channel.GuildID)
			for _, roleFromList := range guild.Roles {
				if roleFromList.Name == m.Author.ID {
					s.GuildRoleDelete(guild.ID, roleFromList.ID)
					s.ChannelMessageSend(m.ChannelID, send)
					log.Print("Done")
					return
				}
			}
			s.ChannelMessageSend(m.ChannelID, "Could not find role.")
		}
	} else {
		log.Print("Can't execute")
	}

	return
}

func (c CommandColor) hexToInt(color string) (colorInt int, err error) {
	tmpColor, _ := strconv.ParseInt(color, 16, 32)
	colorInt = int(tmpColor)
	return
}

func (c CommandColor) createRoleWithColor(s *discordgo.Session, guildID string, roleName string, color int) *discordgo.Role {
	guild, _ := s.State.Guild(guildID)
	for _, roleFromList := range guild.Roles {
		if roleFromList.Name == roleName {
			s.GuildRoleEdit(guildID, roleFromList.ID, roleFromList.Name, color, false, 0, false)
			return roleFromList
		}
	}

	newRole, _ := s.GuildRoleCreate(guildID)
	newRole, _ = s.GuildRoleEdit(guildID, newRole.ID, roleName, color, false, 0, false)
	s.GuildMemberRoleAdd(guildID, roleName, newRole.ID)
	roleList := guild.Roles

	orderRolesByPositon(roleList)
	for i := 2; i < len(roleList)-1; i++ {
		roleList[i].Position--
	}
	roleList[1].Position = len(roleList) - 2
	orderRolesByPositon(roleList)

	newRoleList, err := s.GuildRoleReorder(guildID, roleList)
	if err != nil {
		log.Printf("error: %s", err)
	}

	for _, roleFromList := range newRoleList {
		log.Printf("Rolename: %s, Position: %d", roleFromList.Name, roleFromList.Position)
	}

	return newRole
}
