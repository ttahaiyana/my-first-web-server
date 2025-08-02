package myfirstwebserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (a *API) configureLogger() error {
	logLevel, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(logLevel)
	return nil
}

func (a *API) configureRouter() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! It's me"))
	})
}
