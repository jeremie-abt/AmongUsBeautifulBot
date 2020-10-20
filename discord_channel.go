/// this package should handle all the channel stuff
/// It's still to determine but what I have in mind :
///		Specify a chan which will create among us game
///		Wait for ten person to connect within
///		config rules and special rules
///		etc ...

package main


import (
	"github.com/bwmarrin/discordgo"
)


//type DiscordChannel struct {
//	name string
//}
//
//
//func NewDiscordChan() (*DiscordChannel) {
//	return &{
//		""
//	}
//}
//
//
///// Get all players within chan
//func GetAllPlayerFromChan()


type channelStruct struct {
	Id string
	discordChan *discordgo.Channel
}


func NewDiscordChanStruct(
		Id string, dg *discordgo.Session) (*channelStruct, error) {

	discordChan, err := dg.Channel(Id)

	if err != nil {
		return nil, err
	}
	
	return &channelStruct{
		Id: Id,
		discordChan: discordChan,
	}, nil
}
