package service

import (
	"encoding/json"
	"follme/comment-service/pkg/adapter/serializer"
	"follme/comment-service/pkg/model"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebSocketService struct{}

func NewWebSocketService() *WebSocketService {
	return &WebSocketService{}
}

type connect struct {
	userId string
	conn   *websocket.Conn
	postId string
}

var connPool map[string]*connect = map[string]*connect{}

var _ model.WebSocketSvc = &WebSocketService{}

func (s WebSocketService) HandleConnection(ws *websocket.Conn) {
	connId := string(uuid.New().String())
	connPool[connId] = &connect{
		conn: ws,
	}

	ws.SetReadDeadline(time.Now().Add(15 * time.Second))
	for {
		var message model.Message
		err := ws.ReadJSON(&message)
		if err != nil {
			delete(connPool, connId)
			log.Printf("error occurred: %v", err)
			break
		}
		log.Println(message)
		ws.SetReadDeadline(time.Now().Add(15 * time.Second))
		switch message.Action {
		case model.Join:
			connPool[connId].userId = message.Message
		case model.JoinPost:
			connPool[connId].postId = message.Message
		case model.TypingCmtPost:
			broadCastTyping(connPool[connId].postId, connPool[connId].userId)
		case model.RecoverState:
			recoverState(connId, message.Message)
		}
	}
}

func (s WebSocketService) BroadCastToPostRoom(cmt *model.Comment) {
	log.Println("Do broadcast")
	out, err := json.Marshal(serializer.Comment{
		ID:       cmt.ID(),
		Content:  cmt.Content(),
		ParentID: cmt.ParentID(),
		Author:   cmt.Author(),
	})
	if err != nil {
		return
	}
	message := model.Message{
		Action:  model.Commented,
		Message: string(out),
	}
	for _, connect := range connPool {
		if connect.postId != cmt.PostSlug() {
			continue
		}

		// send message from server
		if err := connect.conn.WriteJSON(message); err != nil {
			log.Printf("error occurred: %v", err)
		}
	}
}

func broadCastTyping(postId string, emitter string) {
	message := model.Message{
		Action: model.TypingCmtPost,
	}
	for _, connect := range connPool {
		if connect.userId == emitter || connect.postId != postId {
			continue
		}

		// send message from server
		if err := connect.conn.WriteJSON(message); err != nil {
			log.Printf("error occurred: %v", err)
		}
	}
}

func recoverState(connId string, message string) {
	actions := model.ClientAction{}
	err := json.Unmarshal([]byte(message), &actions)
	if err != nil {
		return
	}
	connPool[connId].userId = actions.Join
	connPool[connId].postId = actions.JoinPost
}
