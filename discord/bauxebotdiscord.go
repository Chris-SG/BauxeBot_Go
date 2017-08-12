package bauxebotdiscord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/Chris-SG/BauxeBot_Go/discord/commands"
	"github.com/bwmarrin/discordgo"
)

// Session for discord bot
var (
	discord *discordgo.Session
	err     error
	prefix  string
	cmdList []cmd.CommandColor
)

/*type Command interface {
	CreateCommand()
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
}*/

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var channel, _ = discord.State.Channel(m.ChannelID)
	perms, _ := discord.UserChannelPermissions(m.Author.ID, channel.ID)
	log.Printf("chan %s msg %s by id %s", channel.Name, m.Content, m.Author.ID)

	if channel.Name == "colour_requests" {
		if strings.HasPrefix(m.Content, prefix) {
			parts := strings.Split(strings.ToLower(m.Content), " ")
			trimmed := strings.TrimPrefix(parts[0], prefix)
			if (strings.Contains(parts[0], "colour") == true || strings.Contains(parts[0], "color") == true) && len(parts) >= 2 {
				addUserColour(m)
			} else if strings.Contains(trimmed, "clear") == true && (perms&(1<<14)) != 0 {
				log.Print("clearing?")
				msgList, _ := discord.ChannelMessages(channel.ID, 100, channel.LastMessageID, "0", "")

				for i := 0; i < len(msgList); i++ {
					log.Printf("%s\n", msgList[i].ID)
					if !strings.Contains(msgList[i].ID, "245143483403337729") {
						discord.ChannelMessageDelete(channel.ID, msgList[i].ID)
					}
				}
			}
			//245143483403337729 main ID
		} else {
			log.Printf("Trying to delete message \"%s\" in %s", m.Content, channel.ID)
			err = discord.ChannelMessageDelete(channel.ID, m.ID)
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
		}
	}

	/*if channel.Name != "colour_requests" {
		if m.Content[0] == '!' {
			log.Printf("%s : %s", channel.ID, m.Content)
			var msg, merr = s.ChannelMessageSend(channel.ID, "helo")
			if merr != nil {
				log.Printf("Failed to respond: %s", merr)
			} else {
				log.Printf("Wrote: %s", msg.Content)
			}
		}
	}*/

	for _, cmd := range cmdList {
		if strings.HasPrefix(m.Content, (prefix + cmd.Common.Caller)) {
			cmd.Execute(s, m)
		}
	}
}

func addUserColour(m *discordgo.MessageCreate) {
	var user = m.Author.ID
	channel, _ := discord.State.Channel(m.ChannelID)
	if channel == nil {
		log.Printf("Could not find channel: %s", m.ChannelID)
		return
	}

	guild, _ := discord.State.Guild(channel.GuildID)
	if guild == nil {
		log.Printf("Could not find guild: %s", channel.GuildID)
		return
	}

	parts := strings.Split(m.Content, " ")
	var newColor int
	var tmpColor int64
	if parts[1][0] == '#' {
		trimmed := strings.TrimPrefix(parts[1], "#")
		tmpColor, err = strconv.ParseInt(trimmed, 16, 32)
		newColor = int(tmpColor)
		log.Printf("Parsing as a hex value. Int value %d", newColor)
	} else {
		newColor, err = strconv.Atoi(parts[1])
	}
	if err != nil {
		log.Printf("User tried to change colour to %s: Invalid", parts[1])
		discord.ChannelMessageDelete(m.ChannelID, m.ID)
		return
	}

	for _, roleList := range guild.Roles {
		if roleList.Name == user {
			discord.GuildRoleEdit(guild.ID, roleList.ID, roleList.Name, newColor, false, 0, false)
			log.Printf("Updated role: %s", roleList.ID)
			discord.ChannelMessageDelete(m.ChannelID, m.ID)
			return
		}
	}

	newRole, _ := discord.GuildRoleCreate(guild.ID)
	discord.GuildRoleEdit(guild.ID, newRole.ID, m.Author.ID, newColor, false, 0, false)
	newRole.Position = 3
	discord.GuildMemberRoleAdd(guild.ID, m.Author.ID, newRole.ID)
	log.Printf("Created new role: %s", newRole.ID)
	discord.ChannelMessageDelete(m.ChannelID, m.ID)
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	log.Print("Ready!")
	discord.UpdateStatus(0, "Trying to make this work...")
}

func init() {
	log.Printf("Initializing session with env token DISCORD_BOT_TOKEN: %s", os.Getenv("DISCORD_BOT_TOKEN"))
	discord, err = discordgo.New()

	discord.Token = "Bot " + os.Getenv("DISCORD_BOT_TOKEN")

	channels := []string{}
	users := []string{}
	c := cmd.CommandColor{Cooldown: 0, Common: cmd.CommandCommon{Caller: "setcolor", Response: "Setting {NAME}'s color to {ARG1}.", Description: "Sets user's color", Structure: "!setcolor <color> (hex)", Channels: channels, RequiredPermissions: 0, RequiredUsers: users}}
	cmdList = append(cmdList, c)
}

// StartBotDiscord will Start Discord bot
func StartBotDiscord(cmdPrefix string) {
	log.Print("Starting session...")
	prefix = cmdPrefix
	discord.State.User, err = discord.User("@me")
	if err != nil {
		log.Printf("Error: %s", err)
	}

	discord.AddHandler(onReady)
	discord.AddHandler(onMessageCreate)

	log.Print("Opening session...")
	err = discord.Open()
	if err != nil {
		log.Printf("Failed to open session: %s", err)
	}

	// Wait for a signal to quit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
