package routers

import (
	"PromisedLandLab/handlers"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func RegisterRoutes(app *iris.Application, db *gorm.DB, sessionManager *sessions.Sessions) {
	// 创建 handlers 实例
	userHandlers := handlers.NewUserHandlers(db, sessionManager)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to the User System</h1>")
	})

	app.Get("/register", userHandlers.RegisterPage)
	app.Post("/register", userHandlers.Register)
	app.Get("/login", userHandlers.LoginPage)
	app.Post("/login", userHandlers.Login)
	app.Get("/profile", userHandlers.Profile)
}
