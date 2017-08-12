package cmd

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func replace(s string, m *discordgo.MessageCreate) (res string) {
	switch s {
	case "NAME":
		res = m.Author.Username
	case "HL_NAME":
		res = "<@" + m.Author.ID + ">"
	case "ID":
		res = m.Author.ID
	case "ARG1":
		parts := strings.Split(m.Content, " ")
		res = parts[1]
	default:
		res = "[" + s + "]"
	}

	return
}

func insertPlaceholders(s string, m *discordgo.MessageCreate) (res string) {
	//lastSearchPos := -1
	log.Printf("starting as: %s", s)
	searching := true

	// Find text between {} and send to replace func
	// This will have issues if the {} doesn't get replaced
	for searching == true {
		searching = false
		parts := strings.SplitN(s, "{", 1)
		if len(parts) > 1 {
			parts2 := strings.SplitAfterN(parts[1], "}", 1)
			if len(parts2) > 1 {
				res = parts[0] + replace(parts2[0], m) + parts2[1]
				searching = true
			}
		} else {
			res = parts[0]
		}
	}

	log.Printf("finishing as: %s", res)
	return
}
