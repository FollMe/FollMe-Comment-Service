package domain

import (
	"github.com/gorilla/websocket"
)

type WebSocketSvc interface {
	HandleConnection(ws *websocket.Conn)
	BroadCastToPostRoom(cmt *Comment)
}

const (
	Authenticate  = "authenticate"
	Join          = "join"
	JoinPost      = "join_post"
	TypingCmtPost = "typing_cmt_post"
	Commented     = "commented"
	RecoverState  = "recover_state"
	Ping          = "ping"
)

type Message struct {
	UserID  string `json:"userId"`
	Action  string `json:"action"`
	Message string `json:"message"`
}

type ClientAction struct {
	Join     string `json:"join,omitempty"`
	JoinPost string `json:"join_post,omitempty"`
}
