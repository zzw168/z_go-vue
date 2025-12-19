package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.Static("/static", "./statics")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})
	r.GET("/index1.html", func(c *gin.Context) {
		c.HTML(200, "index1.html", nil)
	})
	r.GET("/landing-page.html", func(c *gin.Context) {
		c.HTML(200, "landing-page.html", nil)
	})
	r.GET("/about.html", func(c *gin.Context) {
		c.HTML(200, "about.html", nil)
	})
	r.GET("/shop.html", func(c *gin.Context) {
		c.HTML(200, "shop.html", nil)
	})
	r.GET("/shop-details.html", func(c *gin.Context) {
		c.HTML(200, "shop-details.html", nil)
	})
	r.GET("/blog.html", func(c *gin.Context) {
		c.HTML(200, "blog.html", nil)
	})
	r.GET("/blog-details.html", func(c *gin.Context) {
		c.HTML(200, "blog-details.html", nil)
	})
	r.GET("/faq.html", func(c *gin.Context) {
		c.HTML(200, "faq.html", nil)
	})
	r.GET("/contact.html", func(c *gin.Context) {
		c.HTML(200, "contact.html", nil)
	})
	r.GET("/404.html", func(c *gin.Context) {
		c.HTML(200, "404.html", nil)
	})
	r.GET("/cart.html", func(c *gin.Context) {
		c.HTML(200, "cart.html", nil)
	})
	r.GET("/checkout.html", func(c *gin.Context) {
		c.HTML(200, "checkout.html", nil)
	})
	r.GET("/embed.html", func(c *gin.Context) {
		c.HTML(200, "embed.html", nil)
	})
	r.GET("/users", func(c *gin.Context) {
		c.HTML(200, "/users/index.tmpl", gin.H{
			"title": "http://www.qq.com",
		})
	})
	r.GET("/posts", func(c *gin.Context) {
		c.HTML(200, "/posts/index.tmpl", gin.H{
			"title": "http://www.163.com",
		})
	})
	r.Run()
}
