package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	User     string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/web", func(ctx *gin.Context) {
		// 获取浏览器请求
		// age := ctx.Query("age")
		age := ctx.DefaultQuery("age", "18")
		name, ok := ctx.GetQuery("query")
		if !ok {
			name = "somebody"
		}
		ctx.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBind(&u) //打包变量结构
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			// c.JSON(200, gin.H{
			// 	"username": u.User,
			// 	"password": u.Password,
			// })
			c.HTML(200, "index.html", gin.H{
				"Name":     u.User,
				"Password": u.Password,
			})
		}
	})
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(200, "upload.html", gin.H{})
	})
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			files := form.File["files"]
			for index, file := range files {
				log.Println(file.Filename)
				dst := fmt.Sprintf("./%d_%s", index, file.Filename)
				c.SaveUploadedFile(file, dst)
			}
			c.JSON(200, gin.H{
				"status": "ok",
			})
		}
	})

	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(200, gin.H{
			"year":  year,
			"month": month,
		})
	})

	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "http://www.qq.com")
	})
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //重定向到b请求
		r.HandleContext(c)        //重新处理上下文
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "b",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "http://www.qq.com"})
	})

	r.Run(":8088")

}
