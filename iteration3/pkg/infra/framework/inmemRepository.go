package framework

import (
	"fmt"

	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"
)

type inmemRepository struct {
	// Je prend direct un entity.Game
	// ca m'evite de faire des conversion de donne
	// mais normalement on devrait avoir un propre type
	// avec des methods qui convertisse les donnees entrante
	// et sortante
	game map[string]*entity.Game
}

func NewInMemRepository() Repository {
	return &inmemRepository{
		game: make(map[string]*entity.Game),
	}
}

func (i *inmemRepository) AddGame(id string) error {
	_, found := i.game[id]
	if found {
		return fmt.Errorf(ErrAlreadyExist)
	}
	i.game[id] = entity.NewGame(id, nil)
	return nil
}

func (i *inmemRepository) GetGame(id string) (*entity.Game, error) {
	_, found := i.game[id]
	if !found {
		return nil, fmt.Errorf(ErrNotExisting)
	}

	return i.game[id], nil
}

func (i *inmemRepository) DeleteGame(id string) error {
	_, found := i.game[id]
	if !found {
		return fmt.Errorf(ErrNotExisting)
	}
	delete(i.game, id)
	return nil
}

func (i *inmemRepository) AddPlayer(
	gameId string, player *entity.Player) error {
	game, found := i.game[gameId]
	if !found {
		return fmt.Errorf(ErrNotExisting)
	}
	// Ici je devrais convertir la data normalement
	game.Players[gameId] = player
	return nil
}

func (i *inmemRepository) GetPlayer(
	gameId string, playerId string) (*entity.Player, error) {

	return nil, nil
}
func (i *inmemRepository) UpdatePlayer(
	gameId string, player *entity.Player) error {
	panic("Not impl\n")
	return nil
}

func (i *inmemRepository) SetDeadPlayer(
	gameId string, playerId string) error {

	game, err := i.GetGame(gameId)
	if err != nil {
		return err
	}
	player, found := game.Players[playerId]
	if !found {
		return fmt.Errorf(ErrNotExisting)
	}
	player.SetDead()
	return nil
}

func (i *inmemRepository) DeletePlayer(gameId string, playerId string) error {
	panic("DeletePlayer not implemented for inmemRepo\n")
}
