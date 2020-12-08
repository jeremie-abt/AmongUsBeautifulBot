package services_test

import (
	"testing"
)

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/services"
	"github.com/stretchr/testify/assert"
)

const GUILDID = "6544684644634686"
const PLAYERID = "3353545453534"
const CHANNELID = "65464686464"

func TestForwardMessage(t *testing.T) {
	// TODO Mock NewBotCommandHandler to make reel test
	var err error

	bot := services.NewDiscordBotAdapter(domain.NewBotCommandHandler())
	session, _ := discordgo.New()

	assert := assert.New(t)

	err = bot.ForwardMessage(session, generateDiscordMessage(".bau"))
	assert.EqualError(err, services.ErrWrongCommand)
	err = bot.ForwardMessage(session, generateDiscordMessage(".bau  "))
	assert.EqualError(err, services.ErrWrongCommand)
	err = bot.ForwardMessage(session, generateDiscordMessage(".bau fake"))
	assert.EqualError(err, services.ErrWrongCommand)

	err = bot.ForwardMessage(session, generateDiscordMessage(".bau start"))
	assert.NoError(err)
	err = bot.ForwardMessage(session, generateDiscordMessage(".bau end"))
	assert.NoError(err)
	err = bot.ForwardMessage(session, generateDiscordMessage(".bau   end"))
	assert.NoError(err)

	err = bot.ForwardMessage(session, generateVoiceMessage(true, true))
	if assert.Error(err) {
		assert.EqualError(err, "VoiceStateUpdate")
	}
}

func generateDiscordMessage(msg string) *discordgo.MessageCreate {
	return &discordgo.MessageCreate{
		Message: &discordgo.Message{
			ID:        GUILDID,
			ChannelID: CHANNELID,
			GuildID:   GUILDID,
			Content:   msg,
		},
	}
}

func generateVoiceMessage(mute bool, deaf bool) *discordgo.VoiceStateUpdate {
	return &discordgo.VoiceStateUpdate{
		&discordgo.VoiceState{
			UserID: "654646",
			Mute:   mute,
			Deaf:   deaf,
		},
	}
}
