package domain

import (
	"fmt"
	//	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework"
)

/*
	IbotCommand Ports implementation
*/
type BotCommandHandler struct {
	repo framework.Repository
}

func NewBotCommandHandler(repo framework.Repository) IBotCommand {
	return &BotCommandHandler{
		repo: repo,
	}
}

func (b *BotCommandHandler) StartGame(id string) error {
	err := b.repo.AddGame(id)
	if err != nil {
		return err
	}
	return nil
}

func (b *BotCommandHandler) StopGame(id string) error {
	err := b.repo.DeleteGame(id)
	if err != nil {
		return err
	}
	return nil
}
func (b *BotCommandHandler) IsGameIdExisting(gameId string) bool {
	return false
}

func (b *BotCommandHandler) AddPlayer(gameId string, playerId string) error {
	game, err := b.repo.GetGame(gameId)
	if err != nil {
		return err
	}
	if len(game.Players) >= 10 {
		return fmt.Errorf(ErrTooMuchPlayer)
	}
	// Add
	err = b.repo.AddPlayer(gameId, entity.NewPlayer(playerId, true, "", ""))
	return err
}

func (b *BotCommandHandler) DeletePlayer(gameId string, playerId string) error {
	game, err := b.repo.GetGame(gameId)
	if err != nil {
		return err
	}
	if len(game.Players) >= 10 {
		return fmt.Errorf(ErrTooMuchPlayer)
	}
	err = b.repo.DeletePlayer(gameId, playerId)
	return nil
}
