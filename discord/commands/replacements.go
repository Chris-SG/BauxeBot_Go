package cmd

// Limitation of current implementation is no support for {} in text

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// replace will replace a placeholder
func replace(s string, m *discordgo.MessageCreate) (res string) {
	switch s {
	case "NAME": // Command author's name
		res = m.Author.Username
	case "HL_NAME": // Highlight command author
		res = "<@" + m.Author.ID + ">"
	case "ID": // Author ID
		res = m.Author.ID
	case "ARG1": // argument 1 of author's message
		parts := strings.Split(m.Content, " ")
		res = parts[1]
	default:
		res = s
	}

	return
}

// split will slice a string at all { and }
func split(r rune) bool {
	return r == '{' || r == '}'
}

// insertPlaceholders will fill in placeholder text
func insertPlaceholders(s string, m *discordgo.MessageCreate) (res string) {
	//lastSearchPos := -1
	log.Printf("starting as: %s", s)

	// slice the response
	parts := strings.FieldsFunc(s, split)

	for _, part := range parts {
		res += replace(part, m)
	}

	log.Printf("finishing as: %s", res)
	return
}
