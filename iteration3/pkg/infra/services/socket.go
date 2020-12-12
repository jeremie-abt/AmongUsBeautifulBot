package services

import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain"
import "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"

type socketHandling struct {
	auEventHandler domain.IAmongUsEvent
	botCmdHandler  domain.IBotCommand
}

type PlayerEvent struct {
	Name   string `json:"name"`
	Action string `json:"action"`
	Color  string `json:"color"`
}

func NewSocketHandling(
	auEvtHandler domain.IAmongUsEvent,
	botCmdHandler domain.IBotCommand) ISocketCaptureCodeEvent {
	return &socketHandling{
		auEventHandler: auEvtHandler,
		botCmdHandler:  botCmdHandler,
	}
}

func (s *socketHandling) HandleConnectCode(captureCode string) {
	if !s.botCmdHandler.IsGameIdExisting(captureCode) {
		// TODO: que faire lorsque je recois un code pour une game
		// qui nexiste pas ?
		// Throw un error pour les sockets ??
	}
}

// ???
func (s *socketHandling) HandleLobbyEvent(msg string, captureCode string) {
	evt := transformLobbyEventForDomain(msg, captureCode)

	if evt != nil {
		err := s.auEventHandler.HandleEvent(evt)
		if err != nil {
			panic(err)
		}
	}
}

func (s *socketHandling) HandleStateEvent(msg string, captureCode string) {
	evt := transformStateEventForDomain(msg, captureCode)

	if evt != nil {
		err := s.auEventHandler.HandleEvent(evt)
		if err != nil {
			panic(err)
		}
	}
}
func (s *socketHandling) HandlePlayerEvent(msg *PlayerEvent, captureCode string) {

}

func transformLobbyEventForDomain(
	msg string, captureCode string) *entity.AmongUsEvent {
	// TODO : Implement
	return &entity.AmongUsEvent{
		AttachedGameId: captureCode,
		Evttype:        entity.NOTIMPLEMENT,
	}
}

func (s *socketHandling) HandleDisconnection(captureCode string) {

}

/*
msg.Action :
	0 -> connection au lobby.
	2 -> nom de la personne venant de mourir.
	3 -> changement de couleur.
	5 -> Disconnected.
	6 -> mort pour les votes.
*/
func transformPlayerEventForDomain(
	msg *PlayerEvent, captureCode string) *entity.AmongUsEvent {

	evtReturned := &entity.AmongUsEvent{
		AttachedGameId: captureCode,
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

/*
msg :
	0 -> LOBBY
	1 debut de tour
	3 -> Reunion vocale
*/
func transformStateEventForDomain(
	msg string, captureCode string) *entity.AmongUsEvent {
	if msg == "0" {
		return &entity.AmongUsEvent{
			AttachedGameId: captureCode,
			Evttype:        entity.GAMELOBBY,
		}
	} else if msg == "1" {
		return &entity.AmongUsEvent{
			AttachedGameId: captureCode,
			Evttype:        entity.GAMELOBBY,
		}
	} else if msg == "3" {
		return &entity.AmongUsEvent{
			AttachedGameId: captureCode,
			Evttype:        entity.GAMELOBBY,
		}
	}
	return nil
}
