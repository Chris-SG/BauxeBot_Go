package cmd

import "github.com/bwmarrin/discordgo"

// Command works as an interface for other commands
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
}

// Commands is a struct to hold all bot commands by type
type Commands struct {
	DummyCommands      []CommandDummy
	ColorCommands      []CommandColor
	DebugCommands      []CommandDebug
	ModerationCommands []CommandModeration
}

/*CommandCommon represents fields common between all commands

 Fields are as follows:
 	Caller
		Text used to call command
	Response
		Bot's response to command request
	Description
		Basic description of command result
	Structure
		Format for command
	Channels
		Channel(s) this command can be used in. Blank for any
	RequiredPermissions
		Channel permissions required to use command
	RequiredUsers
		UserID(s) required to use command
*/
type CommandCommon struct {
	Caller              string
	Response            string
	Description         string
	Structure           string
	Action              string
	Channels            []string
	RequiredPermissions int
	RequiredUsers       []string
}

// all permissions
const (
	createInstantInvite = iota //other
	kickMembers         = iota //other
	banMembers          = iota //other
	administrator       = iota //other
	manageChannels      = iota //manage
	manageGuild         = iota //manage
	addReactions        = iota //text
	viewAuditLog        = iota //other
	readMessages        = iota //text
	sendMessages        = iota //text
	sendTTSMessages     = iota //text
	manageMessages      = iota //text
	embedLinks          = iota //text
	attachFiles         = iota //text
	readMessageHistory  = iota //text
	mentionEveryone     = iota //text
	useExternalEmojis   = iota //text
	connect             = iota //voice
	speak               = iota //voice
	muteMembers         = iota //voice
	deafenMembers       = iota //voice
	moveMembers         = iota //voice
	useVAD              = iota //voice
	changeNickname      = iota //other
	manageNicknames     = iota //manage
	manageRoles         = iota //manage
	manageWebhooks      = iota //manage
	manageEmojis        = iota //manage
)

// Check if a user can execute a command
func (c CommandCommon) canExecute(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	userPerms, _ := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

	// Most commands won't have tied permissions, so this speeds most up
	if c.RequiredPermissions == 0 && len(c.RequiredUsers) == 0 {
		return true
	}

	// This is pretty much untested. Need to add check for required users
	if !checkVoicePerms(userPerms, c.RequiredPermissions) ||
		!checkTextPerms(userPerms, c.RequiredPermissions) ||
		!checkManagementPerms(userPerms, c.RequiredPermissions) ||
		!checkOtherPerms(userPerms, c.RequiredPermissions) {
		return false
	}

	return true
}

// Send error
func (c CommandCommon) sendErrorResponse(s *discordgo.Session, channelID string) {
	response := "Must use correct format: " + c.Structure
	s.ChannelMessageSend(channelID, response)
	return
}

// CheckPerm checks if a user has a specified permission
func CheckPerm(userPerms int, perm int) bool {
	if userPerms&(1<<uint(perm)) != 0 {
		return true
	}
	return false
}

func checkOtherPerms(userPerms int, commandPerms int) bool {
	if userPerms&(1<<createInstantInvite) == 0 {
		if commandPerms&(1<<createInstantInvite) != 0 {
			return false
		}
	}
	if userPerms&(1<<kickMembers) == 0 {
		if commandPerms&(1<<kickMembers) != 0 {
			return false
		}
	}
	if userPerms&(1<<banMembers) == 0 {
		if commandPerms&(1<<banMembers) != 0 {
			return false
		}
	}
	if userPerms&(1<<administrator) == 0 {
		if commandPerms&(1<<administrator) != 0 {
			return false
		}
	}
	if userPerms&(1<<viewAuditLog) == 0 {
		if commandPerms&(1<<viewAuditLog) != 0 {
			return false
		}
	}
	if userPerms&(1<<changeNickname) == 0 {
		if commandPerms&(1<<changeNickname) != 0 {
			return false
		}
	}

	return true
}

func checkManagementPerms(userPerms int, commandPerms int) bool {
	if userPerms&(1<<manageChannels) == 0 {
		if commandPerms&(1<<manageChannels) != 0 {
			return false
		}
	}
	if userPerms&(1<<manageEmojis) == 0 {
		if commandPerms&(1<<manageEmojis) != 0 {
			return false
		}
	}
	if userPerms&(1<<manageGuild) == 0 {
		if commandPerms&(1<<manageGuild) != 0 {
			return false
		}
	}
	if userPerms&(1<<manageMessages) == 0 {
		if commandPerms&(1<<manageMessages) != 0 {
			return false
		}
	}
	if userPerms&(1<<manageNicknames) == 0 {
		if commandPerms&(1<<manageNicknames) != 0 {
			return false
		}
	}
	if userPerms&(1<<manageRoles) == 0 {
		if commandPerms&(1<<manageRoles) != 0 {
			return false
		}
	}
	if userPerms&(1<<manageWebhooks) == 0 {
		if commandPerms&(1<<manageWebhooks) != 0 {
			return false
		}
	}

	return true
}

func checkTextPerms(userPerms int, commandPerms int) bool {
	if userPerms&(1<<addReactions) == 0 {
		if commandPerms&(1<<addReactions) != 0 {
			return false
		}
	}
	if userPerms&(1<<readMessages) == 0 {
		if commandPerms&(1<<readMessages) != 0 {
			return false
		}
	}
	if userPerms&(1<<sendMessages) == 0 {
		if commandPerms&(1<<sendMessages) != 0 {
			return false
		}
	}
	if userPerms&(1<<sendTTSMessages) == 0 {
		if commandPerms&(1<<sendTTSMessages) != 0 {
			return false
		}
	}
	if userPerms&(1<<manageMessages) == 0 {
		if commandPerms&(1<<manageMessages) != 0 {
			return false
		}
	}
	if userPerms&(1<<embedLinks) == 0 {
		if commandPerms&(1<<embedLinks) != 0 {
			return false
		}
	}
	if userPerms&(1<<attachFiles) == 0 {
		if commandPerms&(1<<attachFiles) != 0 {
			return false
		}
	}
	if userPerms&(1<<readMessageHistory) == 0 {
		if commandPerms&(1<<readMessageHistory) != 0 {
			return false
		}
	}
	if userPerms&(1<<mentionEveryone) == 0 {
		if commandPerms&(1<<mentionEveryone) != 0 {
			return false
		}
	}
	if userPerms&(1<<useExternalEmojis) == 0 {
		if commandPerms&(1<<useExternalEmojis) != 0 {
			return false
		}
	}

	return true
}

func checkVoicePerms(userPerms int, commandPerms int) bool {
	if userPerms&(1<<connect) == 0 {
		if commandPerms&(1<<connect) != 0 {
			return false
		}
	}
	if userPerms&(1<<speak) == 0 {
		if commandPerms&(1<<speak) != 0 {
			return false
		}
	}
	if userPerms&(1<<muteMembers) == 0 {
		if commandPerms&(1<<muteMembers) != 0 {
			return false
		}
	}
	if userPerms&(1<<deafenMembers) == 0 {
		if commandPerms&(1<<deafenMembers) != 0 {
			return false
		}
	}
	if userPerms&(1<<moveMembers) == 0 {
		if commandPerms&(1<<moveMembers) != 0 {
			return false
		}
	}
	if userPerms&(1<<useVAD) == 0 {
		if commandPerms&(1<<useVAD) != 0 {
			return false
		}
	}

	return true
}
