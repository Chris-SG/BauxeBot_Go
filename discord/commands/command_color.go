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
		// Check if user sent 6 characters (hex representation)
		parts := strings.Split(m.Content, " ")
		if len(parts) < 2 || len(parts[1]) != 6 {
			if len(parts[1]) == 7 && parts[1][0] == '#' {
				parts[1] = strings.TrimPrefix(parts[1], "#")
			}
			c.Common.sendErrorResponse(s, m.ChannelID)
			return
		}
		roleColor, err := c.hexToInt(parts[1])
		if err != nil {
			c.Common.sendErrorResponse(s, m.ChannelID)
			return
		}
		channel, _ := s.State.Channel(m.ChannelID)
		c.createRoleWithColor(s, channel.GuildID, m.Author.ID, roleColor)
		s.ChannelMessageSend(m.ChannelID, send)
		log.Print("Done")
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

	roleList := guild.Roles
	newRole, _ := s.GuildRoleCreate(guildID)
	newRole, _ = s.GuildRoleEdit(guildID, newRole.ID, roleName, color, false, 0, false)
	s.GuildMemberRoleAdd(guildID, roleName, newRole.ID)
	roleList = append(roleList, newRole)
	newRole.Position = len(roleList)

	orderRolesByPositon(roleList)

	s.GuildRoleReorder(guildID, roleList)

	for _, roleFromList := range roleList {
		log.Printf("Rolename: %s, Position: %d", roleFromList.Name, roleFromList.Position)
	}

	return newRole
}
