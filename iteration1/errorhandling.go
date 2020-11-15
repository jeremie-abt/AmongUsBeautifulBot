package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// When we got an error while handling error
// I want to send them somewhere
// TODO : Probably send them in some db or some
// monitoring tools
func logFatalError(err error) {
	fmt.Printf("\n\nFatal ERROR : %+v\n\n", err)
}

func NotifyError(s *discordgo.Session, userId string, content error) {
	chanToSendError, err := s.UserChannelCreate(userId)

	fmt.Printf("Voici la creation du chanId : %+v\n", chanToSendError)
	if err != nil {
		logFatalError(err)
		return
	}

	chanId := chanToSendError.ID
	_, err = s.ChannelMessageSend(chanId, content.Error())

	if err != nil {
		logFatalError(err)
	}
}
