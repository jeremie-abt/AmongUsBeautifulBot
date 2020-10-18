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

func main() {

	Gconf := NewGameConfig()
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
	dcPlayer := NewDiscordPlayer()


	return
	// TODO : Essayer de faire la feature du talkie


	dg, err := discordgo.New("Bot " + "NzY2NzUyODEwOTI2MzQyMjM5.X4n8Mw.SpMt4yaPAtkeksY5RRRdIQQZJnk")
	if err != nil {
		fmt.Printf("Err instantiating bot : %s\n", err)
	}

	dg.AddHandler(MyFirstHandler)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

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


func MyFirstHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	fmt.Printf("Bonour ma fonctio est called %v\n\n", m.Message)
}
