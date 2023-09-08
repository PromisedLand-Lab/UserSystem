package main

import (
	"PromisedLandLab/routers"
	"PromisedLandLab/utils"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12/sessions"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

var db *gorm.DB
var sessionManager *sessions.Sessions

func main() {
	// 初始化数据库连接和会话管理器等
	db = utils.InitDB()
	sessionManager = utils.InitSessionManager()

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	// 注册模板
	app.RegisterView(iris.HTML("./views", ".html"))
	// 注册路由
	routers.RegisterRoutes(app, db, sessionManager)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
