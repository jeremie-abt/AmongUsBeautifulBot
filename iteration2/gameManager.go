package bot

type GameManager interface {
	AddGame(string)
	DeleteGame(string)
}

type GameManagerT struct {
	Games map[string]IGame
	Id    string
}

func NewGameManager(Id string) GameManager {
	return &GameManagerT{
		Games: make(map[string]IGame),
		Id:    Id,
	}
}

/*
	Cahier des charges :
		- Garder le track des differentes games.
		- etre capable de ressortir une game par rapport a un code
*/

func (gameman *GameManagerT) AddGame(gameId string) {
	game := NewGame(gameId, NewDcUsers())
	gameman.Games[gameId] = game
	return
}

func (gameman *GameManagerT) DeleteGame(gameId string) {
	return
}

func (gameman *GameManagerT) Range() <-chan IGame {
	gamemanagerChan := make(chan IGame)

	go func() {
		defer close(gamemanagerChan)
		for _, val := range gameman.Games {
			gamemanagerChan <- val
		}
	}()

	return gamemanagerChan
}
