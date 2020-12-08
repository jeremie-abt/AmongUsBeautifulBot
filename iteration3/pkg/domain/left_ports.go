package domain

type IBotCommand interface {
	StartGame(id string) error
	StopGame(id string) error
	AddPlayer(gameId string, playerId string) error
	DeletePlayer(gameId string, playerId string) error
}
