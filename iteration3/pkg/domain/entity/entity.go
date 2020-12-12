package entity

/*
	Domain entities, all the repositories or framework
	must adapt their return type to those entities
*/

type AmongUsEvtType string

// Attention ca ne doit pas etre lie a discordgo
// je pense que c bien davoir juste ces deux events
// et ne pas implementer tous ceux de discordgo
const (
	PLAYERUPDATECOLOR AmongUsEvtType = "PLAYERUPDATECOLOR"
	PLAYERUPDATENAME                 = "PLAYERUPDATENAME"
	PLAYERDEAD                       = "PLAYERDEAD"
	GAMENEWPLAYER                    = "GAMENEWPLAYER"
	GAMELOBBY                        = "GAMELOBBY"
	GAMEBREAK                        = "GAMEBREAK"
	GAMEBEGINROUND                   = "GAMEBEGINROUND"
	NOTIMPLEMENT                     = "NOTIMPL"
)

const ()

/*
	Le but de se format est de gerer tous les events socket
	quite a avoir des variable non utilises
*/
type AmongUsEvent struct {
	AttachedGame   *Game
	AttachedGameId string
	Evttype        AmongUsEvtType
	PlayerName     string
	PlayerColor    string
}

type IPlayer interface {
	SetAlive()
	SetDead()
	IsAlive() bool
}

type Player struct {
	Id      string
	Color   string
	Name    string
	isAlive bool
}

func (p *Player) SetAlive() {
	p.isAlive = true
}

func (p *Player) SetDead() {
	p.isAlive = false
}

func (p *Player) IsAlive() bool {
	return p.isAlive
}

func NewPlayer(id string) *Player {
	return &Player{
		Id:      id,
		isAlive: true,
	}
}

type Game struct {
	Id      string
	Players map[string]*Player
}

func NewGame(id string) *Game {
	return &Game{
		Id: id,
	}
}
