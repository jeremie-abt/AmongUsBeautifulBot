package test

import (
	//	"fmt"
	"net"
	"net/http"
	"net/url"
	"testing"

	//"github.com/bwmarrin/discordgo"
	//socketio "github.com/googollee/go-socket.io"
	myBot "github.com/jeremie-abt/AmongUsBeautifulBot/iteration2"
)

// TODO: Find a better way to mock this
type mock_socketConn struct {
	context interface{}
}

func TestSocketEvLobby(t *testing.T) {
	mockedGame := NewMockObj(nil)
	mockConn := &mock_socketConn{}
	mockConn.SetContext(mockedGame)
	mockedGame.AssertFuncCalledWith("HandleAmongUsEvent", "0", myBot.LOBBY)
	mockedGame.AssertFuncCalledWith("HandleAmongUsEvent", "1", myBot.LOBBY)
	mockedGame.AssertFuncCalledWith("HandleAmongUsEvent", "2", myBot.LOBBY)

	myBot.SocketEvLobby(mockConn, "0")
	myBot.SocketEvLobby(mockConn, "1")
	myBot.SocketEvLobby(mockConn, "2")

	if err := mockedGame.VerifyFuncCall(); err != nil {
		t.Errorf("Lobby Event bad functions calls : %s\n", err)
	}
}

func TestSocketEvState(t *testing.T) {
	t.Errorf("to impl\n")
}

func TestSocketEvPlayer(t *testing.T) {
	t.Errorf("to impl\n")
}

// -------- socketio Conn struct mocking
func (mckConn *mock_socketConn) ID() string {
	return ""
}

func (mckConn *mock_socketConn) Close() error {
	return nil
}

func (mckConn *mock_socketConn) URL() url.URL {
	ret, _ := url.ParseRequestURI("localhost")
	return *ret
}

func (mckConn *mock_socketConn) LocalAddr() net.Addr {
	return nil
}

func (mckConn *mock_socketConn) RemoteAddr() net.Addr {
	return nil
}

func (mckConn *mock_socketConn) RemoteHeader() http.Header {
	return nil
}

func (mckConn *mock_socketConn) Context() interface{} {
	return mckConn.context
}

func (mckConn *mock_socketConn) SetContext(v interface{}) {
	mckConn.context = v
	return
}

func (mckConn *mock_socketConn) Namespace() string {
	return ""
}

func (mckConn *mock_socketConn) Emit(msg string, v ...interface{}) {
	return
}

func (mckConn *mock_socketConn) Join(room string) {
	return
}

func (mckConn *mock_socketConn) Leave(room string) {
	return
}

func (mckConn *mock_socketConn) LeaveAll() {
	return
}

func (mckConn *mock_socketConn) Rooms() []string {
	return []string{""}
}

///// -------- Game mock method
type mockGame struct {
	*myBot.Game
}

type mockAuPlayers struct {
	*myBot.AuPlayers
}

//func NewMockGame() *mockGame {
//
//	//return &mockGame{
//	//	&myBot.Game{
//	//		ChannelId: "YOtest",
//	//		AuPlayers: &mockAuPlayers{
//	//			&myBot.AuPlayers{
//	//				AuPlayers: make(map[string]*myBot.AuPlayer),
//	//			},
//	//		},
//	//		AuCaptureCode: "testcode",
//	//	},
//	//}
//}
