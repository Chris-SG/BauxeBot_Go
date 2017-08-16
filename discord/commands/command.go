package cmd

import "github.com/bwmarrin/discordgo"

// Command works as an interface for other commands
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
	GetCommons() CommandCommon
}

// Commands is a struct to hold all bot commands by type
type Commands struct {
	DummyCommands      []CommandDummy
	ColorCommands      []CommandColor
	DebugCommands      []CommandDebug
	ModerationCommands []CommandModeration
}

var CmdList *Commands

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

type Permission int

// all permissions
const (
	createInstantInvite Permission = 1 << iota //other
	kickMembers                                //other
	banMembers                                 //other
	administrator                              //other
	manageChannels                             //manage
	manageGuild                                //manage
	addReactions                               //text
	viewAuditLog                               //other
	readMessages                               //text
	sendMessages                               //text
	sendTTSMessages                            //text
	manageMessages                             //text
	embedLinks                                 //text
	attachFiles                                //text
	readMessageHistory                         //text
	mentionEveryone                            //text
	useExternalEmojis                          //text
	connect                                    //voice
	speak                                      //voice
	muteMembers                                //voice
	deafenMembers                              //voice
	moveMembers                                //voice
	useVAD                                     //voice
	changeNickname                             //other
	manageNicknames                            //manage
	manageRoles                                //manage
	manageWebhooks                             //manage
	manageEmojis                               //manage
)

var permissionMap = map[Permission]string{
	createInstantInvite: "CreateInstantInvite",
	kickMembers:         "KickMembers",
	banMembers:          "BanMembers",
	administrator:       "Administrator",
	manageChannels:      "ManageChannels",
	manageGuild:         "ManageGuild",
	addReactions:        "AddReactions",
	viewAuditLog:        "ViewAuditLog",
	readMessages:        "ReadMessages",
	sendMessages:        "SendMessages",
	sendTTSMessages:     "SendTTSMessages",
	manageMessages:      "ManageMessages",
	embedLinks:          "EmbedLinks",
	attachFiles:         "AttachFiles",
	readMessageHistory:  "ReadMessageHistory",
	mentionEveryone:     "MentionEveryone",
	useExternalEmojis:   "UseExternalEmojis",
	connect:             "Connect",
	speak:               "Speak",
	muteMembers:         "MuteMembers",
	deafenMembers:       "DeafenMembers",
	moveMembers:         "MoveMembers",
	useVAD:              "UseVAD",
	changeNickname:      "ChangeNickname",
	manageNicknames:     "ManageNicknames",
	manageRoles:         "ManageRoles",
	manageWebhooks:      "ManageWebhooks",
	manageEmojis:        "ManageEmojis",
}

// Check if a user can execute a command
func (c CommandCommon) canExecute(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	userPerms, _ := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

	// Most commands won't have tied permissions, so this speeds most up
	if c.RequiredPermissions == 0 && len(c.RequiredUsers) == 0 {
		return true
	}

	// This is pretty much untested. Need to add check for required users
	/*if !checkVoicePerms(userPerms, c.RequiredPermissions) ||
		!checkTextPerms(userPerms, c.RequiredPermissions) ||
		!checkManagementPerms(userPerms, c.RequiredPermissions) ||
		!checkOtherPerms(userPerms, c.RequiredPermissions) {
		return false
	}*/
	if userPerms&c.RequiredPermissions != c.RequiredPermissions {
		return false
	}

	return true
}

func getRequiredPermissionNames(permissions int) (perms *[]string) {
	for val, name := range permissionMap {
		if val&Permission(permissions) != 0 {
			*perms = append(*perms, name)
		}
	}
	return
}

// Send error
func (c CommandCommon) sendErrorResponse(s *discordgo.Session, channelID string) {
	response := "Must use correct format: " + c.Structure
	s.ChannelMessageSend(channelID, response)
	return
}

