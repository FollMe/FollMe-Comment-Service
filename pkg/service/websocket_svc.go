package service

import (
	"encoding/json"
	"follme/comment-service/pkg/adapter/serializer"
	"follme/comment-service/pkg/config"
	"follme/comment-service/pkg/model"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebSocketService struct{}

func NewWebSocketService() *WebSocketService {
	go func() {
		message := model.Message{
			Action: model.Ping,
		}
		for {
			for _, connection := range connPool {
				log.Print("Ping")
				connection.conn.WriteJSON(message)
			}
			time.Sleep(30 * time.Second)
		}
	}()
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

	ws.SetReadDeadline(time.Now().Add(time.Duration(2) * time.Second))

	for {
		var message model.Message
		err := ws.ReadJSON(&message)
		if err != nil {
			delete(connPool, connId)
			log.Printf("error occurred: %v", err)
			break
		}
		log.Println(message)

		if ok := authenticate(connId, ws, &message); !ok || message.Action == model.Authenticate {
			continue
		}

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

func authenticate(connId string, ws *websocket.Conn, message *model.Message) bool {
	if message.Action == model.Authenticate && message.Message == config.AppConfig.WSToken {
		connPool[connId] = &connect{
			conn: ws,
		}
		log.Printf("connId %v authenticated successfully!", connId)
	}

	if _, ok := connPool[connId]; !ok {
		log.Printf("connId %v authenticated fail!", connId)
		return false
	}

	ws.SetReadDeadline(time.Now().Add(time.Duration(35000+rand.Intn(1000)) * time.Millisecond))
	return true
}
