package test

import (
	myBot "github.com/jeremie-abt/AmongUsBeautifulBot/iteration2"
	"testing"
)

var functions map[string]func(...interface{}) interface{}

// --------- Test on AuPlayers
func TestAddAuPlayer(t *testing.T) {
	functions = map[string]func(...interface{}) interface{}{
		"GetEntityId": func(args ...interface{}) interface{} {
			if args[0] == "testReturnId" {
				return "testReturnedExistingId"
			}
			return ""
		},
	}
	mockobj := NewMockObj(functions)
	var auPlayers = myBot.NewAuPlayers(mockobj)

	auPlayers.AddAuPlayer("jeremie")
	auPlayers.AddAuPlayer("jeremie")
	auPlayers.AddAuPlayer("jeremie")
	if len(auPlayers.AuPlayers) != 1 {
		t.Errorf("len(auPlayers) should be equal to 1")
	}
	auPlayers.AddAuPlayer("jeremie1")
	auPlayers.AddAuPlayer("jeremie2")
	auPlayers.AddAuPlayer("jeremie3")
	auPlayers.AddAuPlayer("jeremie4")
	auPlayers.AddAuPlayer("jeremie5")
	auPlayers.AddAuPlayer("jeremie6")
	auPlayers.AddAuPlayer("jeremie7")
	auPlayers.AddAuPlayer("jeremie8")
	auPlayers.AddAuPlayer("jeremie9")
	auPlayers.AddAuPlayer("jeremie10")
	auPlayers.AddAuPlayer("jeremie11")
	auPlayers.AddAuPlayer("jeremie12")
	auPlayers.AddAuPlayer("jeremie13")
	if len(auPlayers.AuPlayers) != 10 {
		t.Errorf("len(auPlayers) should be equal to 10")
	}

	mockobj = NewMockObj(functions)
	auPlayers = myBot.NewAuPlayers(mockobj)

	auPlayers.AddAuPlayer("testReturnId")

	if auPlayers.GetAuPlayer("testReturnId").LinkedTo != "testReturnedExistingId" {
		t.Errorf("Should have been linked to : %s\n", "testReturnedExistingId")
	}
}

func TestGetAuPlayer(t *testing.T) {
	mockobj := NewMockObj(functions)
	var auPlayers = myBot.NewAuPlayers(mockobj)

	auPlayers.GetAuPlayer("unknown")

	auPlayers.AddAuPlayer("jeremie11")
	auPlayers.AddAuPlayer("julie")
	auPlayers.AddAuPlayer("jeremie11")
	auPlayers.AddAuPlayer("jeremie11")
	if auPlayers.GetAuPlayer("unknown") != nil {
		t.Errorf("Unknown should be nil")
	}
	auplayer := auPlayers.GetAuPlayer("jeremie11")
	if auplayer == nil {
		t.Errorf("Unknown jeremie11")
	}
	if auplayer.Name != "jeremie11" {
		t.Errorf("bad name %s - should be jeremie11\n", auplayer.Name)
	}

}

func TestUpdateAuPlayer(t *testing.T) {
	mockobj := NewMockObj(functions)
	var auPlayers = myBot.NewAuPlayers(mockobj)

	auPlayers.DeleteAuPlayer("sadasdasd")
	auPlayers.AddAuPlayer("julie")
	auPlayers.AddAuPlayer("jeremie12")
	auPlayers.AddAuPlayer("jeremie13")
	auPlayers.UpdateAuPlayer("jeremie12", "jeremie154")
	if len(auPlayers.AuPlayers) != 3 {
		t.Errorf("auPlayers len : %d, should be 3", len(auPlayers.AuPlayers))
	}
	if ret := auPlayers.GetAuPlayer("jeremie12"); ret != nil {
		t.Errorf("jeremie12 should not exist anymore")
	}
	if ret := auPlayers.GetAuPlayer("jeremie154"); ret == nil {
		t.Errorf("jeremie154 should exist anymore")
	}
}

func TestDeleteAuPlayer(t *testing.T) {
	mockobj := NewMockObj(functions)
	var auPlayers = myBot.NewAuPlayers(mockobj)

	auPlayers.DeleteAuPlayer("sadasdasd")
	auPlayers.AddAuPlayer("julie")
	auPlayers.AddAuPlayer("jeremie12")
	auPlayers.AddAuPlayer("jeremie13")
	auPlayers.DeleteAuPlayer("jeremie12")
	if len(auPlayers.AuPlayers) != 2 {
		t.Errorf("auPlayers len : %d, should be 2", len(auPlayers.AuPlayers))
	}
}

func TestChangeColor(t *testing.T) {
	t.Errorf("to impl\n")
}
