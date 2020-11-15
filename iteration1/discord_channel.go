/// this package should handle all the channel stuff
/// It's still to determine but what I have in mind :
///		Specify a chan which will create among us game
///		Wait for ten person to connect within
///		config rules and special rules
///		etc ...

package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func CreateNewChan(
	s *discordgo.Session, guildID string, name string,
) *discordgo.Channel {
	ret, err := s.GuildChannelCreate(
		guildID, name, discordgo.ChannelTypeGuildVoice)
	if err != nil {
		// TODO : Manage les erreurs
		panic("error creating chan\n")
	}
	return ret
}

func DeleteChan(s *discordgo.Session, chanId string) error {
	_, err := s.ChannelDelete(chanId)
	if err != nil {
		return fmt.Errorf("could not delete chan")
	}
	return nil
}

// TODO: A voir mais normalement j'en ai plus besoins
//type channelStruct struct {
//	Id string
//	discordChan *discordgo.Channel
//}
//
//
//func NewDiscordChanStruct(
//		Id string, dg *discordgo.Session) (*channelStruct, error) {
//
//	discordChan, err := dg.Channel(Id)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &channelStruct{
//		Id: Id,
//		discordChan: discordChan,
//	}, nil
//}
