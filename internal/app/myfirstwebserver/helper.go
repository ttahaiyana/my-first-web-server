package myfirstwebserver

import (
	"github.com/sirupsen/logrus"

	"github.com/ttahaiyana/my-first-web-server/storage"
)

var (
	prefix string = "/api/v1"
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
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles", a.CreateArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticle).Methods("DELETE")
	a.router.HandleFunc(prefix+"/atricles/{id}", a.UpdateArticle).Methods("PUT")

	a.router.HandleFunc(prefix+"/users", a.CreateUser).Methods("GET")
}

func (a *API) configureStorage() error {
	storage := storage.NewStorage(*a.config.Storage)
	if err := storage.Open(); err != nil {
		return nil
	}
	a.storage = storage
	return nil
}
