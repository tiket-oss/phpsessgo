package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tiket-oss/phpsessgo"
)

var sessionManager phpsessgo.SessionManager

func main() {

	sessionManager = phpsessgo.NewSessionManager(phpsessgo.SessionManagerConfig{
		Expiration:     time.Hour,
		CookiePath:     "/",
		CookieHttpOnly: true,
		CookieDomain:   "localhost:1323",
		CookieSecure:   false,
	})

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handlerFunc, sessionHandlerFunc)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func handlerFunc(ctx echo.Context) error {
	val := ctx.Get("session")
	if val == nil {
		return ctx.String(http.StatusInternalServerError, ctx.Get("session_error").(string))
	}

	session := val.(*phpsessgo.Session)
	session.Value["some-key"] = "somevalue"
	return ctx.String(http.StatusOK, session.SessionID)
}

func sessionHandlerFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := sessionManager.Start(c.Response().Writer, c.Request())
		if err != nil {
			c.Set("session_error", err.Error())
		} else {
			c.Set("session", session)
			c.Response().Before(func() {
				sessionManager.Save(session)
			})
		}
		return next(c)
	}
}
