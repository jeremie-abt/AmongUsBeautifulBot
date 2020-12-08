package services

import "github.com/bwmarrin/discordgo"

type IbotCommandService interface {
	ForwardMessage(*discordgo.Session, interface{}) error

	//	StartGame(*discordgo.Session, *discordgo.MessageCreate)
	//	DeleteGame(*discordgo.Session, *discordgo.MessageCreate)

	// TODO : Ou tu mets les method pour handle les voicestateUdpate event ?
}
