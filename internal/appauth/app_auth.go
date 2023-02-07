package appauth

import (
	"os"

	"github.com/joho/godotenv"
)

type AppAuth struct {
	EnvVars map[string]string
}

func NewAppAuth() *AppAuth {
	appAuth := new(AppAuth)
	appAuth.EnvVars = make(map[string]string)
	appAuth.setEnv()

	return appAuth
}

func (aa *AppAuth) setEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	aa.buildEnvVars()
	return nil
}

func (aa *AppAuth) buildEnvVars() {

	// PostgreSQL
	aa.EnvVars["AUTH_POSTGRESQL_USER"] = os.Getenv("AUTH_POSTGRESQL_USER")
	aa.EnvVars["AUTH_POSTGRESQL_PORT"] = os.Getenv("AUTH_POSTGRESQL_PORT")
	aa.EnvVars["AUTH_POSTGRESQL_DB"] = os.Getenv("AUTH_POSTGRESQL_DB")
	aa.EnvVars["AUTH_POSTGRESQL_PASS"] = os.Getenv("AUTH_POSTGRESQL_PASS")
	aa.EnvVars["AUTH_POSTGRESQL_HOST"] = os.Getenv("AUTH_POSTGRESQL_HOST")
	aa.EnvVars["AUTH_POSTGRESQL_LOGLEVEL"] = os.Getenv("AUTH_POSTGRESQL_LOGLEVEL")

	// Host
	aa.EnvVars["AUTH_HOST_PORT"] = os.Getenv("AUTH_HOST_PORT")
	aa.EnvVars["AUTH_HOST_ADDRESS"] = os.Getenv("AUTH_HOST_ADDRESS")
}
