package cmd

import (
	"container/list"

	"github.com/bwmarrin/discordgo"
)

// Command works as an interface for other commands
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
}

// CommandCommon represents fields common between all commands
type CommandCommon struct {
	Caller              string
	Response            string
	Description         string
	Channels            *list.List
	RequiredPermissions int
	RequiredUsers       *list.List
}

const (
	create_instant_invite = iota
	kick_members          = iota
	ban_members           = iota
	administrator         = iota
	manage_channels       = iota
	manage_guild          = iota
	add_reactions         = iota
	view_audit_log        = iota
	read_messages         = iota
	send_messages         = iota
	send_tts_messages     = iota
	manage_messages       = iota
	embed_links           = iota
	attach_files          = iota
	read_message_history  = iota
	mention_everyone      = iota
	use_external_emojis   = iota
	connect               = iota //voice
	speak                 = iota //voice
	mute_members          = iota //voice
	deafen_members        = iota //voice
	move_members          = iota //voice
	use_vad               = iota //voice
	change_nickname       = iota
	manage_nicknames      = iota
	manage_roles          = iota
	manage_webhooks       = iota
	manage_emojis         = iota
)

func (c CommandCommon) canExecute(s *discordgo.Session, m *discordgo.MessageCreate) (canExecute bool) {
	userPerms, _ := s.UserChannelPermissions(m.Author.ID, m.ChannelID)

	canExecute = checkVoicePerms(userPerms, c.RequiredPermissions)

	return true
}

func checkVoicePerms(userPerms int, commandPerms int) (canExecute bool) {
	if userPerms&(1<<connect) != 0 {
		if commandPerms&(1<<connect) == 0 {
			return false
		}
	}
	if userPerms&(1<<speak) != 0 {
		if commandPerms&(1<<speak) == 0 {
			return false
		}
	}
	if userPerms&(1<<mute_members) != 0 {
		if commandPerms&(1<<mute_members) == 0 {
			return false
		}
	}
	if userPerms&(1<<deafen_members) != 0 {
		if commandPerms&(1<<deafen_members) == 0 {
			return false
		}
	}
	if userPerms&(1<<move_members) != 0 {
		if commandPerms&(1<<move_members) == 0 {
			return false
		}
	}
	if userPerms&(1<<use_vad) != 0 {
		if commandPerms&(1<<use_vad) == 0 {
			return false
		}
	}

	return true
}
