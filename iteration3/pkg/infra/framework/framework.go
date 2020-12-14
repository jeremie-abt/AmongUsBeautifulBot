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
	GetPlayer(gameId string, playerId string) (*entity.Player, error)
	UpdatePlayer(gameId string, player *entity.Player) error
	SetDeadPlayer(gameId string, playerId string) error
	DeletePlayer(gameId string, playerId string) error
}

type VoipServer interface {
	MuteAll(game *entity.Game) error
	UnMuteAll(game *entity.Game) error
	UpdateColor(game *entity.Game, newPlayer *entity.Player) error
	UpdateName(game *entity.Game, newPlayer *entity.Player) error
}
