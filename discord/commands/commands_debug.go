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
	roles, _ := s.GuildRoles(guild.ID)

	switch content := params[1]; content {
	case "help":
		log.Printf("Usage options as follows: ")
		log.Printf("\troleinfo")
		log.Printf("\tupdaterolename @user rolename")
	case "roleinfo":
		log.Printf("Displaying roles for server: %s (%s)", guild.Name, guild.ID)
		for _, roleFromList := range roles {
			log.Printf("Name: %s - ID: %s - Perms: %d - Pos: %d", roleFromList.Name, roleFromList.ID, roleFromList.Permissions, roleFromList.Position)
		}
	case "updaterolename":
		if len(params) < 4 {
			log.Printf("!%s updaterolename <user> <role>", c.Common.Caller)
			return
		}
		newRoleName := strings.TrimPrefix(params[2], "<@")
		newRoleName = strings.TrimSuffix(newRoleName, ">")
		roleID := params[3]
		for _, roleFromList := range roles {
			if roleFromList.Name == roleID {
				s.GuildRoleEdit(guild.ID, roleFromList.ID, newRoleName, roleFromList.Color, roleFromList.Hoist, roleFromList.Permissions, roleFromList.Mentionable)
				log.Printf("Updated role")
				return
			}
		}
		log.Printf("Could not find role")
	case "checkperms":
		if len(params) < 3 {
			return
		}
		c, err := findCommandByName(*CmdList, params[2])
		if err != nil {
			log.Printf("error: %s", err.Error())
		} else {
			log.Printf("perms: %v", getRequiredPermissionNames(c.GetCommons().RequiredPermissions))
		}
	}

	return
}

// GetCommons gets the common struct for the command
func (c CommandDebug) GetCommons() CommandCommon {
	return c.Common
}
