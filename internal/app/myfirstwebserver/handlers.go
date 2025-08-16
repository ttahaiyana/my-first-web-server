package myfirstwebserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"github.com/ttahaiyana/my-first-web-server/internal/app/myfirstwebserver/middleware"
	"github.com/ttahaiyana/my-first-web-server/internal/app/myfirstwebserver/models"
)

func initHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func newMessage(statusCode int, message string, isError bool) *Message {
	return &Message{
		StatusCode: statusCode,
		Message:    message,
		IsError:    isError,
	}
}

func (a *API) respondWithLog(w http.ResponseWriter, statusCode int, logMsg string, clientMsg string, isError bool, err error) {
	w.WriteHeader(statusCode)

	if err != nil {
		a.logger.Info(logMsg, err)
	} else {
		a.logger.Info(logMsg)
	}

	msg := newMessage(statusCode, clientMsg, isError)
	if err = json.NewEncoder(w).Encode(msg); err != nil {
		a.logger.Info("Failed to encode JSON response:", err)
	}
}

func (a *API) GetArticleById(w http.ResponseWriter, req *http.Request) {
	initHandler(w)
	a.logger.Info("Get article by id GET api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		a.respondWithLog(w, 400,
			"Troubles while parsing {id} param. err: ",
			"Unappropriate id. Don't use id as uncasting to int value.",
			true, err)
		return
	}

	article, ok, err := a.storage.Article().FindByID(id)
	if err != nil {
		a.respondWithLog(w, 500,
			"Error while accessing database table with id. err: ",
			"We have troubles accessing database. Try again later.",
			true, err)
		return
	}

	if !ok {
		a.respondWithLog(w, 404,
			"Can not find article with that id in database.",
			"Article with that id does not exists in database.",
			true, err)
		return
	}

	w.WriteHeader(200)
	if err = json.NewEncoder(w).Encode(article); err != nil {
		a.logger.Info("Failed to encode JSON response:", err)
	}
}

func (a *API) GetAllArticles(w http.ResponseWriter, req *http.Request) {
	initHandler(w)
	a.logger.Info("Get all articles GET api/v1/articles")

	articles, err := a.storage.Article().SelectAll()
	if err != nil {
		a.respondWithLog(w, 500,
			"Error while Article().SelectAll:",
			"We have troubles accessing database. Try again later.",
			true, err)
		return
	}

	w.WriteHeader(200)
	if err = json.NewEncoder(w).Encode(articles); err != nil {
		a.logger.Info("Failed to encode JSON response:", err)
	}
}

func (a *API) CreateArticle(w http.ResponseWriter, req *http.Request) {
	initHandler(w)
	a.logger.Info("Create article POST api/v1/articles")

	var article models.Article
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		a.respondWithLog(w, 400,
			"Invalid json received from client.",
			"Provided json is invalid.",
			true, nil)
		return
	}

	_, err = a.storage.Article().Create(&article)
	if err != nil {
		a.respondWithLog(w, 500,
			"Error while Article().Create()",
			"We have troubles accessing database. Try again later.",
			true, err)
		return
	}

	w.WriteHeader(200)
	if err = json.NewEncoder(w).Encode(article); err != nil {
		a.logger.Info("Failed to encode JSON response:", err)
	}
}

func (a *API) DeleteArticle(w http.ResponseWriter, req *http.Request) {
	initHandler(w)
	a.logger.Info("Delete article by id DELETE api/v1/articles/{id}")

	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		a.respondWithLog(w, 400,
			"Troubles with parsing id param: ",
			"Unappropriate value. Can not use uncasting to int id.",
			true, err)
		return
	}

	_, ok, err := a.storage.Article().FindByID(id)
	if err != nil {
		a.respondWithLog(w, 500,
			"Trouble with accessing database with that id. err: ",
			"We have troubles accessing database. Try again later.",
			true, err)
		return
	}

	if !ok {
		a.respondWithLog(w, 404,
			"Can not find article with that id in database.",
			"Article with that id does not exist in database.",
			true, nil)
		return
	}

	_, err = a.storage.Article().Delete(id)
	if err != nil {
		a.respondWithLog(w, 500,
			"Trouble with deleting article with that id. err: ",
			"We have troubles accessing database. Try again later.",
			true, err)
		return
	}

	w.WriteHeader(202)
	msg := newMessage(202, fmt.Sprintf("Article with ID %d successfully deleted.", id), false)
	if err = json.NewEncoder(w).Encode(msg); err != nil {
		a.logger.Info("Failed to encode JSON response:", err)
	}
}

