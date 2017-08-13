package bauxebotdiscord

import (
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/Chris-SG/BauxeBot_Go/discord/commands"
	"github.com/bwmarrin/discordgo"
)

// Session for discord bot
var (
	discord *discordgo.Session
	err     error
	prefix  string
	cmdList cmd.Commands
	bot     *discordgo.User
)

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Echo out response
	var channel, _ = discord.State.Channel(m.ChannelID)
	log.Printf("chan %s msg %s by id %s", channel.Name, m.Content, m.Author.ID)

	// We don't want bot-loop sort things, maybe add support for disabling commands for all bots
	if m.Author.ID == bot.ID {
		return
	}

	// Check all dummy commands
	for _, dummyCmd := range cmdList.DummyCommands {
		if strings.HasPrefix(m.Content, (prefix + dummyCmd.Common.Caller)) {
			go dummyCmd.Execute(s, m)
			return
		}
	}

	// Check all color commands
	for _, colorCmd := range cmdList.ColorCommands {
		if strings.HasPrefix(m.Content, (prefix + colorCmd.Common.Caller)) {
			go colorCmd.Execute(s, m)
			return
		}
	}

	// Check all debug commands
	for _, debugCmd := range cmdList.DebugCommands {
		if strings.HasPrefix(m.Content, (prefix + debugCmd.Common.Caller)) {
			go debugCmd.Execute(s, m)
			return
		}
	}
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	log.Print("Ready!")
	discord.UpdateStatus(0, "Trying to make this work...")
	bot, _ = discord.User("@me")
}

func init() {
	// Get token from env variables. %%TO ADD TO XML%%
	log.Printf("Initializing session with env token DISCORD_BOT_TOKEN: %s", os.Getenv("DISCORD_BOT_TOKEN"))
	discord, err = discordgo.New()

	discord.Token = "Bot " + os.Getenv("DISCORD_BOT_TOKEN")

	// Test commands, will make more elegant in time
	var c cmd.Command
	c = cmd.CommandColor{Common: cmd.CommandCommon{Caller: "color", Response: "Setting {HL_NAME}'s color to #{ARG1}.", Description: "Sets user's color", Structure: "!setcolor <color> (hex)", Channels: []string{}, RequiredPermissions: 0, RequiredUsers: []string{}}}
	cmdList.ColorCommands = append(cmdList.ColorCommands, c.(cmd.CommandColor))
	c = cmd.CommandDummy{Common: cmd.CommandCommon{Caller: "helo", Response: "helo", Description: "helo", Structure: "!helo", Channels: []string{}, RequiredPermissions: 0, RequiredUsers: []string{}}}
	cmdList.DummyCommands = append(cmdList.DummyCommands, c.(cmd.CommandDummy))
	c = cmd.CommandDebug{Common: cmd.CommandCommon{Caller: "debug", Response: "", Description: "debug", Structure: "!debug <param>", Channels: []string{}, RequiredPermissions: 8, RequiredUsers: []string{}}}
	cmdList.DebugCommands = append(cmdList.DebugCommands, c.(cmd.CommandDebug))
}

// StartBotDiscord will Start Discord bot
func StartBotDiscord(cmdPrefix string) {
	log.Print("Starting session...")
	prefix = cmdPrefix
	discord.State.User, err = discord.User("@me")
	if err != nil {
		log.Printf("Error: %s", err)
	}

	// Add discord handlers
	discord.AddHandler(onReady)
	discord.AddHandler(onMessageCreate)

	// Open discord session
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
