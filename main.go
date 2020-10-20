package main


import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func setup() {
	// I dont know where this func is going for 
	// now But It'll Done somme setup work

	// TODO : Assign randomly Tag (based on some config)
}



var GlobalVarManager GlobalVarManagerType
func main() {

	// TODO : Recheck un peu les struct dans struct.go
	// et init mon gobal var
	
	dg, err := discordgo.New("Bot " + "NzY2NzUyODEwOTI2MzQyMjM5.X4n8Mw.gS1DzyAEiO29ELQqdA-I2zaO5ec")

	testGuid, err := dg.Guild("766750463524732968")
	fmt.Printf("yo le rap%v\n\n", testGuid)

	/*
	**	Not sure yet of the design yet, but I'll do
	**	a big struct containing litle struct more
	**	specific in the part of the code they're using
	**	I'm instantiating for now all the litle struct
	**	under this and I'll do the struct after
	*/	


	// Config struct
	Gconf := NewGameConfig()
	// Discord Channel struct
	discordChan, err := NewDiscordChanStruct("766941016699305984", dg)
	if err != nil {
		fmt.Printf("Err getting channel : %v\n", err)
		return
	}
	// Player struct (recup depuis le chan)

	fmt.Printf("init connexion %v ...\n\n", Gconf)

	// TODO : Needed avant de faire cette feature :
	//	List de Personne
	//	Fonction de list de personne qui en choisit une random
	//	lui assigne le role
	//	Attention a la compat avec les infos que je voudrais plus tard pour une personne
	// TODO : Implementation de la logique qui gere les channels
	// discord, c'est lui qui devra ensuite appeler
	// la creation de player je pense

	// TODO : Receveoir une liste de player et non pas un seul player
	//dcPlayer := NewDiscordPlayer()


	// TODO: Gros gros refacto pour gerer tous le monde
	// Il faudra gerer les sessions propre a chacun etc ...
	// je vais plus ou moins mocker tous ca le temps de
	// me familiariser avec le go puis de lire de trois
	// trucs sur la concurency mais faudra faire ca clean
	// ca va etre plutot style


	fmt.Printf("chantest : %v\n", discordChan)
	if err != nil {
		fmt.Printf("Err instantiating bot : %s\n", err)
		return
	}

	//dg.AddHandler(VoiceStateHandler)
	dg.AddHandler(testHandler)

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

func testHandler(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
		fmt.Printf("discord : %+v\n\n", m.VoiceState)
}

// TODO: etudier le pattern que discordgo utilse pour les events
