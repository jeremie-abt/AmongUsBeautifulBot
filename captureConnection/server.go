package main

import (
	"fmt"
	"net/http"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func testFunc(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)
	fmt.Printf("upgrade : %+v", conn)
	fmt.Printf("upgrade : %+v", conn)

	if err != nil {
		fmt.Printf("\n\niiiiierr : %+v\n", err)
		return
	}
	code := []byte("958677")
	if err := conn.WriteMessage(websocket.TextMessage, code); err != nil {
		fmt.Printf("2err responding : %+v\n", err)
		return
	}
	for {
		// Read
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("1err : %s\n", err)
			return
		}

		// And respond :
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Printf("2err responding : %+v\n", err)
			return
		}
	}
}

func main() {
	router := mux.NewRouter()

	// TODO: Debug / trouver un moyen de fix ca
	// En gros quand je passe par ngrok il
	// pense que je tape sur /socket.io/
	router.HandleFunc("/socket.io/", testFunc)

	// TODO : jai casi tous setup, par contre je narrive pas a communique depuis ngrok
	// je ne recois rien ici je ne sais pas pk
	// -> A voir pour debug mux direct, tenter de debug le package
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
