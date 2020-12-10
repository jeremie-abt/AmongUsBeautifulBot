package services_test

import (
	"testing"
)

import (
	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/mocks"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/services"
	"github.com/stretchr/testify/assert"
)

const GUILDID = "6544684644634686"
const PLAYERID = "3353545453534"
const CHANNELID = "65464686464"

func TestForwardMessage(t *testing.T) {
	// TODO Mock NewBotCommandHandler to make reel test
	var err error

	assert := assert.New(t)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Domain bot logic interface mock
	mock := mocks.NewMockIBotCommand(mockCtrl)
	mock.EXPECT().StartGame(CHANNELID).Return(nil).Times(3)
	mock.EXPECT().StopGame(CHANNELID).Return(nil).Times(2)

	bot := services.NewDiscordBotAdapter(mock)
	bot.HandleWrittenMessage(generateDiscordMessage(".bau start"))
	bot.HandleWrittenMessage(generateDiscordMessage(".bau start  "))
	bot.HandleWrittenMessage(generateDiscordMessage(".bau   START  "))

	bot.HandleWrittenMessage(generateDiscordMessage(".bau   end  "))
	bot.HandleWrittenMessage(generateDiscordMessage(".bau   End  "))

	err = bot.HandleWrittenMessage(generateDiscordMessage(".bau"))
	assert.EqualError(err, services.ErrWrongCommand)
	err = bot.HandleWrittenMessage(generateDiscordMessage(".bau      "))
	assert.EqualError(err, services.ErrWrongCommand)
	err = bot.HandleWrittenMessage(generateDiscordMessage(".bau   asdadsa   "))
	assert.EqualError(err, services.ErrWrongCommand)

	err = bot.HandleWrittenMessage(generateDiscordMessage(".bau"))
	assert.EqualError(err, services.ErrWrongCommand)
	err = bot.HandleWrittenMessage(generateDiscordMessage(".bau      "))
	assert.EqualError(err, services.ErrWrongCommand)
	err = bot.HandleWrittenMessage(generateDiscordMessage(".bau   asdadsa   "))
	assert.EqualError(err, services.ErrWrongCommand)
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
