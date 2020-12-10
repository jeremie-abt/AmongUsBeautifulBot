package framework

import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"

const ErrAlreadyExist = "this Id already Exist"
const ErrNotExisting = "this Id is not existing"

type Repository interface {
	AddGame(id string) error
	GetGame(id string) (*entity.Game, error)
	DeleteGame(id string) error

	// Player
	AddPlayer(gameId string, player *entity.Player) error
	DeletePlayer(gameId string, playerId string) error
}
