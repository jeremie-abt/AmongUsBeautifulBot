package bot

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	//"github.com/gorilla/mux"
	//	"github.com/jeremie-abt/AmongUsBeautifulBot/iteration2/logger"
)

func SocketServer() {
	server, err := socketio.NewServer(nil)

	if err != nil {
		// TODO: handle error
		fmt.Printf("err : %s", err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		return nil
	})
	server.OnEvent("/", "connectCode", SocketEvConnectCode)
	server.OnEvent("/", "lobby", SocketEvLobby)
	server.OnEvent("/", "state", SocketEvState)
	server.OnEvent("/", "player", SocketEvPlayer)
}

func SocketEvConnectCode(s socketio.Conn, msg string) error {
	// TODO
	s.SetContext(msg)
	return nil
}

// TODO: Reflexion : Refacto en une seule fonction (Je pense qu'il faut le faire)
func SocketEvLobby(s socketio.Conn, msg string) error {
	game := s.Context().(IGame)

	game.HandleAmongUsEvent(msg, LOBBY)
	return nil
}

func SocketEvState(s socketio.Conn, msg string) error {
	game := s.Context().(IGame)

	game.HandleAmongUsEvent(msg, STATE)
	return nil
}

func SocketEvPlayer(s socketio.Conn, msg interface{}) error {
	game := s.Context().(IGame)

	game.HandleAmongUsEvent(msg, PLAYER)
	return nil
}
