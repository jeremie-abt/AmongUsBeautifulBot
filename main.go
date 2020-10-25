package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
)


// TODO : Comment faire ca plus clean que 50 ?
var GlobalVarManager GlobalVarManagerType

func main() {
	
	test_map_readjson := make(map[string]string)

	json_config_file, _ := os.Open("config.json")
	ret, _ := ioutil.ReadAll(json_config_file)

	json.Unmarshal(ret, &test_map_readjson)
	bot_token := test_map_readjson["bot_token"]
	
	dg, err := discordgo.New("Bot " + bot_token)

	jejemsMockGuild := NewGuildManager(dg, "766750463524732968")
	GlobalVarManager.GuildManagers = (
		append(GlobalVarManager.GuildManagers, jejemsMockGuild))


	// Config struct
	Gconf := NewGameConfig()
	// Discord Channel struct
//	discordChan, err := NewDiscordChanStruct("766750463524732968", dg)
//	if err != nil {
//		fmt.Printf("Err getting channel : %v\n", err)
//		return
//	}
	// Player struct (recup depuis le chan)

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

	dg.AddHandler(voiceChangeHandler)
	dg.AddHandler(messageSendHanlder)

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


// Managing VoiceUpdate change
func voiceChangeHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	currentGuild := GlobalVarManager.getGuildObj(m.VoiceState.GuildID)
	currentGuild.HandleVoiceChange(m.VoiceState)

}

func messageSendHanlder(s *discordgo.Session, m *discordgo.MessageCreate) {
	curMessage := m.Message
	currentGuild := GlobalVarManager.getGuildObj(curMessage.GuildID)

	if strings.HasPrefix(strings.ToLower(curMessage.Content), ".creategame") {
		// get le channel pour commencer une game

		if len(curMessage.Content) <= 11 {
			return // TODO : Gerer la gestion d'erreur
			// Probablement renvoyer un usage
		} else {
			gameChanName := curMessage.Content[12:]
			allChan, err := s.GuildChannels(curMessage.GuildID)
			if err != nil {
				fmt.Printf("Voici le err : %+v\n", err)
				return
			}
			// TODO : Il reste q check si le chan existe, que ce soit bien un type quil faut
			// Et sinon creer le chan jimagine
			// et ensuite finalement creer la game
			var chanMatched string
			for _, curChan := range(allChan) {
				if (curChan.Name == gameChanName &&
						curChan.Type == discordgo.ChannelTypeGuildVoice) {
					chanMatched = curChan.ID
				}
			}

			newGame := NewGame(chanMatched, NewGameConfig())
			currentGuild.AttachGame(newGame)

		}
	}
}
