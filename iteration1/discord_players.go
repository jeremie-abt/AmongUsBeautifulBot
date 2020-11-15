/// This package will handle all about
/// Discord players, Struct to maintain internal state
/// Or things like that

package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)


const RATE_LIMIT_BUCKET = "RateLimitBucket"

/// TODO : This struct will probably evolve
/// IN order to use more realistics stuff
/// like discordgo id of discord player
/// etc ...
type DiscordPlayer struct {
	// TODO : Define ID properly with discordgo type
	Id					string
	Name				string
	AmongPlayerRole 	CustomRoleName
	ChannelID 			string
	IsAlive				bool

	isMute    			bool
	isDeaf    			bool
	AmongUsLinkedName	string
}

func NewDiscordPlayer(userId string, name string) *DiscordPlayer {
	/// TODO : unmock, pour linstant je mock pour test
	/// la feature talkie
	return &DiscordPlayer{
		Id:             	userId,
		Name:				name,
		AmongPlayerRole:	Talkie,
		AmongUsLinkedName:	"",
		IsAlive:			true,
	}
}

type test struct {
	Mute bool `json:"mute"`
}

// test et faire le unmute
func (player *DiscordPlayer) Mute(guildId string, mute bool) {
	/*
		Mute discord player
	*/

	if player.IsAlive == false {
		return
	}
	data := struct{
		Mute bool `json:"mute"`
	}{mute}
	_, err := G_Gvm.Sess.RequestWithBucketID("PATCH", discordgo.EndpointGuildMember(guildId, player.Id), data, RATE_LIMIT_BUCKET)

	if err != nil {
		// TODO log
		fmt.Printf("Error : %s\n", err)
		return
	}
}
