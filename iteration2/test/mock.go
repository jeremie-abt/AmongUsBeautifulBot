package test

import (
	"errors"
	"fmt"
	"reflect"

	socketio "github.com/googollee/go-socket.io"
	mybot "github.com/jeremie-abt/AmongUsBeautifulBot/iteration2"
)

// Test to implement this kind of mock :
// 	https://dev.to/jonfriesen/mocking-dependencies-in-go-1h4d
type Mock interface {

	// ** Game
	HandleAmongUsEvent(interface{}, mybot.AuEventType)
	HandleLobbyEvent(string)
	HandlePlayerEvent(*mybot.AmongUsPlayerEvent)
	HandleStateEvent(mybot.AuEventState)
	ResetGameState()
	BeginRound()
	BeginChattingState()
	GetGameFromCode(string) *mybot.Game

	// ** Socket

	SocketEvLobby(s socketio.Conn, msg string)

	// ** AuPlayer

	IsAlive()
	SetIsAlive(isAlive bool)
	UnMute()

	// ** Mute exist for entities and for AuPlayer
	Mute(bool)

	// ** Entities

	GetEntityId(string) string
	GetEntity(entityId string) interface{}
	AddEntity(entity interface{})
	RemoveEntity(entity string)

	// ** mock reel code logic
	VerifyFuncCall() error
	AssertFuncCalledWith(funcName string, args ...interface{})
}

func logUnknownMockFunc(...interface{}) interface{} {
	return nil
}

func NewMockObj(functions map[string]func(...interface{}) interface{}) Mock {
	// TODO: Is there any better pattern than type assertion here ??
	mockObj := &MockObj{}
	mockReflect := reflect.ValueOf(mockObj)
	mockReflect = mockReflect.Elem()
	nbFields := mockReflect.NumField()
	mockReflectType := mockReflect.Type()

	for i := 0; i < nbFields; i++ {
		v := mockReflectType.Field(i)
		field := mockReflect.FieldByName(v.Name)
		// Remove Func from MockObj field : HandleAmongUsEventFunc -> HandleAmongUsEvent
		funcName := v.Name[0 : len(v.Name)-4]

		if !field.CanSet() {
			continue
		}
		if ret, found := functions[funcName]; found {
			field.Set(reflect.ValueOf(ret))
		} else {
			field.Set(reflect.ValueOf(logUnknownMockFunc))
		}
	}
	mockObj.funcCallWithArgs = []*functionCall{}
	return mockObj
}

type MockObj struct {
	HandleAmongUsEventFunc func(...interface{}) interface{}
	HandleLobbyEventFunc   func(...interface{}) interface{}
	HandlePlayerEventFunc  func(...interface{}) interface{}
	SocketEvLobbyFunc      func(...interface{}) interface{}
	GetEntityIdFunc        func(...interface{}) interface{}
	MuteFunc               func(...interface{}) interface{}
	GetEntityFunc          func(...interface{}) interface{}
	AddEntityFunc          func(...interface{}) interface{}
	RemoveEntityFunc       func(...interface{}) interface{}
	ResetGameStateFunc     func(...interface{}) interface{}
	BeginRoundFunc         func(...interface{}) interface{}
	BeginChattingStateFunc func(...interface{}) interface{}
	IsAliveFunc            func(...interface{}) interface{}
	SetIsAliveFunc         func(...interface{}) interface{}
	UnMuteFunc             func(...interface{}) interface{}
	GetGameFromCodeFunc    func(...interface{}) interface{}
	HandleStateEventFunc   func(...interface{}) interface{}
	funcCallWithArgs       []*functionCall
	assertFuncCallWithArgs []*functionCall
}

/*
	struct for keeping track of call order and with which args
*/
type functionCall struct {
	funcName string
	args     []interface{}
}

/*
	All the complex reflect logic should appen there not in the
	two previous func (assertCalledWith and recordFuncCall)
*/
func (mockObj *MockObj) VerifyFuncCall() error {
	assertFuncCalls := mockObj.assertFuncCallWithArgs
	funcCalls := mockObj.funcCallWithArgs

	if len(funcCalls) < len(assertFuncCalls) {
		return errors.New("not enough function called")
	}

	curIndex := 0
	for _, assertCall := range assertFuncCalls {
		isCalled := false
		wrongArgs := false

		if curIndex >= len(funcCalls) {
			return fmt.Errorf("some functions not called\n")
		}
		for subIndex, funcCall := range funcCalls[curIndex:] {
			if assertCall.funcName == funcCall.funcName {
				wrongArgs = true
				if isFuncCalledWithSameArgs(assertCall.args, funcCall.args) {
					curIndex = subIndex + curIndex
					curIndex += 1
					isCalled = true
					break
				}
			}
			//fmt.Printf("funcCall : %+v\n", funcCall)
		}
		if isCalled == false {
			if wrongArgs == true {
				//	fmt.Printf("\n\nici\n\n")
				//	fmt.Printf("assert : %+v\n", assertCall)

				return fmt.Errorf("func : %s not called with good args\n", assertCall.funcName)
			}
			return fmt.Errorf("func : %s not called\n", assertCall.funcName)
		}
	}

	return nil
}

