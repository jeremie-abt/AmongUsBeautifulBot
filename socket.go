package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	socketio "github.com/googollee/go-socket.io"

)

// la partie des sockets doit en connaitre le moins possible
// sur le reste du code (l'etat des strucs etc ...)
// On va donc juste matcher des events avec des fonctions qui
// savent quoi faire
func propagateEvent(conn socketio.Conn, msg *eventType) {
	// Test avec mac

	if msg.eventType == CONNECTCODE {
		// Il faut lie la bonne struct game
		// Avec ce code


		if SocketConnectCodeHandle(msg.msg) {
			conn.SetContext(msg.msg)
		}
	}
}

/***
**		All SocketEvXXX func are websocket handlerj
***/
type Socketcontext struct {
		
	captureCode string
}

// Receive connect code -> Don't really know yet the utility
// But maybe to link guild account or something
// pour ce qui est de lidentification des clients, je peux rajouter
// des var dans le context de la connexion a prio
// donc mettre un id de game ou un truc du genre
func socketEvConnectCode(conn socketio.Conn, msg string) {
	fmt.Printf("connection code event : %s\n", msg)

	propagateEvent(conn, &eventType{
		eventType: CONNECTCODE,
		msg: msg,
	})
}

// je ne sais pas encore quel est l'utilite de cet Event
// msg : "LobbyCode":"BNHODF","Region":0}
func socketEvLobby(conn socketio.Conn, msg string) {
	// TODO
	fmt.Printf("lobby event : %s\n", msg)

	//propagateEvent(conn, &eventType{
	//	eventType: LOBBY,
	//	msg: msg,
	//})
}

// msg :
// 0 -> LOBBY
// 1 debut de tour je crois (jetais imposteur) / reprise dun tour egalement
// 2 -> Kill / je lai eu aussi en appuyant sur le bouton Wtfff
// 3 -> MENU
func socketEvState(conn socketio.Conn, msg string) {
	// TODO
	fmt.Printf("state event : %s\n", msg)

	//propagateEvent(conn, &eventType{
	//	eventType: STATE,
	//	msg: msg,
	//})
}

// msg example :
//	{"Action":4,"Name":"jejelaterr","IsDead":false,"Disconnected":false,"Color":7}
// {"Action":5,"Name":"I","IsDead":false,"Disconnected":true,"Color":5}
// Action :
// 0 -> connection au lobby
// 1 -> ???
// 2 -> ???
// 	5 -> Disconnected (a confirmer)
// 6 -> Je crois que c quand une personne est ejecte aux votes
func socketEvPlayer(conn socketio.Conn, msg string) {
	// TODO
	fmt.Printf("player event : %s\n", msg)

	//PropagateEvent(conn, &eventType{
	//	eventType: PLAYER,
	//	msg: msg,
	//})
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