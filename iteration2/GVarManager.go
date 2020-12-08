package bot

import (
	"fmt"
)

type GVarManagerT map[string]GameManager

func NewGVarManager() GVarManagerT {
	return make(GVarManagerT)
}

/*
	Typiquement le val.(*GameManagerT) c'est pas ouf
	mais comme jai plusieurs Range method, je ne peux pas
	la mettre dans l'interface a cause du mock
*/
func (gvar GVarManagerT) GetGameFromCode(code string) IGame {
	for _, val := range gvar {
		for game := range val.(*GameManagerT).Range() {
			if ret := game.(*Game).GetGameFromCode(code); ret != nil {
				return ret
			}
		}
	}
	return nil
}

func (gvar GVarManagerT) AddGvarManager(Id string, gamemanager GameManager) error {
	if _, found := gvar[Id]; found {
		return fmt.Errorf("Key %s already in memory\n", Id)
	}

	gvar[Id] = gamemanager
	return nil
}
