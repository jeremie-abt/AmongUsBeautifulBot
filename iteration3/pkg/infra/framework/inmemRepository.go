package framework

import (
	"fmt"
)
import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"

type inmemRepository struct {
	// Je prend direct un entity.Game
	// ca m'evite de faire des conversion de donne
	// mais normalement on devrait avoir un propre type
	// avec des methods qui convertisse les donnees entrante
	// et sortante
	game map[string]*entity.Game
}

func (i *inmemRepository) AddGame(id string) error {
	_, found := i.game[id]
	if found {
		return fmt.Errorf(ErrAlreadyExist)
	}
	i.game[id] = entity.NewGame(id)
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

func (i *inmemRepository) AddPlayer(gameId string, player *entity.Player) error {
	game, found := i.game[gameId]
	if !found {
		return fmt.Errorf(ErrNotExisting)
	}
	// Ici je devrais convertir la data normalement
	game.Players = append(game.Players, player)
	return nil
}
