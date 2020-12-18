package framework

import (
	"github.com/bwmarrin/discordgo"
)

import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"

type DiscordFramework struct {
	sess *discordgo.Session
}

func NewDiscordFramework(botToken string) VoipServer {
	sess, _ := discordgo.New()
	return &DiscordFramework{
		sess: sess,
	}
}

// Voir s'il faut en changer
const BUCKETID = "6546asd46a4sd6a4sdwkjqwhd HSADGDS"

func (d *DiscordFramework) MuteAll(game *entity.Game) error {
	return nil
}

func (d *DiscordFramework) UnMuteAll(game *entity.Game) error {
	return nil
}

func (d *DiscordFramework) muteDiscord(game *entity.Game) {
	for _, player := range game.Players {
		d.sess.RequestWithBucketID(
			"patch",
			discordgo.EndpointGuildMember(game.Id, player.Id),
			struct {
				Mute bool
			}{
				Mute: true,
			}, BUCKETID)
	}
}

func (d *DiscordFramework) UpdateColor(
	game *entity.Game, newPlayer *entity.Player) error {

	return nil
}

func (d *DiscordFramework) UpdateName(
	game *entity.Game, newPlayer *entity.Player) error {

	return nil
}