func (a *API) UpdateArticle(w http.ResponseWriter, req *http.Request) {
	initHandler(w)
	a.logger.Info("Update article PUT api/v1/articles")

	var article models.Article
	err := json.NewDecoder(req.Body).Decode(&article)
	if err != nil {
		a.respondWithLog(w, 400,
			"Invalid json received from client.",
			"Provided json is invalid.",
			true, nil)
		return
	}

	_, err = a.storage.Article().Update(&article)
	if err != nil {
		a.respondWithLog(w, 500,
			"Error while Article().Update(). err: ",
			"We have troubles accessing database. Try again later.",
			true, err)
		return
	}

	w.WriteHeader(202)
	msg := newMessage(202, "Article successfully updated.", false)
	if err = json.NewEncoder(w).Encode(msg); err != nil {
		a.logger.Info("Failed to encode JSON response:", err)
	}
}

func (a *API) CreateUser(w http.ResponseWriter, req *http.Request) {
	initHandler(w)
	a.logger.Info("Create user POST api/v1/users")

	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		a.respondWithLog(w, 400,
			"Invalid json received from client.",
			"Provided json is invalid.",
			true, nil)
		return
	}

	_, err = a.storage.User().Create(&user)
	if err != nil {
		a.respondWithLog(w, 500,
			"Error while User().Create()",
			"We have troubles accessing database. Try again later.",
			true, err)
		return
	}

	w.WriteHeader(200)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		a.logger.Info("Failed to encode JSON response:", err)
	}
}

func (a *API) DeleteAllUsers(w http.ResponseWriter, req *http.Request) {
	initHandler(w)
	a.logger.Info("Delete all users DELETE api/v1/users")

	err := a.storage.User().DeleteAll()
	if err != nil {
		a.respondWithLog(w, 500,
			"Trouble with deleting all users. err: ",
			"We have troubles with accessing database. Try again later.",
			true, err)
		return
	}

	w.WriteHeader(202)
	msg := newMessage(202, "All users successfully deleted.", false)
	if err = json.NewEncoder(w).Encode(msg); err != nil {
		a.logger.Info("Failed to encode JSON response:", err)
	}
}

func (a *API) PostToAuth(w http.ResponseWriter, req *http.Request) {
	initHandler(w)
	a.logger.Info("User Authorization POST api/v1/users/auth")

	var userFromJson models.User
	err := json.NewDecoder(req.Body).Decode(&userFromJson)
	if err != nil {
		a.respondWithLog(w, 400,
			"Invalid json received from client.",
			"Provided json is invalid.",
			true, nil)
		return
	}

	userFromDB, ok, err := a.storage.User().FindByLogin(userFromJson.Login)
	if err != nil {
		a.respondWithLog(w, 500,
			"Trouble with accessing database with that login. err: ",
			"We have troubles accessing database. Try again later.",
			true, err)
		return
	}

	if !ok {
		a.respondWithLog(w, 404,
			"Can not find article with that login in database.",
			"User with that login does not exist in database.",
			true, nil)
		return
	}

	if userFromJson.Password != userFromDB.Password {
		a.respondWithLog(w, 404,
			"Invalid credentials to auth.",
			"Your password is invalid.",
			true, nil)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		a.respondWithLog(w, 500,
			"Can not claim jwt token. err: ",
			"We have some troubles. Try again later.",
			true, err)
		return
	}

	a.respondWithLog(w, 201,
		"", tokenString,
		false, nil)
}