func isFuncCalledWithSameArgs(args1 []interface{}, args2 []interface{}) bool {
	if len(args1) != len(args2) {
		return false
	}
	for index, arg1 := range args1 {
		arg2 := args2[index]
		if reflect.ValueOf(arg2).Kind() != reflect.ValueOf(arg1).Kind() {
			return false
		}

		kind := reflect.ValueOf(arg2).Kind()
		// If not a composite type
		if kind != reflect.UnsafePointer && kind != reflect.Struct &&
			kind != reflect.Slice && kind != reflect.Ptr &&
			kind != reflect.Map && kind != reflect.Interface &&
			kind != reflect.Func && kind != reflect.Chan &&
			kind != reflect.Array {

			if arg1 != arg2 {
				return false
			}
		}
	}
	return true
}

func (mockObj *MockObj) AssertFuncCalledWith(funcName string, args ...interface{}) {
	recordFuncCall(&mockObj.assertFuncCallWithArgs, funcName, args...)
	return
}

func recordFuncCall(functionCalls *[]*functionCall, funcName string, args ...interface{}) {
	*functionCalls = append(*functionCalls, &functionCall{
		funcName: funcName,
		args:     args,
	})
}

func (mockobj *MockObj) HandleAmongUsEvent(arg1 interface{}, arg2 mybot.AuEventType) {
	_ = mockobj.HandleAmongUsEventFunc(arg1, arg2)
	recordFuncCall(&mockobj.funcCallWithArgs, "HandleAmongUsEvent", arg1, arg2)
}

func (mockobj *MockObj) HandleLobbyEvent(msg string) {
	_ = mockobj.HandleLobbyEventFunc(msg)
	recordFuncCall(&mockobj.funcCallWithArgs, "HandleLobbyEvent", msg)
}

func (mockobj *MockObj) HandlePlayerEvent(arg *mybot.AmongUsPlayerEvent) {
	_ = mockobj.HandlePlayerEventFunc(arg)
	recordFuncCall(&mockobj.funcCallWithArgs, "HandlePlayerEvent", arg)
}

func (mockobj *MockObj) SocketEvLobby(sess socketio.Conn, msg string) {
	_ = mockobj.SocketEvLobbyFunc(sess, msg)
	recordFuncCall(&mockobj.funcCallWithArgs, "SocketEvLobby", sess, msg)
}

func (mockobj *MockObj) GetEntityId(name string) string {
	ret := mockobj.GetEntityIdFunc(name)
	recordFuncCall(&mockobj.funcCallWithArgs, "GetEntityId", name)
	if reflect.ValueOf(ret).Kind() == reflect.String {
		return ret.(string)
	} else {
		return ""
	}
}

func (mockobj *MockObj) Mute(arg bool) {
	_ = mockobj.MuteFunc(arg)
	recordFuncCall(&mockobj.funcCallWithArgs, "Mute", arg)
}

func (mockobj *MockObj) GetEntity(msg string) interface{} {
	ret := mockobj.GetEntityFunc(msg)
	recordFuncCall(&mockobj.funcCallWithArgs, "GetEntity", msg)
	return ret
}

func (mockobj *MockObj) AddEntity(msg interface{}) {
	_ = mockobj.AddEntityFunc(msg)
	recordFuncCall(&mockobj.funcCallWithArgs, "AddEntity", msg)
}

func (mockobj *MockObj) RemoveEntity(msg string) {
	_ = mockobj.RemoveEntityFunc(msg)
	recordFuncCall(&mockobj.funcCallWithArgs, "RemoveEntity", msg)
}

func (mockobj *MockObj) ResetGameState() {
	_ = mockobj.ResetGameStateFunc()
	recordFuncCall(&mockobj.funcCallWithArgs, "ResetGameState")
}

func (mockobj *MockObj) BeginRound() {
	_ = mockobj.BeginRoundFunc()
	recordFuncCall(&mockobj.funcCallWithArgs, "BeginRound")
}

func (mockobj *MockObj) BeginChattingState() {
	_ = mockobj.BeginChattingStateFunc()
	recordFuncCall(&mockobj.funcCallWithArgs, "BeginChattingState")
}

func (mockobj *MockObj) IsAlive() {
	_ = mockobj.IsAliveFunc()
	recordFuncCall(&mockobj.funcCallWithArgs, "IsAlive")
}

func (mockobj *MockObj) SetIsAlive(arg bool) {
	_ = mockobj.SetIsAliveFunc(arg)
	recordFuncCall(&mockobj.funcCallWithArgs, "SetIsAlive")
}

func (mockobj *MockObj) UnMute() {
	_ = mockobj.UnMuteFunc()
	recordFuncCall(&mockobj.funcCallWithArgs, "BeginChattingState")
}

func (mockobj *MockObj) GetGameFromCode(code string) *mybot.Game {
	_ = mockobj.GetGameFromCodeFunc()
	recordFuncCall(&mockobj.funcCallWithArgs, "GetGameFromCode", code)
	return nil
}

func (mockobj *MockObj) HandleStateEvent(event mybot.AuEventState) {
	_ = mockobj.HandleStateEventFunc()
	recordFuncCall(&mockobj.funcCallWithArgs, "HandleStateEvent", event)
}
