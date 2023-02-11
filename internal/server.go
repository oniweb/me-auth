package internal

import (
	"log"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
)

type AppServer struct {
	manager *manage.Manager
}

func NewAppServer(
	manage *manage.Manager,
) *AppServer {
	return &AppServer{
		manager: manage,
	}
}

func (as *AppServer) AppServerConfig() (*server.Server, error) {
	srv := server.NewDefaultServer(as.manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return srv, nil
}