// CheckPerm checks if a user has a specified permission
func CheckPerm(userPerms int, perm int) bool {
	if userPerms&(perm) != 0 {
		return true
	}
	return false
}

func checkOtherPerms(userPerms Permission, commandPerms Permission) bool {
	if userPerms&createInstantInvite == 0 {
		if commandPerms&createInstantInvite != 0 {
			return false
		}
	}
	if userPerms&kickMembers == 0 {
		if commandPerms&kickMembers != 0 {
			return false
		}
	}
	if userPerms&banMembers == 0 {
		if commandPerms&banMembers != 0 {
			return false
		}
	}
	if userPerms&administrator == 0 {
		if commandPerms&administrator != 0 {
			return false
		}
	}
	if userPerms&viewAuditLog == 0 {
		if commandPerms&viewAuditLog != 0 {
			return false
		}
	}
	if userPerms&changeNickname == 0 {
		if commandPerms&changeNickname != 0 {
			return false
		}
	}

	return true
}

func checkManagementPerms(userPerms Permission, commandPerms Permission) bool {
	if userPerms&manageChannels == 0 {
		if commandPerms&manageChannels != 0 {
			return false
		}
	}
	if userPerms&manageEmojis == 0 {
		if commandPerms&manageEmojis != 0 {
			return false
		}
	}
	if userPerms&manageGuild == 0 {
		if commandPerms&manageGuild != 0 {
			return false
		}
	}
	if userPerms&manageMessages == 0 {
		if commandPerms&manageMessages != 0 {
			return false
		}
	}
	if userPerms&manageNicknames == 0 {
		if commandPerms&manageNicknames != 0 {
			return false
		}
	}
	if userPerms&manageRoles == 0 {
		if commandPerms&manageRoles != 0 {
			return false
		}
	}
	if userPerms&manageWebhooks == 0 {
		if commandPerms&manageWebhooks != 0 {
			return false
		}
	}

	return true
}

func checkTextPerms(userPerms Permission, commandPerms Permission) bool {
	if userPerms&addReactions == 0 {
		if commandPerms&addReactions != 0 {
			return false
		}
	}
	if userPerms&readMessages == 0 {
		if commandPerms&readMessages != 0 {
			return false
		}
	}
	if userPerms&sendMessages == 0 {
		if commandPerms&sendMessages != 0 {
			return false
		}
	}
	if userPerms&sendTTSMessages == 0 {
		if commandPerms&sendTTSMessages != 0 {
			return false
		}
	}
	if userPerms&manageMessages == 0 {
		if commandPerms&manageMessages != 0 {
			return false
		}
	}
	if userPerms&embedLinks == 0 {
		if commandPerms&embedLinks != 0 {
			return false
		}
	}
	if userPerms&attachFiles == 0 {
		if commandPerms&attachFiles != 0 {
			return false
		}
	}
	if userPerms&readMessageHistory == 0 {
		if commandPerms&readMessageHistory != 0 {
			return false
		}
	}
	if userPerms&mentionEveryone == 0 {
		if commandPerms&mentionEveryone != 0 {
			return false
		}
	}
	if userPerms&useExternalEmojis == 0 {
		if commandPerms&useExternalEmojis != 0 {
			return false
		}
	}

	return true
}

func checkVoicePerms(userPerms Permission, commandPerms Permission) bool {
	if userPerms&connect == 0 {
		if commandPerms&connect != 0 {
			return false
		}
	}
	if userPerms&speak == 0 {
		if commandPerms&speak != 0 {
			return false
		}
	}
	if userPerms&muteMembers == 0 {
		if commandPerms&muteMembers != 0 {
			return false
		}
	}
	if userPerms&deafenMembers == 0 {
		if commandPerms&deafenMembers != 0 {
			return false
		}
	}
	if userPerms&moveMembers == 0 {
		if commandPerms&moveMembers != 0 {
			return false
		}
	}
	if userPerms&useVAD == 0 {
		if commandPerms&useVAD != 0 {
			return false
		}
	}

	return true
}
