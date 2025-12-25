package initialize

import (
	"fmt"
	"net/http"
	"server/global"
	"server/model"

	"github.com/gin-gonic/gin"
)

// 获取TODO数据
func GetAllTodo() (todoList []*model.Todo, err error) {
	if err = global.GVA_DB.Debug().Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodoList(c *gin.Context) {
	todoList, err := GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		for i, t := range todoList {
			fmt.Printf("[%d] pointer=%p, value=%+v\n", i, t, *t)
		}
		c.JSON(http.StatusOK, todoList)
	}
}

// 获取Msg_Data数据
func GetAllMsg() (msgList []*model.Msg_Data, err error) {
	if err = global.GVA_DB.Debug().Find(&msgList).Error; err != nil {
		return nil, err
	}
	return
}

func GetMsgList(c *gin.Context) {
	msgList, err := GetAllMsg()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		for i, t := range msgList {
			fmt.Printf("[%d] pointer=%p, value=%+v\n", i, t, *t)
		}
		c.JSON(http.StatusOK, msgList)
	}
}
