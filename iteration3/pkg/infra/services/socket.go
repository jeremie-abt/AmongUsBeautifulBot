package services

import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain"
import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework"
import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"

type socketHandling struct {
	auEventHandler domain.IAmongUsEvent
	botCmdHandler  domain.IBotCommand
}

type SocketEvt string

const (
	PLAYERUPDATECOLOR SocketEvt = "PLAYERUPDATECOLOR"
	PLAYERUPDATENAME            = "PLAYERUPDATENAME"
	PLAYERDEAD                  = "PLAYERDEAD"
)

type PlayerEvent struct {
	Name   string                      `json:"name"`
	Action entity.AmongUsEvtPlayerType `json:"action"`
	Color  entity.AmongUsColor         `json:"color"`
}

func NewSocketAdapter() ISocketEventAdapter {

	redisRepo := framework.NewRedisRepository()
	discordVoip := framework.NewDiscordFramework("Nil")

	auEvtHandler := domain.NewAmongUsEvtHandler(discordVoip, redisRepo)
	botCmdHandler := domain.NewBotCommandHandler(redisRepo)
	return &socketHandling{
		auEventHandler: auEvtHandler,
		botCmdHandler:  botCmdHandler,
	}
}

func (s *socketHandling) HandleConnectCode(gameId string) {
	if !s.botCmdHandler.IsGameIdExisting(gameId) {
		// TODO: que faire lorsque je recois un code pour une game
		// qui nexiste pas ?
		// Throw un error pour les sockets ??
	}
}

// ???
func (s *socketHandling) HandleLobbyEvent(msg string, gameId string) {
	evt := transformLobbyEventForDomain(msg, gameId)

	if evt != nil {
		err := s.auEventHandler.HandleEvent(*evt)
		if err != nil {
			panic(err)
		}
	}
}

func (s *socketHandling) HandleStateEvent(msg string, gameId string) {
	evt := transformStateEventForDomain(msg, gameId)

	if evt != nil {
		err := s.auEventHandler.HandleEvent(*evt)
		if err != nil {
			panic(err)
		}
	}
}
func (s *socketHandling) HandlePlayerEvent(msg *PlayerEvent, gameId string) {

}

func (s *socketHandling) HandleDisconnection(gameId string) {

}

// Private methods

/*
msg.Action :
	0 -> connection au lobby.
	2 -> nom de la personne venant de mourir.
	3 -> changement de couleur.
	5 -> Disconnected.
	6 -> mort pour les votes.
*/
func transformPlayerEventForDomain(
	msg *PlayerEvent, gameId string) *entity.AmongUsEvent {

	evtReturned := &entity.AmongUsEvent{
		AttachedGameId: gameId,
		PlayerName:     msg.Name,
	}

	if msg.Action == "0" {
		//	evtReturned.Evttype = entity.GAMENEWPLAYER
		evtReturned.Evttype = entity.NOTIMPLEMENT
	} else if msg.Action == "2" {
		evtReturned.Evttype = entity.PLAYERDEAD
	} else if msg.Action == "3" {
		evtReturned.Evttype = entity.PLAYERUPDATECOLOR
	} else if msg.Action == "5" {
		evtReturned.Evttype = entity.NOTIMPLEMENT
	} else if msg.Action == "6" {
		evtReturned.Evttype = entity.PLAYERDEAD
	}
	return evtReturned
}

func transformLobbyEventForDomain(
	msg string, gameId string) *entity.AmongUsEvent {
	// TODO : Implement
	return &entity.AmongUsEvent{
		AttachedGameId: gameId,
		Evttype:        entity.NOTIMPLEMENT,
	}
}

/*
msg :
	0 -> LOBBY
	1 debut de tour
	3 -> Reunion vocale
*/
func transformStateEventForDomain(
	msg string, gameId string) *entity.AmongUsEvent {
	if msg == "0" {
		return &entity.AmongUsEvent{
			AttachedGameId: gameId,
			Evttype:        entity.GAMELOBBY,
		}
	} else if msg == "1" {
		return &entity.AmongUsEvent{
			AttachedGameId: gameId,
			Evttype:        entity.GAMELOBBY,
		}
	} else if msg == "3" {
		return &entity.AmongUsEvent{
			AttachedGameId: gameId,
			Evttype:        entity.GAMELOBBY,
		}
	}
	return nil
}
