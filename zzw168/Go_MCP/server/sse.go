package server

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type SSEServer struct {
	clients map[chan string]bool
	lock    sync.Mutex
}

func NewSSEServer() *SSEServer {
	return &SSEServer{
		clients: make(map[chan string]bool),
	}
}

// SSE 连接处理
func (s *SSEServer) SSEHandler(c *gin.Context) {
	ch := make(chan string)
	s.lock.Lock()
	s.clients[ch] = true
	s.lock.Unlock()

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-ch; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}

// 客户端发消息
func (s *SSEServer) MessageHandler(c *gin.Context) {
	var json struct {
		Message string `json:"message"`
	}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	go s.Broadcast(fmt.Sprintf("用户发来消息：%s", json.Message))
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// 广播消息到所有客户端
func (s *SSEServer) Broadcast(msg string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for ch := range s.clients {
		select {
		case ch <- msg:
		default:
			close(ch)
			delete(s.clients, ch)
		}
	}
}
