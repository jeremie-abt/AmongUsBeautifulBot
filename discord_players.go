/// This package will handle all about
/// Discord players, Struct to maintain internal state
/// Or things like that

package main

import (
//"github.com/bwmarrin/discordgo"
)

/// TODO : This struct will probably evolve
/// IN order to use more realistics stuff
/// like discordgo id of discord player
/// etc ...
type DiscordPlayer struct {
	// TODO : Define ID properly with discordgo type
	Id              string
	AmongPlayerRole CustomRoleName

	isMute    bool
	isDeaf    bool
	channelID string
	// TODO: Reflechir au chan Id jsuis pas sur que ce soit
	// judicieux de le mettre la !
}

func NewDiscordPlayer(userId string) *DiscordPlayer {
	/// TODO : unmock, pour linstant je mock pour test
	/// la feature talkie
	return &DiscordPlayer{
		Id:              userId,
		AmongPlayerRole: Talkie,
	}
}

// TODO: Continuer de voir si c'est possible de trouver
// func GetAllPlayerOfdiscordChan(*discordgo.Channel) {
//
// }
