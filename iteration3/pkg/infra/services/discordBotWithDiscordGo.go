package services

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain"
	"strings"
)

const ErrWrongCommand = "Wrong command, Couldn't parse It\n"
const InternalError = "Internal error !\n"

type discordBotAdapter struct {
	botCommandHandler domain.IBotCommand
}

// Dans la conf il faisait en sorte que ca prenne une interface en params
func NewDiscordBotAdapter(botCmdHandler domain.IBotCommand) IDiscordBotWithDiscordGo {
	return &discordBotAdapter{
		botCommandHandler: botCmdHandler,
	}
}

// appel la fonction de parse si l'interface est bonne, cette fonction
// est donc la liaison entre discord go et ce module
// Probablement qu'il serait interessant de rajouter de l'abstraction ici
// pour eviter d'etre colle a discordgo
// Surtout quand tu prends une lib random en debut de projet car tu as juste envie
// de commencer a coder.
func (s *discordBotAdapter) HandleWrittenMessage(
	sess *discordgo.Session,
	msg *discordgo.MessageCreate) {

	s.parseWrittenCommand(msg)
}

func (s *discordBotAdapter) StartGame(gameId string) error {
	// On entre cette fonction sans avoir fais de validation, pour ca qu'une fois de plus
	// je ne pense pas qu'elle doit etre publique

	err := s.botCommandHandler.StartGame(gameId)
	if err != nil {
		return err
	}
	return nil
}

func (s *discordBotAdapter) DeleteGame(gameId string) error {
	// On entre cette fonction sans avoir fais de validation, pour ca qu'une fois de plus
	// je ne pense pas qu'elle doit etre publique

	err := s.botCommandHandler.StopGame(gameId)

	if err != nil {
		return err
	}
	return nil
}

// Drive logic de dispatching de command
// ca doit etre cette fonction qui appel StartGame etc ...
// Pour ca que je ne pense pas que ces methods doivent etre public
func (s *discordBotAdapter) parseWrittenCommand(msg *discordgo.MessageCreate) error {
	/*
		Rules for command :
			Should begin with .bau
			The following should be in the defined commands
	*/

	body := msg.Message.Content
	commandLists := map[string]func(string) error{
		"start": s.StartGame,
		"end":   s.DeleteGame,
	}
	body = strings.ToLower(body)
	if len(body) <= 5 {
		return fmt.Errorf(ErrWrongCommand)
	}
	if !strings.HasPrefix(body, ".bau ") {
		return fmt.Errorf(ErrWrongCommand)
	}
	command := getCommand(body)
	handlingFunc, found := commandLists[command]
	if !found {
		return fmt.Errorf(ErrWrongCommand)
	}
	return handlingFunc(msg.Message.ChannelID)
}

func getCommand(message string) string {
	return strings.TrimSpace(message[5:])
}
