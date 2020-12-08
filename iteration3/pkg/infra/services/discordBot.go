package services

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain"
	"reflect"
	"strings"
)

const ErrWrongCommand = "Wrong command, Couldn't parse It\n"

type discordBotAdapter struct {
	botCommandHandler domain.IBotCommand
}

// Dans la conf il faisait en sorte que ca prenne une interface en params
func NewDiscordBotAdapter(botCmdHandler domain.IBotCommand) IbotCommandService {
	return &discordBotAdapter{
		botCommandHandler: botCmdHandler,
	}
}

// pas utile je pense car depuis discordgo je vais register
// une fonction par event
// Donc inutile
func (s *discordBotAdapter) ForwardMessage(sess *discordgo.Session, i interface{}) error {
	InterfaceName := reflect.TypeOf(i).Elem().Name()
	if InterfaceName == "MessageCreate" {
		return parseCommand(i.(*discordgo.MessageCreate).Message.Content)
	}
	return fmt.Errorf(InterfaceName)
}

func parseCommand(body string) error {
	/*
		Rules for command :
			Should begin with .bau
			The following should be in the defined commands
	*/

	commandLists := map[string]bool{
		"start": true,
		"end":   true,
	}
	if len(body) <= 5 {
		return fmt.Errorf(ErrWrongCommand)
	}
	if !strings.HasPrefix(body, ".bau ") {
		return fmt.Errorf(ErrWrongCommand)
	}
	command := strings.TrimSpace(body[5:])
	if _, found := commandLists[command]; !found {
		return fmt.Errorf(ErrWrongCommand)
	}
	return nil
}
