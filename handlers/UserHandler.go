package handlers

import (
	"PromisedLandLab/statics"
	"PromisedLandLab/utils"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type UserHandlers struct {
	DB             *gorm.DB
	SessionManager *sessions.Sessions
}

func NewUserHandlers(db *gorm.DB, sessionManager *sessions.Sessions) *UserHandlers {
	return &UserHandlers{DB: db, SessionManager: sessionManager}
}

func (uh *UserHandlers) RegisterPage(ctx iris.Context) {
	ctx.View("register.html")
}

func (uh *UserHandlers) Register(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")

	var existingUser statics.User
	uh.DB.Where("username = ?", username).First(&existingUser)

	if existingUser.ID != 0 {
		ctx.ViewData("Error", "Username already exists")
		ctx.View("register.html")
		return
	}

	hashedPassword, _ := utils.HashPassword(password)

	user := statics.User{Username: username, Password: hashedPassword}
	uh.DB.Create(&user)

	ctx.Redirect("/login")
}

func (uh *UserHandlers) LoginPage(ctx iris.Context) {
	ctx.View("login.html")
}

func (uh *UserHandlers) Login(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")

	var user statics.User
	uh.DB.Where("username = ?", username).First(&user)

	if user.ID == 0 || !utils.CheckPassword(password, user.Password) {
		ctx.Redirect("/login")
		return
	}

	session := uh.SessionManager.Start(ctx)
	session.Set("userID", user.ID)

	ctx.Redirect("/profile")
}

func (uh *UserHandlers) Profile(ctx iris.Context) {
	session := uh.SessionManager.Start(ctx)
	userID := session.Get("userID")

	if userID == nil {
		ctx.Redirect("/login")
		return
	}

	var user statics.User
	uh.DB.First(&user, userID)

	ctx.ViewData("user", user)
	ctx.View("profile.html")
}
