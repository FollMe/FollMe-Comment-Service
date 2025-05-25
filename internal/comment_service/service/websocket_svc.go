package service

import (
	"encoding/json"
	"follme/comment-service/internal/comment_service/domain"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	wsToken string
}

func NewWebSocketService(wsToken string) *WebSocketService {
	go func() {
		message := domain.Message{
			Action: domain.Ping,
		}
		for {
			for _, connection := range connPool {
				log.Print("Ping")
				connection.conn.WriteJSON(message)
			}
			time.Sleep(30 * time.Second)
		}
	}()
	return &WebSocketService{
		wsToken: wsToken,
	}
}

type connect struct {
	userId string
	conn   *websocket.Conn
	postId string
}

var connPool map[string]*connect = map[string]*connect{}

var _ domain.WebSocketSvc = &WebSocketService{}

func (s WebSocketService) HandleConnection(ws *websocket.Conn) {
	connId := string(uuid.New().String())

	ws.SetReadDeadline(time.Now().Add(time.Duration(2) * time.Second))

	for {
		var message domain.Message
		err := ws.ReadJSON(&message)
		if err != nil {
			delete(connPool, connId)
			log.Printf("error occurred: %v", err)
			break
		}
		log.Println(message)

		if ok := s.authenticate(connId, ws, &message); !ok || message.Action == domain.Authenticate {
			continue
		}

		switch message.Action {
		case domain.Join:
			connPool[connId].userId = message.Message
		case domain.JoinPost:
			connPool[connId].postId = message.Message
		case domain.TypingCmtPost:
			broadCastTyping(connPool[connId].postId, connPool[connId].userId)
		case domain.RecoverState:
			recoverState(connId, message.Message)
		}
	}
}

func (s WebSocketService) BroadCastToPostRoom(cmt *domain.Comment) {
	log.Println("Do broadcast")
	out, err := json.Marshal(domain.CommentRes{
		ID:       cmt.ID(),
		Content:  cmt.Content(),
		ParentID: cmt.ParentID(),
		Author:   cmt.Author(),
	})
	if err != nil {
		return
	}
	message := domain.Message{
		Action:  domain.Commented,
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
	message := domain.Message{
		Action: domain.TypingCmtPost,
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
	actions := domain.ClientAction{}
	err := json.Unmarshal([]byte(message), &actions)
	if err != nil {
		return
	}
	connPool[connId].userId = actions.Join
	connPool[connId].postId = actions.JoinPost
}

func (s WebSocketService) authenticate(connId string, ws *websocket.Conn, message *domain.Message) bool {
	if message.Action == domain.Authenticate && message.Message == s.wsToken {
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
