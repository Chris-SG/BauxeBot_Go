package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CommandModeration represents a struct for moderation commands
type CommandModeration struct {
	Common CommandCommon
}

// Execute acts upon a given message
func (c CommandModeration) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !c.Common.canExecute(s, m) {
		log.Print("Insufficient Permissions.")
	}

	channel, _ := s.State.Channel(m.ChannelID)

	params := strings.Split(m.Content, " ")
	switch action := c.Common.Action; action {
	case "deletebulk":
		log.Print("Deleting messages")
		count, err := strconv.Atoi(params[1])
		if err != nil {
			log.Printf("Error: %s", err)
			return
		}
		var messages []*discordgo.Message
		messages, _ = s.ChannelMessages(channel.ID, count, "", "", channel.LastMessageID)
		var messagesToDelete []string
		for _, message := range messages {
			messagesToDelete = append(messagesToDelete, message.ID)
		}
		s.ChannelMessagesBulkDelete(m.ChannelID, messagesToDelete)
	}
}

// GetCommons gets the common struct for the command
func (c CommandModeration) GetCommons() CommandCommon {
	return c.Common
}
