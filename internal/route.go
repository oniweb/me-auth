package internal

import (
	"github.com/gorilla/mux"
	"github.com/hfleury/me-auth/internal/appauth"
)

type AppRoute struct {
	appAuth appauth.AppAuth
}

func NewAppRoute(
	appauth appauth.AppAuth,
) *AppRoute {
	return &AppRoute{
		appAuth: appauth,
	}
}

func (ar *AppRoute) AppRouteConfig() *mux.Router {
	r := mux.NewRouter()
	// Only match if the domain is the AUTH_HOST_ADDRESS env var
	sbrouter := r.Host(ar.appAuth.EnvVars["AUTH_HOST_ADDRESS"]).Subrouter()

	return sbrouter
}
