package dao

import (
	"fmt"
	"zzw168/Go_Bubble/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySql() (err error) {
	dsn := "root:root@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func InitModel() {
	// 模型绑定
	DB.AutoMigrate(&models.Todo{})
}

func Close() {
	DB.Close()
}

/*
	dao 负责所有增删改查操作
*/

// CreateATodo 创建Todo
func CreateATodo(todo *models.Todo) (err error) {
	err = DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*models.Todo, err error) {
	if err = DB.Debug().Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodoBy(id int) (*models.Todo, error) {
	fmt.Println(id)
	var todo models.Todo
	err := DB.Debug().First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, err
}

func UpdateATodo(todo *models.Todo) (err error) {
	err = DB.Save(&todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = DB.Where("id=?", id).Delete(&models.Todo{}).Error
	return
}
