package cmd

import "github.com/bwmarrin/discordgo"

// This file will be used for actions that may be required by all commands
// This will include:
// Delete caller message
// Delete respond message after time

// DeleteMessage a
func DeleteMessage() {

}

func orderRolesByPositon(r discordgo.Roles) {
	for i := 0; i < len(r)-1; i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i].Position > r[j].Position {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
}
