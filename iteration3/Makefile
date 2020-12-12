unittest:
	@go test -v pkg/domain/tests/AmongUsEvtHandler_test.go
	@go test -v pkg/infra/services/discordBotWithDiscordgo_test.go
	@go test -v pkg/domain/tests/botCommandHandler_test.go

runBot:
	@go run cmd/bot/bot.go

updateMock:
	@mockgen -destination=mocks/mock_voipServer.go -package=mocks "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework" VoipServer
	@mockgen -destination=mocks/mock_repository.go -package=mocks "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework" Repository
	@mockgen -destination=mocks/mock_domain_botcommand.go -package=mocks "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain" IBotCommand

