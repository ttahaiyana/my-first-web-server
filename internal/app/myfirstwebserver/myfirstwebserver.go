package myfirstwebserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ttahaiyana/my-first-web-server/storage"
)

type API struct {
	config  Config
	logger  *logrus.Logger
	router  mux.Router
	storage *storage.Storage
}

func New(config Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: *mux.NewRouter(),
	}
}

func (a *API) Start() error {
	err := a.configureLogger()
	if err != nil {
		return err
	}

	a.configureRouter()

	err = a.configureStorage()
	if err != nil {
		return err
	}

	a.logger.Info("starting http server at port", a.config.BindAddr)

	return http.ListenAndServe(a.config.BindAddr, &a.router)
}
