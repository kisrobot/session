package manager

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/kisrobot/middlewares"
	"github.com/kisrobot/session"
	"github.com/kisrobot/session/gorilla"
)

// SessionManager default session manager
var SessionManager session.ManagerInterface = gorilla.New("_session", sessions.NewCookieStore([]byte("secret")))

func init() {
	middlewares.Use(middlewares.Middleware{
		Name: "session",
		Handler: func(handler http.Handler) http.Handler {
			return SessionManager.Middleware(handler)
		},
	})
}
