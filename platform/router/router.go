package router

import (
	"01-Login/platform/authenticator"
	"01-Login/platform/middleware"
	"01-Login/web/app/callback"
	"01-Login/web/app/data"
	"01-Login/web/app/home"
	"01-Login/web/app/login"
	"01-Login/web/app/logout"
	"01-Login/web/app/types"
	"01-Login/web/app/user"
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(types.ProfileType{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", home.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/data", middleware.IsAuthenticated, data.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
