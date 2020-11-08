// En phase de test mais le code a ete bouge dans eventHandling.go
//
////package main
//
//import (
//	"fmt"
//	"net/http"
//	"log"
//	"time"
//
//	socketio "github.com/googollee/go-socket.io"
//	"github.com/gorilla/mux"
//)
//
////TODO Integrer tous ce code dans mon bot
//
///***
//**		All SocketEvXXX func are websocket handlerj
//***/
//
//// Receive connect code -> Don't really know yet the utility
//// But maybe to link guild account or something
//// pour ce qui est de lidentification des clients, je peux rajouter
//// des var dans le context de la connexion a prio
//// donc mettre un id de game ou un truc du genre
//func socketEvConnectCode(conn socketio.Conn, msg string) {
//	// TODO
//	fmt.Printf("connection code event : %s\n", msg)
//}
//
//// je ne sais pas encore quel est l'utilite de cet Event
//// msg : "LobbyCode":"BNHODF","Region":0}
//func socketEvLobby(conn socketio.Conn, msg string) {
//	// TODO
//	fmt.Printf("lobby event : %s\n", msg)
//}
//
//// msg :
//// 0 -> LOBBY
//// 1 debut de tour je crois (jetais imposteur) / reprise dun tour egalement
//// 2 -> Kill / je lai eu aussi en appuyant sur le bouton Wtfff
//// 3 -> MENU
//func socketEvState(conn socketio.Conn, msg string) {
//	// TODO
//	fmt.Printf("state event : %s\n", msg)
//}
//
//// msg example :
////	{"Action":4,"Name":"jejelaterr","IsDead":false,"Disconnected":false,"Color":7}
//// {"Action":5,"Name":"I","IsDead":false,"Disconnected":true,"Color":5}
//// Action :
//// 0 -> connection au lobby
//// 1 -> ???
//// 2 -> ???
//// 	5 -> Disconnected (a confirmer)
//// 6 -> Je crois que c quand une personne est ejecte aux votes
//func socketEvPlayer(conn socketio.Conn, msg string) {
//	// TODO
//	fmt.Printf("player event : %s\n", msg)
//}
//
//func handleSocketError(conn socketio.Conn, err error) {
//	// TODO
//	fmt.Printf("socket errror : %s\n", err)
//}
//
//func handleDisconnection(conn socketio.Conn, reason string) {
//	// TODO
//	fmt.Printf("socket disconnection : %s\n", reason)
//}
//
//func main() {
//	router := mux.NewRouter()
//	server, err := socketio.NewServer(nil)
//
//	if err != nil {
//		fmt.Printf("err %s\n", err)
//	}
//	// TODO: note persos : Allez voir la doc
//	server.OnConnect("/", func(s socketio.Conn) error {
//		s.SetContext("")
//		println("Yo je suis connecte !!\n")
//		return nil
//	})
//
//	server.OnEvent("/", "connectCode", socketEvConnectCode)
//	server.OnEvent("/", "lobby", socketEvLobby)
//	server.OnEvent("/", "state", socketEvState)
//	server.OnEvent("/", "player", socketEvPlayer)
//
//	server.OnError("/", handleSocketError)
////	server.OnDisconnect("")
//
//	go server.Serve()
//	defer server.Close()
//
//	router.Handle("/socket.io/", server)
//
//	// TODO : jai casi tous setup, par contre je narrive pas a communique depuis ngrok
//	// je ne recois rien ici je ne sais pas pk
//	// -> A voir pour debug mux direct, tenter de debug le package
//	srv := &http.Server{
//		Handler: router,
//		Addr:    "127.0.0.1:3000",
//		// Good practice: enforce timeouts for servers you create!
//		WriteTimeout: 15 * time.Second,
//		ReadTimeout:  15 * time.Second,
//	}
//
//	log.Fatal(srv.ListenAndServe())
//}
