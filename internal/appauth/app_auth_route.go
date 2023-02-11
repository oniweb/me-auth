package appauth

import (
	"encoding/json"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/google/uuid"
	"github.com/hfleury/me-auth/pkg/appauth/middleware"
)

type AppAuthRoute struct {
	server *server.Server
	http   *http.Server
}

func NewAppAuthRoute(
	server *server.Server,
) *AppAuthRoute {
	return &AppAuthRoute{
		server: server,
	}
}

func (ar *AppAuthRoute) AppAuthRouteConfig() {
	http.HandleFunc("/protected", middleware.ValidateToken(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected"))
	}, ar.server))

	http.HandleFunc("/credentials", func(w http.ResponseWriter, r *http.Request) {
		clientId := uuid.New().String()
		clientSecret := uuid.New().String()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		ar.server.HandleTokenRequest(w, r)
	})
}
