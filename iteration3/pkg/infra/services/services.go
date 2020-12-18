package services

import "github.com/bwmarrin/discordgo"

//TODO : Reflechir a cette partie du design, il y a un concept
// qui n'ait pas bien abstrait ici
// implementation du bot pour discord avec la lib discordgo
type IDiscordBotWithDiscordGo interface {
	// TODO: Je laisse ces deux fonction mais je ne pense pas qu'elles
	// Doivent rester a terme ?
	StartGame(gameId string) error
	DeleteGame(gameId string) error

	HandleWrittenMessage(sess *discordgo.Session, msg *discordgo.MessageCreate)
	//HandleWrittenMessage(message *discordgo.MessageCreate) error

	// TODO : Ou tu mets les method pour handle les voicestateUdpate event ?
}

/*
point d'entree pour les sockets
ca definit les points d'entree pour ce program :
https://github.com/denverquane/amonguscapture

Domain dependencies:
	AmongUsEvent
*/
type ISocketEventAdapter interface {
	HandleConnectCode(msg string)
	HandleLobbyEvent(msg string, gameId string)
	HandleStateEvent(msg string, gameId string)
	HandlePlayerEvent(msg *PlayerEvent, gameId string)
	HandleDisconnection(gameId string)
}
