package entity

/*
	Domain entities, all the repositories or framework
	must adapt their return type to those entities
*/

type Player struct {
	Id      string
	IsAlive bool
}

func NewPlayer(id string) *Player {
	return &Player{
		Id:      id,
		IsAlive: true,
	}
}

type Game struct {
	Id      string
	Players []*Player
}

func NewGame(id string) *Game {
	return &Game{
		Id: id,
	}
}
