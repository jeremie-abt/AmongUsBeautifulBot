package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var GlobalVarManager GlobalVarManagerType

func main() {

	test_map_readjson := make(map[string]string)

	json_config_file, _ := os.Open("config.json")
	ret, _ := ioutil.ReadAll(json_config_file)

	json.Unmarshal(ret, &test_map_readjson)
	bot_token := test_map_readjson["bot_token"]

	dg, err := discordgo.New("Bot " + bot_token)

	jejemsMockGuild := NewGuildManager(dg, "766750463524732968")
	GlobalVarManager.GuildManagers = (append(GlobalVarManager.GuildManagers, jejemsMockGuild))

	Gconf := NewGameConfig("config_role.json")

	fmt.Printf("init connexion %v ...\n\n", Gconf)

	// TODO: Gros gros refacto pour gerer tous le monde
	// Il faudra gerer les sessions propre a chacun etc ...
	// je vais plus ou moins mocker tous ca le temps de
	// me familiariser avec le go puis de lire de trois
	// trucs sur la concurency mais faudra faire ca clean
	// ca va etre plutot style

	if err != nil {
		fmt.Printf("Err instantiating bot : %s\n", err)
		return
	}

	dg.AddHandler(VoiceChangeHandler)
	dg.AddHandler(MessageSendHandler)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)

	// Open socket and begin listen
	err = dg.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	defer close(sc)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
