package framework_test

import (
	"testing"

	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework"
)

const GAMEID = "adasdasdasdasdgameid"
const PLAYERID = "asdasdljasdlsjdla654"
const PLAYERID2 = "sdasdlkjasd"
const PLAYERID3 = "sdasdlkjasd1"
const PLAYERID4 = "sdasdlkjasd2"
const PLAYERID5 = "sdasdlkjasd3"

func TestEntryPoint(t *testing.T) {
	redisRepo := framework.NewRedisRepository()
	redisRepo.AddGame(GAMEID)
	redisRepo.AddPlayer(GAMEID, entity.NewPlayer(PLAYERID, true, "", ""))
	redisRepo.AddPlayer(GAMEID, entity.NewPlayer(PLAYERID2, true, "", ""))
	redisRepo.AddPlayer(GAMEID, entity.NewPlayer(PLAYERID3, true, "", ""))
	redisRepo.AddPlayer(GAMEID, entity.NewPlayer(PLAYERID4, true, "", ""))
	redisRepo.AddPlayer(GAMEID, entity.NewPlayer(PLAYERID5, true, "", ""))

	player, err := redisRepo.GetPlayer(GAMEID, PLAYERID2)
	if err != nil {
		t.Errorf("Error getiiiiting player : %s\n", err)
	}
	if player.Id != PLAYERID2 {
		t.Errorf("Error getting player : Ids are not the same\n")
	}

	game, _ := redisRepo.GetGame(GAMEID)
	if len(game.Players) != 5 {
		t.Errorf("Not all players added to game, needed 5 got %v\n", len(game.Players))
	}

	redisRepo.DeletePlayer(GAMEID, PLAYERID)
	game, _ = redisRepo.GetGame(GAMEID)
	if len(game.Players) != 4 {
		t.Errorf("player not deleted\n")
	}
	_, found := game.Players[PLAYERID]
	if found {
		t.Errorf("Wrong Player deleted\n")
	}

	newPlayer := entity.NewPlayer(PLAYERID2, false, "red", "jeanbaptiste")
	newPlayer2 := entity.NewPlayer(PLAYERID3, true, "violet", "mozart")
	redisRepo.UpdatePlayer(GAMEID, newPlayer)
	redisRepo.UpdatePlayer(GAMEID, newPlayer2)
	game, _ = redisRepo.GetGame(GAMEID)

	newPlayerUpdate := game.Players[PLAYERID2]
	newPlayerUpdate2 := game.Players[PLAYERID3]
	if newPlayerUpdate.Name != "jeanbaptiste" || newPlayerUpdate.Color != "red" {
		t.Errorf("wrong updated player")
	}
	if newPlayerUpdate2.Name != "mozart" || newPlayerUpdate2.Color != "violet" {
		t.Errorf("wrong updated player")
	}

	newPlayerUpdate, err = redisRepo.GetPlayer(GAMEID, PLAYERID2)
	newPlayerUpdate2, err = redisRepo.GetPlayer(GAMEID, PLAYERID3)
	if newPlayerUpdate.Name != "jeanbaptiste" || newPlayerUpdate.Color != "red" {
		t.Errorf("wrong updated player")
	}
	if newPlayerUpdate2.Name != "mozart" || newPlayerUpdate2.Color != "violet" {
		t.Errorf("wrong updated player")
	}

	redisRepo.DeletePlayer(GAMEID, PLAYERID5)
	redisRepo.DeletePlayer(GAMEID, PLAYERID4)
	redisRepo.DeletePlayer(GAMEID, PLAYERID3)
	game, _ = redisRepo.GetGame(GAMEID)
	if len(game.Players) != 1 {
		t.Errorf("Should have only one players now\n")
	}
	_, found = game.Players[PLAYERID2]
	if !found {
		t.Errorf("PLAYER2 should not have been deleted")
	}

	redisRepo.DeleteGame(GAMEID)
	game, _ = redisRepo.GetGame(GAMEID)
	if len(game.Players) != 0 {
		t.Errorf("Game should not have players anymore\n")
	}
}
