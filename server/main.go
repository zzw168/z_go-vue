package main

import (
	"server/core"
	"server/global"
	"server/initialize"

	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// 这部分 @Tag 设置用于排序, 需要排序的接口请按照下面的格式添加
// swag init 对 @Tag 只会从入口文件解析, 默认 main.go
// 也可通过 --generalInfo flag 指定其他文件
// @Tag.Name        Base
// @Tag.Name        SysUser
// @Tag.Description 用户

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.8.6
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	// 初始化系统
	initializeSystem()
	// 运行服务器
	core.RunServer()
}

// initializeSystem 初始化系统所有组件
// 提取为单独函数以便于系统重载时调用
func initializeSystem() {
	global.GVA_VP = core.Viper() // 初始化Viper配置管理库
	initialize.OtherInit()
	// OtherInit() 启动期初始化函数，
	// 用于校验 JWT 时间配置、初始化 JWT 黑名单缓存，
	// 并在未配置的情况下自动从 go.mod 中识别项目的 Go module 名称，
	// 为后续鉴权与代码生成提供基础环境。
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()                // 初始化定时任务
	initialize.DBList()               // 初始化多数据库连接
	initialize.SetupHandlers()        // 注册全局函数
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
	}
}
