package apppsql

import (
	"fmt"
	"strconv"

	"github.com/hfleury/me-auth/internal/appauth"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
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

func (ap *AppPsql) ConPsql() (pgx.ConnConfig, error) {
	ui64, err := strconv.ParseUint(ap.AppAuth.EnvVars["AUTH_POSTGRESQL_PORT"], 10, 64)
	if err != nil {
		fmt.Printf("Error to parse db port string to uint16")
		return pgx.ConnConfig{}, err
	}
	portui16 := uint16(ui64)

	logLevel, err := strconv.Atoi(ap.AppAuth.EnvVars["AUTH_POSTGRESQL_LOGLEVEL"])
	if err != nil {
		fmt.Printf("Error to parse log level string to int")
		return pgx.ConnConfig{}, err
	}

	pgConfig := pgx.ConnConfig{
		Config: pgconn.Config{
			Host:     ap.AppAuth.EnvVars["AUTH_POSTGRESQL_HOST"],
			Port:     portui16,
			Database: ap.AppAuth.EnvVars["AUTH_POSTGRESQL_DB"],
			User:     ap.AppAuth.EnvVars["AUTH_POSTGRESQL_USER"],
			Password: ap.AppAuth.EnvVars["AUTH_POSTGRESQL_PASS"],
		},
		LogLevel: pgx.LogLevel(logLevel),
	}

	return pgConfig, nil
}
