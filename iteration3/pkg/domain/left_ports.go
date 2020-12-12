package domain

import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"

type IBotCommand interface {
	StartGame(id string) error
	StopGame(id string) error
	AddPlayer(gameId string, playerId string) error
	DeletePlayer(gameId string, playerId string) error
	IsGameIdExisting(gameId string) bool
}

// Logic for managing among us event
type IAmongUsEvent interface {
	HandleEvent(*entity.AmongUsEvent) error
}
