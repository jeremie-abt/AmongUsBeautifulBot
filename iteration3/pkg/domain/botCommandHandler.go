package domain

/*
	IbotCommand Ports implementation
*/
type BotCommandHandler struct{}

func NewBotCommandHandler() *BotCommandHandler {
	return &BotCommandHandler{}
}

func (b *BotCommandHandler) StartGame(id string) error {
	return nil
}

func (b *BotCommandHandler) StopGame(id string) error {
	return nil
}

func (b *BotCommandHandler) AddPlayer(gameId string, playerId string) error {
	return nil
}

func (b *BotCommandHandler) DeletePlayer(gameId string, playerId string) error {
	return nil
}
