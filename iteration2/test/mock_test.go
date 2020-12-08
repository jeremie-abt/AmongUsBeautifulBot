package test

import (
	"testing"
	"time"
)

import (
	mybot "github.com/jeremie-abt/AmongUsBeautifulBot/iteration2"
	//	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration2/logger"
)

/*
	Test for the VerifyFuncCall func
*/
func TestVerifyFuncCall(t *testing.T) {
	mockobj := NewMockObj(nil)
	if err := mockobj.VerifyFuncCall(); err != nil {
		t.Errorf("verify funcCall should not throw err : %s !\n", err)
	}
	mockobj.AssertFuncCalledWith("HandlePlayerEvent", &mybot.AmongUsPlayerEvent{
		Action: 5,
	})
	mockobj.AssertFuncCalledWith("GetEntityId", "6546654648")
	mockobj.AssertFuncCalledWith("Mute", false)
	mockobj.AssertFuncCalledWith("Mute", true)
	mockobj.AssertFuncCalledWith("RemoveEntity", "jjeeemss")
	mockobj.AssertFuncCalledWith("HandleAmongUsEvent", "jjems", "jjems")

	mockobj.HandlePlayerEvent(&mybot.AmongUsPlayerEvent{
		Action: 5,
	})
	mockobj.GetEntityId("asdasda")
	mockobj.GetEntityId("asdasda")
	mockobj.GetEntityId("6546654648")
	mockobj.GetEntityId("654665464")
	mockobj.GetEntityId("6546654648")
	mockobj.Mute(true)
	mockobj.Mute(false)
	mockobj.Mute(true)
	mockobj.RemoveEntity("jjeeemss")
	mockobj.RemoveEntity("emoveEntity")
	mockobj.HandleAmongUsEvent("adadsada", "saad")
	mockobj.HandleAmongUsEvent("jjems", "jjems")

	if err := mockobj.VerifyFuncCall(); err != nil {
		//TODO: Bug, a prio pour une comparaison de string, a chaque
		// fois la comparaison du deuxieme arg est a false, meme si les
		// deux strings sont equals
		t.Errorf("VerifyFuncCall error but should be ok : %s !\n", err)
	}

	functions := make(map[string]func(...interface{}) interface{})
	chanFuncCalled := make(chan interface{}, 1)
	functions["HandleAmongUsEvent"] = func(args ...interface{}) interface{} {
		chanFuncCalled <- true
		return nil
	}

	mockobj = NewMockObj(functions)
	mockobj.HandleAmongUsEvent("sdad", "asdads")
	select {
	case <-chanFuncCalled:
	case <-time.After(3 * time.Second):
		t.Errorf("Own handleAmongUsEvent Func not called !\n")
	}

	mockobj = NewMockObj(nil)

	mockobj.AssertFuncCalledWith("GetEntityId", "6546654648")
	mockobj.AssertFuncCalledWith("Mute", false)
	mockobj.AssertFuncCalledWith("Mute", true)
	mockobj.AssertFuncCalledWith("RemoveEntity", "jjeeemss")
	mockobj.AssertFuncCalledWith("HandleAmongUsEvent", "jjems", "jjems")

	mockobj.GetEntityId("asdasda")
	mockobj.GetEntityId("asdasda")
	mockobj.GetEntityId("6546654648")
	mockobj.GetEntityId("654665464")
	mockobj.GetEntityId("6546654648")
	mockobj.Mute(true)
	mockobj.Mute(false)
	mockobj.Mute(true)
	mockobj.RemoveEntity("jjeeemss")
	mockobj.RemoveEntity("emoveEntity")
	mockobj.HandleAmongUsEvent("adadsada", "saad")
	mockobj.HandleAmongUsEvent("jiasdadjems", "jjems")
	if err := mockobj.VerifyFuncCall(); err == nil {
		t.Errorf("VerifyFuncCall should throw an error !\n")
	}

	return
}
