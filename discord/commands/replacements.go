package bauxebotdiscordcmd

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func replace(s string, m *discordgo.MessageCreate) (res string) {
	switch s {
	case "NAME":
		res = m.Author.Username
	case "ID":
		res = m.Author.ID
	default:
		res = "{" + s + "}"
	}

	return
}

func insertPlaceholders(s string, m *discordgo.MessageCreate) (res string) {
	//lastSearchPos := -1
	searching := true

	// Find text between {} and send to replace func
	// This will have issues if the {} doesn't get replaced
	for searching == true {
		searching = false
		parts := strings.SplitN(s, "{", 1)
		if len(parts) > 1 {
			parts2 := strings.SplitAfterN(parts[1], "}", 1)
			if len(parts2) > 1 {
				s = parts[0] + replace(parts2[0], m) + parts2[1]
				searching = true
			}
		}
	}
	return
}
