package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandColor represents a struct for changing name colors
type CommandColor struct {
	Cooldown int
	Common   CommandCommon
}

// Execute reporesents acting upon the color command
func (c CommandColor) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	send := insertPlaceholders(c.Common.Response, m)
	log.Printf("Trying to set color %s", m.Content)
	if c.Common.canExecute(s, m) {
		parts := strings.Split(m.Content, " ")
		if len(parts) < 2 || len(parts[1]) != 6 {
			c.Common.sendErrorResponse(s, m.ChannelID)
			return
		}
		roleColor, err := c.hexToInt(parts[1])
		if err != nil {
			c.Common.sendErrorResponse(s, m.ChannelID)
			return
		}
		channel, _ := s.State.Channel(m.ChannelID)
		log.Print("Step1")
		c.createRoleWithColor(s, channel.GuildID, m.Author.ID, roleColor)
		log.Print("Step2")
		s.ChannelMessageSend(m.ChannelID, send)
		log.Print("Step3")
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
	for _, roleList := range guild.Roles {
		if roleList.Name == roleName {
			s.GuildRoleEdit(guildID, roleList.ID, roleList.Name, color, false, 0, false)
			return roleList
		}
	}

	newRole, _ := s.GuildRoleCreate(guildID)
	s.GuildRoleEdit(guildID, newRole.ID, roleName, color, false, 0, false)
	s.GuildMemberRoleAdd(guildID, roleName, newRole.ID)
	roleList := guild.Roles

	for i := len(roleList) - 1; i > 3; i++ {
		roleList[i] = roleList[i-1]
	}
	roleList[3] = newRole
	s.GuildRoleReorder(guildID, roleList)

	return newRole
}
