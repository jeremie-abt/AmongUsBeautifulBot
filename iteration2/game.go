package bot

import (
	"math/rand"
	"time"
)

import (
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration2/logger"
)

type IGame interface {
	HandleAmongUsEvent(msg interface{}, evtType AuEventType)
	HandleLobbyEvent(event string)
	HandlePlayerEvent(event *AmongUsPlayerEvent)
	ResetGameState()
	BeginRound()
	BeginChattingState()
	GetGameFromCode(string) *Game
}

// TODO : fields to define
type Game struct {
	ChannelId     string
	DcUsers       DcUsers
	AuPlayers     *AuPlayers
	CurrentPlayer int
	AuCaptureCode string
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//	{"Action":4,"Name":"jejelaterr","IsDead":false,"Disconnected":false,"Color":7}
type AmongUsPlayerEvent struct {
	Action       AuEventPlayer `json:"Action"`
	Name         string        `json:Name`
	IsDead       bool          `json:"IsDead"`
	Disconnected bool          `json:"Disconnected"`
	Color        ColorPlayer   `json:"Color"`
}

type AuEventType string
type AuEventState string
type AuEventPlayer int
type AuPlayerColor int

const (
	LOBBY  AuEventType = "LOBBY"
	PLAYER             = "PLAYER"
	STATE              = "STATE"
)

const (
	LOBBYSTATE      AuEventState = "0"
	LOBBYBEGINROUND              = "2"
	LOBBYMENU                    = "3"
)

const (
	CHANGECOLOR AuEventPlayer = iota
	ISDEAD
)

func NewGame(channelId string, playerEntities PlayerEntities) IGame {
	return &Game{
		ChannelId:     channelId,
		DcUsers:       make(map[string]*DcUser),
		AuPlayers:     NewAuPlayers(playerEntities),
		CurrentPlayer: 0,
		//		AuCaptureCode: generateCode(10),
		AuCaptureCode: "aaaa", // TODO: replace by generatedCode
	}
}

func (game *Game) GetGameFromCode(code string) *Game {
	return nil
}

func (game *Game) ResetGameState() {
	for auplayer := range game.AuPlayers.Range() {
		auplayer.SetIsAlive(true)
		auplayer.Mute(false)
	}
	return
}

func (game *Game) BeginRound() {
	for auplayer := range game.AuPlayers.Range() {
		if auplayer.IsAlive() == true {
			auplayer.Mute(true)
		}
	}
}

func (game *Game) BeginChattingState() {
	for auplayer := range game.AuPlayers.Range() {
		if auplayer.IsAlive() == true {
			auplayer.Mute(false)
		}
	}
}

/*
	Socket.go is calling this function
	It's the entry point from the socket part of the code.
*/
func (game *Game) HandleAmongUsEvent(msg interface{}, evtType AuEventType) {
	if evtType == LOBBY {
		game.HandleLobbyEvent(msg.(string))
		return
	} else if evtType == PLAYER {
		return
	}
	log.Warnlog.Printf("HandleAmongUsEvent, unknown evtType : %s\n", evtType)
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func generateCode(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// -------- Socket handling
// 0 LOBYYY
// 1 debut de tour
// 2 fin de tour (speaking time)
// 3 menu
func (game *Game) HandleLobbyEvent(event string) {

	if event == "0" {
		//TODO: impl
		return
	} else if event == "1" {
		game.BeginRound()
		return
	} else if event == "2" {
		game.BeginChattingState()
		return
	} else if event == "3" {
		// TODO: impl
		return
	}
	return
}

// msg :
// 0 -> LOBBY
// 1 debut de tour je crois (jetais imposteur) / reprise dun tour egalement
// 2 -> Kill / je lai eu aussi en appuyant sur le bouton Wtfff
// 3 -> MENU

func (game *Game) HandleStateEvent(event AuEventState) {
	// TODO: Impl
	log.Debuglog.Printf("HandleStateEvent ...")
	if event == LOBBYSTATE {
		log.Debuglog.Printf("LOBBYSTATE entered ...")

	} else if event == LOBBYBEGINROUND {
		log.Debuglog.Printf("LOBBYBEGINROUND entered ...")
		game.BeginRound()
	} else if event == LOBBYMENU {
		log.Debuglog.Printf("LOBBYMENU entered ...")
	}
	log.Debuglog.Printf("HandleStateEvent ended ...")
	return
}

func (game *Game) HandlePlayerEvent(event *AmongUsPlayerEvent) {
	if event.Action == CHANGECOLOR {
		auPlayer := game.AuPlayers.GetAuPlayer(event.Name)
		if err := auPlayer.ChangeColor(event.Color); err != nil {
			log.Warnlog.Printf("could not change color : ", event.Color)
		}
		return
	} else if event.Action == ISDEAD {
		auPlayer := game.AuPlayers.GetAuPlayer(event.Name)
		auPlayer.SetIsAlive(false)
	}
}
