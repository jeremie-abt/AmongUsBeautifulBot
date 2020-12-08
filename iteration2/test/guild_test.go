package test

import (
	"testing"

	myBot "github.com/jeremie-abt/AmongUsBeautifulBot/iteration2"
)

func TestNewGuild(t *testing.T) {
	_ = myBot.NewGuild("123456789")

	t.Errorf("to impl\n")
}

func TestAddGame(t *testing.T) {
	t.Errorf("not impl\n")
}

func TestDeleteGame(t *testing.T) {
	t.Errorf("not impl\n")
}

// TODO : Trouver un moyen de faire ce test malgres la fonction portant le meme nom
//func TestGetGameFromCode(t *testing.T) {
//	t.Errorf("not impl\n")
//}
