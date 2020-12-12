package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	//	"fmt"
	"io/ioutil"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/domain"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/framework"
	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration3/pkg/infra/services"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	json_config_file, _ := os.Open("config.json")
	ret, _ := ioutil.ReadAll(json_config_file)
	fmt.Printf("Yeaahhhh : %s\n", ret)
	test_map_readjson := make(map[string]string)
	json.Unmarshal(ret, &test_map_readjson)
	fmt.Printf("Yeaahhhh : %+v\n", test_map_readjson)
	bot_token := test_map_readjson["bot_token"]
	dg, _ := discordgo.New("Bot " + bot_token)

	//	dg.AddHandler(MessageSendHandler)
	discordBotAdapter := services.NewDiscordBotAdapter(
		domain.NewBotCommandHandler(
			framework.NewInMemRepository()))
	dg.AddHandler(discordBotAdapter.HandleWrittenMessage)
	//	dg.AddHandler(HandleJoinChannel)

	err := dg.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	defer close(sc)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
