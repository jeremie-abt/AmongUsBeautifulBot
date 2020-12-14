unittest:
	@go test -coverpkg=./... pkg/domain/tests/AmongUsEvtHandler_test.go
	@go test -coverpkg=./... pkg/infra/services/discordBotWithDiscordgo_test.go
	@go test -coverpkg=./... pkg/domain/tests/botCommandHandler_test.go
	
integrationtest:
	# Test for redisRepo
	@sudo docker run --rm --name redis --net redis -p 6701:6379 redis:6.0-alpine > /dev/null &
	@sleep 5
	@REDIS_HOST=6701 go test -coverpkg=./... pkg/infra/framework/redisRepositor* pkg/infra/framework/framework.go
	@sudo docker container stop redis > /dev/null

runBot:
	@go run cmd/bot/bot.go

updateMock:
	@mockgen -destination=mocks/mock_voipServer.go -package=mocks "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework" VoipServer
	@mockgen -destination=mocks/mock_repository.go -package=mocks "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework" Repository
	@mockgen -destination=mocks/mock_domain_botcommand.go -package=mocks "github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain" IBotCommand

startStack: startRedisStack

stopStack : stopRedisStack

startRedisStack:
	@sudo docker run -it --rm --name redis --net redis -p 6700:6379 redis:6.0-alpine

stopRedisStack:
	@sudo docker stop redis

