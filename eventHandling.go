package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type AUEventType string

type socketContext struct {
	code string
}

const (
	CONNECTCODE AUEventType = "connectCode"
	LOBBY                   = "lobby"
	STATE                   = "state"
	PLAYER                  = "player"
)

type eventType struct {
	eventType AUEventType
	msg       string
}

/*
**	Managing Discord event
*/

// Managing VoiceUpdate change
func VoiceChangeHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	currentGuild := G_Gvm.getGuildObj(m.VoiceState.GuildID)
	currentGuild.HandleVoiceChange(m.VoiceState)
}

func MessageSendHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var currentGuild *GuildManagerType

	curMessage := m.Message
	currentGuild = G_Gvm.getGuildObj(curMessage.GuildID)
	if currentGuild == nil {
		currentGuild = G_Gvm.AddGuild(curMessage.GuildID)
	}
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

/*
**	Managing socket Event (amongus Capture)
*/

func SocketConnectCodeHandle(connectCode string) bool {

	gameToConnect := G_Gvm.GetGameCode(connectCode)
	if gameToConnect == nil {
		// TODO: Gerer et enregistrer l'erreur proprement
		fmt.Printf("game Not Existing")
		return false
	}
	// TODO: Faire un vraie system de logj
	println("Game connected\n")
	gameToConnect.IsCapturedConn = true
	return true
}
