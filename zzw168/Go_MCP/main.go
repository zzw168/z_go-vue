package main

import (
	"net/http"
	"zzw168/Go_MCP/server"

	"github.com/gin-gonic/gin"
)

func McpRun() *server.SSEServer {
	// 这里相当于你的初始化逻辑
	s := server.NewSSEServer()
	return s
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	sseServer := McpRun()

	r.GET("/sse", sseServer.SSEHandler)
	r.POST("/message", sseServer.MessageHandler)

	// r.StaticFile("/", "./index.html")
	r.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	http.ListenAndServe(":8080", r)
}
