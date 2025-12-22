package main

import (
	"zzw168/Go_Bubble/dao"
	"zzw168/Go_Bubble/routers"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
		// fmt.Println(err)
	}
	defer dao.Close() // 程序退出关闭数据库链接
	// 模型绑定
	dao.InitModel()
	// dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	r.Run(":8088")
}
