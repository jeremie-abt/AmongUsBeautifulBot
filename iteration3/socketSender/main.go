package main

/*
	This litle program is used to fake the among us capture code
	websocket program
*/

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:3000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	secEndpoint := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	thirdEndpoint := url.URL{Scheme: "ws", Host: *addr, Path: "/other"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	secDialer, _, err := websocket.DefaultDialer.Dial(secEndpoint.String(), nil)
	thirdDialer, _, err := websocket.DefaultDialer.Dial(thirdEndpoint.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	defer close(done)

	_ = thirdDialer.WriteMessage(websocket.TextMessage, []byte("CoucouCnousOther"))
	//_ = c.WriteMessage(websocket.TextMessage, []byte("CoucouCnous"))
	_ = secDialer.WriteMessage(websocket.TextMessage, []byte("CoucouCnous2"))
	_ = thirdDialer.WriteMessage(websocket.TextMessage, []byte("CoucouCnousOther"))
	//_ = c.WriteMessage(websocket.TextMessage, []byte("CoucouCnous"))
	_ = secDialer.WriteMessage(websocket.TextMessage, []byte("CoucouCnous2"))

	sc := make(chan os.Signal, 1)
	defer close(sc)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
