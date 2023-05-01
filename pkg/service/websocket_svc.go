package service

import (
	"encoding/json"
	"follme/comment-service/pkg/adapter/serializer"
	"follme/comment-service/pkg/model"
	"log"

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

	for {
		var message model.Message
		err := ws.ReadJSON(&message)
		if err != nil {
			delete(connPool, connId)
			log.Printf("error occurred: %v", err)
			break
		}

		log.Println(message)
		switch message.Action {
		case model.Join:
			connPool[connId].userId = message.Message
		case model.JoinPost:
			connPool[connId].postId = message.Message
		case model.TypingCmtPost:
			// TODO: here
		}

		// send message from server
		if err := ws.WriteJSON(message); err != nil {
			log.Printf("error occurred: %v", err)
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
