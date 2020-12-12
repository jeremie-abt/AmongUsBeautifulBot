package domain_test

import "testing"

import (
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/mocks"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain/entity"
)

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const gameId = "5464646464854"

func TestHandleEvent(t *testing.T) {
	// Mock right ports
	// Tchek that the logic is ok
	mockCtrlVoip := gomock.NewController(t)
	mockCtrlRepo := gomock.NewController(t)
	defer mockCtrlRepo.Finish()
	defer mockCtrlVoip.Finish()

	assert := assert.New(t)

	voipServMock := mocks.NewMockVoipServer(mockCtrlVoip)
	repoMock := mocks.NewMockRepository(mockCtrlRepo)

	// GAME EVENT TEST (see AmongUsEvtType in entity file)
	voipServMock.EXPECT().UnMuteAll(entity.NewGame(gameId)).Return(nil).Times(2)
	voipServMock.EXPECT().MuteAll(entity.NewGame(gameId)).Return(nil).Times(1)
	assert.NoError(nil)

	amongUsEvtHandlerDomain := domain.NewAmongUsEvtHandler(
		voipServMock,
		repoMock,
	)
	eventsToTests := []entity.AmongUsEvtType{
		entity.GAMELOBBY, entity.GAMEBREAK, entity.GAMEBEGINROUND}
	for _, event := range eventsToTests {
		amongUsEvtHandlerDomain.HandleEvent(generateEvent(event))
	}

}

func generateEvent(lb entity.AmongUsEvtType) entity.AmongUsEvent {

	return entity.AmongUsEvent{
		AttachedGame: &entity.Game{
			Id: gameId,
		},
		Evttype: lb,
	}
}
