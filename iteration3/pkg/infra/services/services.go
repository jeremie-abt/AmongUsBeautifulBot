package services

import "github.com/bwmarrin/discordgo"

type IbotCommandService interface {
	// TODO: Je laisse ces deux fonction mais je ne pense pas qu'elles
	// Doivent rester a terme ?
	StartGame(gameId string) error
	DeleteGame(gameId string) error

	HandleWrittenMessage(message *discordgo.MessageCreate) error

	// TODO : Ou tu mets les method pour handle les voicestateUdpate event ?
}
