package domain

import (
	"fmt"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework"
)

type AmongUsEvtHandler struct {
	voipServerFramework framework.VoipServer
	repo                framework.Repository
}

func NewAmongUsEvtHandler(
	vp framework.VoipServer,
	repo framework.Repository) *AmongUsEvtHandler {
	return &AmongUsEvtHandler{
		voipServerFramework: vp,
		repo:                repo,
	}
}

/*
	Gere une grosse partie de la logique.
	Possibilite de split en plusieurs fonctions qui vont chacune
	organise leurs gestion d'un event en particulier
*/
func (a *AmongUsEvtHandler) HandleEvent(evt entity.AmongUsEvent) error {
	// Parser levent
	if evt.AttachedGame == nil {
		game, err := a.getGame(&evt)
		if err != nil {
			return nil
		}
		evt.AttachedGame = game
	}
	if evt.Evttype == entity.GAMELOBBY ||
		evt.Evttype == entity.GAMEBREAK ||
		evt.Evttype == entity.GAMEBEGINROUND {
		return a.handleGameStateChangeEvent(&evt)
	}
	if evt.Evttype == entity.PLAYERUPDATECOLOR ||
		evt.Evttype == entity.PLAYERUPDATENAME ||
		evt.Evttype == entity.PLAYERDEAD {

		return a.handlePlayerEvent(&evt)
	}

	return nil
}

func (a *AmongUsEvtHandler) handlePlayerEvent(evt *entity.AmongUsEvent) error {
	var player *entity.Player

	game := evt.AttachedGame
	player = newPlayerEntity(evt.PlayerName, evt.PlayerColor)

	if evt.Evttype == entity.PLAYERUPDATECOLOR {
		err := a.repo.UpdatePlayer(evt.AttachedGameId, player)
		if err != nil {
			return err
		}
		a.voipServerFramework.UpdateColor(evt.AttachedGame, player)
	} else if evt.Evttype == entity.PLAYERUPDATENAME {
		err := a.repo.UpdatePlayer(evt.AttachedGameId, player)
		if err != nil {
			return err
		}
		a.voipServerFramework.UpdateName(evt.AttachedGame, player)

	} else if evt.Evttype == entity.PLAYERDEAD {
		// TODO: Comment chopper le playerId de manier clean
		err := a.repo.SetDeadPlayer(game.Id, "asdasdasdasd")
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AmongUsEvtHandler) handleGameStateChangeEvent(evt *entity.AmongUsEvent) error {
	// GetGame
	game := evt.AttachedGame

	if evt.Evttype == entity.GAMELOBBY {
		for _, player := range game.Players {
			player.SetAlive()
		}
		a.voipServerFramework.UnMuteAll(game)
	}
	if evt.Evttype == entity.GAMEBEGINROUND {
		a.voipServerFramework.MuteAll(game)
	}
	if evt.Evttype == entity.GAMEBREAK {
		a.voipServerFramework.UnMuteAll(game)
	}
	return nil
}

func (a *AmongUsEvtHandler) getGame(evt *entity.AmongUsEvent) (*entity.Game, error) {
	if evt.AttachedGame != nil {
		return evt.AttachedGame, nil
	}
	return a.repo.GetGame(evt.AttachedGameId)
}

func getPlayer(playerId string, game *entity.Game) (*entity.Player, error) {
	player, found := game.Players[playerId]

	if !found {
		return nil, fmt.Errorf(ErrNotExist)
	}
	return player, nil
}

/*
Je fais cette fonction ici pour la raison suivante :
	Jai besoins de linstancier dune certaine facon
	et je ne pense pas que ce soit bien dexposer cette
	method dans les domaines direct
	A voir
*/
func newPlayerEntity(name string, color string) *entity.Player {
	return &entity.Player{
		Name:  name,
		Color: color,
	}
}
