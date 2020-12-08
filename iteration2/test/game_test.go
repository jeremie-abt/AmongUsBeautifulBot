package test

import (
	"testing"

	myBot "github.com/jeremie-abt/AmongUsBeautifulBot/iteration2"
)

// Player :
// msg example :
//	{"Action":4,"Name":"jejelaterr","IsDead":false,"Disconnected":false,"Color":7}
// {"Action":5,"Name":"I","IsDead":false,"Disconnected":true,"Color":5}
// Action :
// 0 -> connection au lobby
// 1 -> ???
// 2 -> mourir -> Tu as le name de la person qui vient de mourir
// 3 -> changement de couleur
// 	5 -> Disconnected (a confirmer)
// 6 -> Je crois que c quand une personne est ejecte aux votes

// dans un premier temps, je vais simuler les events (en declarant moi meme tel
// event == tel action), pour pouvoir faire ma logique de test et ensuite
// jirai voir vraiment quel event = quel event pour poffiner le tout !

var game *myBot.Game

const CHANID = "54546864684"

func TestResetGame(t *testing.T) {
	// TODO
	return
}

func TestInitMock(t *testing.T) {
	funcMap := make(map[string]func(...interface{}))
	funcMap["HandleAmongUsEvent"] = func(vars ...interface{}) {
		return
	}
}

func TestHandleLobbyEvent(t *testing.T) {
	// Voir commentaire de TestHandleStateEvent
	return
}

func TestHandleAmongUsEvent(t *testing.T) {
	// pareil, je ne peux pas vraiment tracker les call de fonctions
}

func TestHandleStateEvent(t *testing.T) {
	// Voila un exemple exacte de pourquoi cet idee de mock n'est pas une
	// bonne idee
	// En tout le fait de tous mettre dans un meme fichier n'est pas viable
	// Pas testable en l'etat (game repose sur amongus)
	//	functions := make(map[string]func(...interface{}) interface{})
	//	functions["HandleStateEvent"] = myBot.HandleStateEvent
	//	mock := NewMockObj(functions)
	//
	//	mock.AssertFuncCalledWith("BeginRound")
	//	mock.HandleStateEvent(myBot.LOBBYBEGINROUND)
	//	if err := mock.VerifyFuncCall(); err != nil {
	//		t.Errorf("should not be err : %s\n", err)
	//	}
	return
}

func TestHandlePlayerEvent(t *testing.T) {
	// Pas testable en l'etat (game repose sur amongus)
	return
}

func TestBeginRound(t *testing.T) {
	// je ne pense pas que ce soit vraiment testable
	// car il n'y a pas vraiment de logique, c'est uniquement
	// une fonction qui se contente d'appeler les layers de code
	// du dessous
	return
}

func TestBeginChattingState(t *testing.T) {
	// je ne pense pas que ce soit vraiment testable
	// car il n'y a pas vraiment de logique, c'est uniquement
	// une fonction qui se contente d'appeler les layers de code
	// du dessous
	return
}
