package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func index(c *gin.Context) {
	t, err := template.New("index.tmpl").
		// Delims("{[", "]}").
		ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")
	if err != nil {
		return
	}
	msg := "这是INDEX页面"
	t.ExecuteTemplate(c.Writer, "index.tmpl", msg)
}

func home(c *gin.Context) {
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home.tmpl")
	if err != nil {
		return
	}
	msg := "这是home页面"
	t.ExecuteTemplate(c.Writer, "home.tmpl", msg)

}

func demo1(c *gin.Context) {
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		return
	}
	name := "小王子"
	t.Execute(c.Writer, name)
}

func f1(c *gin.Context) {
	k := func(name string) (string, error) {
		return name + "ok", nil
	}
	t := template.New("f.tmpl")
	t.Funcs(template.FuncMap{ // 给模板添加函数
		"kua": k,
	})
	_, err := t.ParseFiles("./f.tmpl") //解析模板
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "小王子"
	t.Execute(c.Writer, name)
}

func sayHello(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"message": "hello world!",
	// })
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v", err)
		return
	}
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}

	m1 := map[string]interface{}{
		"name":   "小王子",
		"gender": "男",
		"age":    18,
	}

	hobbyList := []string{
		"篮球",
		"足球",
		"排球",
	}
	err = t.Execute(c.Writer, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
	if err != nil {
		fmt.Printf("render template failed, err:%v", err)
		return
	}
}

func main() {
	//创建一个默认路由引擎
	r := gin.Default()
	// GET:请求方式； /hello: 请求路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.GET("/hello", sayHello)
	r.GET("/demo1", demo1)
	r.GET("/f1", f1)
	r.GET("/index", index)
	r.GET("/home", home)

	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	//启动服务
	r.Run()
}
