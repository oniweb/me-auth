package apppsql

import (
	"fmt"

	"github.com/hfleury/me-auth/internal/appauth"
)

type AppPsql struct {
	AppAuth *appauth.AppAuth
}

func NewAppPsql(
	AppAuth *appauth.AppAuth,
) *AppPsql {
	return &AppPsql{
		AppAuth: AppAuth,
	}
}

func (ap *AppPsql) ConnPsql() string {
	usn := ap.AppAuth.EnvVars["AUTH_POSTGRESQL_USER"]
	pass := ap.AppAuth.EnvVars["AUTH_POSTGRESQL_PASS"]
	host := ap.AppAuth.EnvVars["AUTH_POSTGRESQL_HOST"]
	port := ap.AppAuth.EnvVars["AUTH_POSTGRESQL_PORT"]
	dbname := ap.AppAuth.EnvVars["AUTH_POSTGRESQL_DB"]

	strConn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", usn, pass, host, port, dbname)

	return strConn
}
