package framework

import "context"
import "github.com/go-redis/redis"

import (
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"
	"os"
	"strings"
)

const GAMEID = "54654646464"
const PLAYERID = "5465464646askdad4"
const PREFIXGAME = "gameId"
const PREFIXPLAYER = "playerId"

type redisRepo struct {
	client  *redis.Client
	context context.Context
}

func NewRedisRepository() Repository {
	redisPort := os.Getenv("REDIS_HOST")
	if redisPort == "" {
		// Default port
		redisPort = "6700"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + redisPort, // use default Addr
		Password: "",                       // no password set
		DB:       0,                        // use default DB
	})
	// TODO: Gerer les envs prod et dev rapidement
	return &redisRepo{
		client:  rdb,
		context: context.Background(),
	}
}

/*
	Mes types Redis :
		Game:
			- id string
			- player string (objet player separe par virgule)
		Player:
			- id string (PREFIXGAME + PREFIXPLAYER + id of player)
			- color int
			- isAlive strign("false" || "true")
*/

func (rd *redisRepo) AddGame(id string) error {
	rd.client.HSet(getGameKey(id), "id", id)
	//  pas utile de set une key avec juste une key val nil
	// https://stackoverflow.com/questions/13817865/
	// redis-nil-or-empty-list-or-set
	return nil
}

func (rd *redisRepo) GetGame(id string) (*entity.Game, error) {
	strResults, err := rd.client.HGetAll(getGameKey(id)).Result()
	if err != nil {
		return nil, err
	}
	players, found := strResults["players"]
	if !found {
		players = ""
	}
	playersId := getPlayerIdsFromString(players)
	gamePlayers := make(map[string]*entity.Player)
	for _, val := range playersId {
		// TODO
		player, _ := rd.GetPlayer(id, val)
		gamePlayers[player.Id] = player
	}
	return entity.NewGame(id, gamePlayers), nil
}

func (rd *redisRepo) DeleteGame(id string) error {

	// Pas fou je devrais manipuler des structs interne mais bon ...
	game, err := rd.GetGame(id)
	for _, player := range game.Players {
		rd.DeletePlayer(id, player.Id)
	}
	keys, err := rd.client.HKeys(getGameKey(id)).Result()
	if err != nil {
		return err
	}
	for _, key := range keys {
		rd.client.HDel(getGameKey(id), key)
	}
	return nil
}

func (rd *redisRepo) AddPlayer(gameId string, player *entity.Player) error {
	var newPlayersVal string

	updatePlayer(rd.client, gameId, player)

	// Udpate game.players variable
	game, err := rd.GetGame(gameId)
	if err != nil {
		return err
	}
	strPlayers := getPlayerIdsFromMap(game.Players)
	if strPlayers == "" {
		newPlayersVal = player.Id
	} else {
		newPlayersVal = strPlayers + "," + player.Id
	}
	rd.client.HSet(getGameKey(gameId), "players", newPlayersVal)
	return nil
}

func (rd *redisRepo) UpdatePlayer(gameId string, player *entity.Player) error {
	// Get le player et lupdate (je nai pas besoins de faire dupdate sur la game)

	updatePlayer(rd.client, gameId, player)
	return nil
}

func (rd *redisRepo) GetPlayer(
	gameId string, playerId string) (*entity.Player, error) {
	player, err := rd.client.HGetAll(getPlayerKey(gameId, playerId)).Result()

	if err != nil {
		return nil, err
	}
	return entity.NewPlayer(
		player["id"], stringToBool(player["isAlive"]),
		player["color"], player["name"]), nil
}

func (rd *redisRepo) SetDeadPlayer(gameId string, playerId string) error {
	return nil
}

func (rd *redisRepo) DeletePlayer(gameId string, playerId string) error {
	rd.client.HDel(getPlayerKey(gameId, playerId))

	game, err := rd.GetGame(gameId)
	if err != nil {
		return err
	}

	// udpate game.players variable
	delete(game.Players, playerId)
	strPlayers := getPlayerIdsFromMap(game.Players)
	rd.client.HSet(getGameKey(gameId), "players", strPlayers)
	return nil
}

/// --- Private methods
func updatePlayer(rd *redis.Client, gameId string, player *entity.Player) {

	playerKey := getPlayerKey(gameId, player.Id)
	rd.HSet(playerKey, "isAlive", boolToString(player.IsAlive()))
	rd.HSet(playerKey, "id", player.Id)
	rd.HSet(playerKey, "name", player.Name)
	rd.HSet(playerKey, "color", player.Color)
}

func getPlayerIdsFromString(players string) []string {
	if players == "" {
		return nil
	}
	return strings.Split(players, ",")
}

func getPlayerIdsFromMap(players map[string]*entity.Player) string {
	var ret string
	ret = ""
	for key, _ := range players {
		ret = ret + key + ","
	}
	if len(ret) > 0 {
		// enlever la dernier virgule
		return ret[0 : len(ret)-1]
	}
	return ret
}

func getPlayerKey(gameId string, playerId string) string {
	return PREFIXPLAYER + gameId + playerId
}
func getGameKey(gameId string) string {
	return PREFIXGAME + gameId
}
func stringToBool(boolVal string) bool {
	if boolVal == "true" {
		return true
	} else if boolVal == "false" {
		return false
	}
	panic("stringToBool : Not correct string\n")
}
func boolToString(val bool) string {
	if val == true {
		return "true"
	} else {
		return "false"
	}
}
