package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index in...")
	name, ok := c.Get("name")
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//计时
	start := time.Now()
	c.Next() //调用后续处理函数
	// c.Abort() //阻止调用后续的处理函数
	//return
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}

// 中间件
func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "小王子")
	fmt.Println("m2 out...")
}

// 中间件
func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if doCheck {
			fmt.Println("doCheck ok!")
			c.Next()
		} else {
			c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(m1) //全局注册中间件函数m1
	r.GET("/index", indexHandler)
	r.GET("/shop", m2, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	// userGroup := r.Group("/user", authMiddleware(true))
	userGroup := r.Group("/user")
	userGroup.Use(authMiddleware(true), m2)
	{
		userGroup.GET("/a", indexHandler)
	}

	r.Run()
}
