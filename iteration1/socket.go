// la partie des sockets doit en connaitre le moins possible
// sur le reste du code (l'etat des strucs etc ...)
// On va donc juste matcher des events avec des fonctions qui
// savent quoi faire
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
)

// Pulse : Chan for sending heartbeat
func BeginListenSocket(crashReport chan<- interface{}) {
	router := mux.NewRouter()
	server, err := socketio.NewServer(nil)

	if err != nil {
		fmt.Printf("err : %s\n", err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		// Je pense que cette fonction permet exactement
		// de jouer avec le context
		s.Context()
		println("Connected ...")
		return nil
	})
	server.OnEvent("/", "connectCode", socketEvConnectCode)
	server.OnEvent("/", "lobby", socketEvLobby)
	server.OnEvent("/", "state", socketEvState)
	server.OnEvent("/", "player", socketEvPlayer)

	server.OnError("/", handleSocketError)
	server.OnDisconnect("/", handleDisconnection)
	go server.Serve()
	defer server.Close()

	router.Handle("/socket.io/", server)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Serving...\n")
	log.Fatal(srv.ListenAndServe())
	// send notif to relauch this goroutines
	crashReport <- struct{}{}
}

/*
	All SocketEvXXX func are websocket handlerj
*/
func socketEvConnectCode(conn socketio.Conn, msg string) {
	fmt.Printf("connection code event : %s\n", msg)

	fmt.Printf("Voici le context : %+v\n", conn.Context())
	propagateEvent(conn, &eventType{
		eventType: CONNECTCODE,
		msg:       msg,
	})
}

func socketEvLobby(conn socketio.Conn, msg string) {
	// msg : "LobbyCode":"BNHODF","Region":0}
	// TODO
	fmt.Printf("lobby event : %s\n", msg)

	//propagateEvent(conn, &eventType{
	//	eventType: LOBBY,
	//	msg: msg,
	//})
}

// msg :
func (s *socketHandling) HandleDisconnection(msg string, captureCode string) {

}

// 0 -> LOBBY
// 1 debut de tour je crois (jetais imposteur) / reprise dun tour egalement
// 2 -> Kill / je lai eu aussi en appuyant sur le bouton Wtfff
// 3 -> MENU
func socketEvState(conn socketio.Conn, msg string) {
	// TODO
	fmt.Printf("state event : %s\n", msg)

	propagateEvent(conn, &eventType{
		eventType: STATE,
		msg:       msg,
	})
}

// msg example :
//	{"Action":4,"Name":"jejelaterr","IsDead":false,"Disconnected":false,"Color":7}
// {"Action":5,"Name":"I","IsDead":false,"Disconnected":true,"Color":5}
// Action :
// 0 -> connection au lobby
// 1 -> ???
// 2 -> mourir -> Tu as le name de la person qui vient de mourir
// 3 -> changement de couleur
// 	5 -> Disconnected (a confirmer)
// 6 -> Je crois que c quand une personne est ejecte aux votes
func socketEvPlayer(conn socketio.Conn, msg string) {
	propagateEvent(conn, &eventType{
		eventType: PLAYER,
		msg:       msg,
	})
}

func handleSocketError(conn socketio.Conn, err error) {
	// TODO
	fmt.Printf("socket errror : %s\n", err)
}

// TODO: Propager les Events de maniere clean
// Pour cela, je vais tenter de cree une struct event
// et de la faire passer en chan
func handleDisconnection(conn socketio.Conn, reason string) {
	// TODO
	fmt.Printf("socket disconnection : %s\n", reason)
}

// ------------------------------------ PRIVATE Struct
// {"Action":5,"Name":"I","IsDead":false,"Disconnected":true,"Color":5}
type playerEventT struct {
	Action       int    `json:"Action"`
	Name         string `json:"Name"`
	Isdead       bool   `json:"IsDead"`
	Disconnected bool   `json:"Disconnected"`
	Color        int    `json:"Color"`
}

// ------------------------------------ PRIVATE
func propagateEvent(conn socketio.Conn, msg *eventType) {
	// Test avec mac

	if msg.eventType == CONNECTCODE {
		if CheckAndLinkSocketCode(msg.msg) {
			conn.SetContext(msg.msg)
		}
		return
	}
	captureCode := conn.Context().(string)
	if msg.eventType == PLAYER {
		handleSocketPlayerEvent(msg.msg, captureCode)
	}
	if msg.eventType == STATE {
		handleStateEvent(msg.msg, captureCode)
	}

}

func handleSocketPlayerEvent(msg string, captureCode string) {
	// je recois un joueurs (Event 0 -> il vient de se connecter)
	// Comme il vient de se connecter, il ne doit etre nul part

	var playerEvent playerEventT
	err := json.Unmarshal([]byte(msg), &playerEvent)
	if err != nil {
		// TODO : Error handling
		// Probably should probably log
		fmt.Printf("Error : %s\n", err)
	}
	fmt.Printf("PLayer Event : %+v\n", playerEvent)
	handlePlayerEvent(&playerEvent, captureCode)
}

//------------------------------- Outbound function
/*
	Ces fonctions servent a interagir avec le reste du code
*/

func handlePlayerEvent(playerEvt *playerEventT, captureCode string) {

	curGame, _ := GetGameFromCode(captureCode)
	if curGame == nil {
		fmt.Printf("Warning: Following catpure code doesn't exist : %s\n", captureCode)
		return
	}
	discordPlayer := curGame.GetDiscordPlayerFromSocketEvt(playerEvt.Name)
	fmt.Printf("Bonjour voici le discord player : %+v\n\n", discordPlayer)
	if playerEvt.Action == 2 {
		discordPlayer.IsAlive = false
	}
}

func handleStateEvent(msg string, captureCode string) {

	curGame, guildManager := GetGameFromCode(captureCode)
	fmt.Printf("all player : %+v\n", curGame.DiscordPlayers)
	if curGame == nil {
		fmt.Printf("Warning: Following catpure code doesn't exist : %s\n", captureCode)
		return
	}

	if msg == "1" {
		/*
			Debut de tour, on cherche a mute tous le monde
		*/
		fmt.Printf("Yeah\n")
		curGame.MuteAllPlayer(guildManager.GuildId, true)
	} else if msg == "2" || msg == "3" {
		// On va tenter de faire une simple method pour mute
		// tous le user
		curGame.MuteAllPlayer(guildManager.GuildId, false)
	}
}
