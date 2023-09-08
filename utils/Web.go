package utils

import "github.com/kataras/iris/v12/sessions"

func InitSessionManager() *sessions.Sessions {
	return sessions.New(sessions.Config{
		Cookie:       "session_cookie",
		AllowReclaim: true,
	})
}
