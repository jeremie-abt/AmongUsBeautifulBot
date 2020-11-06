package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Managing VoiceUpdate change
func VoiceChangeHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	currentGuild := GlobalVarManager.getGuildObj(m.VoiceState.GuildID)
	currentGuild.HandleVoiceChange(m.VoiceState)

}

func MessageSendHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	curMessage := m.Message
	currentGuild := GlobalVarManager.getGuildObj(curMessage.GuildID)

	// TODO: Faire gaf a gerer le type du channel, je ne
	// veux pas listen les dm (pour l'instant)

	if strings.HasPrefix(
		strings.ToLower(curMessage.Content), ".creategame") {
		// get le channel pour commencer une game
		HandleCreateGame(s, curMessage, currentGuild)

	} else if strings.HasPrefix(
		strings.ToLower(curMessage.Content), ".stopgame") {

		// TODO: si pas d'arg, propal de toutes les games
		// a delete en fonction de sa guild
		HandleStopGame(s, curMessage, currentGuild)
	}
}
