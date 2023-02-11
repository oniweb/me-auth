package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-oauth2/oauth2/v4/manage"

	"github.com/hfleury/me-auth/internal"
	"github.com/hfleury/me-auth/internal/appauth"
	"github.com/hfleury/me-auth/internal/apppsql"
	pgx4 "github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

func main() {
	ctx := context.Background()
	authApp := appauth.NewAppAuth()
	appPsql := apppsql.NewAppPsql(authApp)

	connConfig := appPsql.ConnPsql()
	pgxConn, err := pgx4.Connect(ctx, connConfig)
	if err != nil {
		log.Fatalf("error connecting to db - error: %v", err)
		panic(err)
	}
	defer pgxConn.Close(ctx)

	manager := manage.NewDefaultManager()

	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, _ := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	defer tokenStore.Close()

	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)
	clientStore, _ := pg.NewClientStore(adapter)

	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(clientStore)

	// setting server
	appServer := internal.NewAppServer(manager)
	srv, err := appServer.AppServerConfig()
	if err != nil {
		log.Fatalf("error setting server - error: %v", err)
		panic(err)
	}

	// setting route
	aar := appauth.NewAppAuthRoute(srv)
	aar.AppAuthRouteConfig()

	log.Fatal(http.ListenAndServe(":9096", nil))
}
