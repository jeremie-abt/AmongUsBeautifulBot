package main

import "fmt"
import "github.com/gorilla/websocket"

/*
	Stock some context within the ws conn
*/
type wbSession struct {
	ctx interface{}
}

func NewWbSession() *wbSession {
	return &wbSession{ctx: nil}
}

func (wbSess *wbSession) setContext(ctx interface{}) {
	wbSess.ctx = ctx
}

func (wbSess *wbSession) context() interface{} {
	return wbSess.ctx
}

type Hub struct {
	register   chan (*websocket.Conn)
	unregister chan (*websocket.Conn)
	msg        chan (*websocket.Conn)
	clients    map[*websocket.Conn]*wbSession
}

func NewHub() *Hub {
	return &Hub{
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		msg:        make(chan *websocket.Conn),
		clients:    make(map[*websocket.Conn]*wbSession),
	}
}

func (h *Hub) Run() {

	for {
		select {
		case con := <-h.register:
			h.clients[con] = nil
		case con := <-h.unregister:
			delete(h.clients, con)
		case msg := <-h.msg:
			fmt.Printf("Message received : %s\n", msg)
		}
	}
}

func (h *Hub) SetContext(c *websocket.Conn, data interface{}) {
	item, found := h.clients[c]
	if !found {
		fmt.Errorf("clients not existing ..\n")
		return
	}
	if item == nil {
		h.clients[c] = NewWbSession()
	}
	h.clients[c].setContext(data)
}

func (h *Hub) Context(c *websocket.Conn) interface{} {
	return h.clients[c].context()
}
