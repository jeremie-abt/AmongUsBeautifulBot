package domain_test

import (
	"fmt"
	"testing"
)

import (
	"github.com/golang/mock/gomock"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/mocks"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework"
	"github.com/stretchr/testify/assert"
)

const GameId = "654313813464"
const PlayerId = "65654646464446464"

func TestStartGame(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	assert := assert.New(t)

	repoMock := mocks.NewMockRepository(mockCtrl)
	repoMock.EXPECT().AddGame(GameId).Return(nil).Times(1)
	repoMock.EXPECT().AddGame(GameId).Return(fmt.Errorf(framework.ErrAlreadyExist)).Times(2)

	bot := domain.NewBotCommandHandler(repoMock)
	err := bot.StartGame(GameId)
	assert.NoError(err)
	err = bot.StartGame(GameId)
	assert.EqualError(err, framework.ErrAlreadyExist)
	err = bot.StartGame(GameId)
	assert.EqualError(err, framework.ErrAlreadyExist)
}

/*
	Je stop mes tests ici pour linstant sur le domaine car :
		Il na aucune logique intrinseque, il ne fait que call
		le repository pour faire des operation crud
*/
