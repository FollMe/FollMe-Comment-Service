package model

import (
	"github.com/gorilla/websocket"
)

type WebSocketSvc interface {
	HandleConnection(ws *websocket.Conn)
	BroadCastToPostRoom(cmt *Comment)
}

const (
	Join          = "join"
	JoinPost      = "join_post"
	TypingCmtPost = "typing_cmt_post"
	Commented     = "commented"
)

type Message struct {
	UserID  string `json:"userId"`
	Action  string `json:"action"`
	Message string `json:"message"`
}
