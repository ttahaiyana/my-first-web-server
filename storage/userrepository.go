package storage

import (
	"fmt"
	"log"

	"github.com/ttahaiyana/my-first-web-server/internal/app/myfirstwebserver/models"
)

type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING id", tableUser)
	if err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	var founded bool
	var foundedUser *models.User
	users, err := ur.SelectAll()
	if err != nil {
		return foundedUser, founded, err
	}
	for _, user := range users {
		if user.Login == login {
			foundedUser = user
			founded = true
			break
		}
	}
	return foundedUser, founded, nil
}

func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableUser)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		u := models.User{}
		err := rows.Scan(&u.ID, &u.Login, &u.Password)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &u)
	}
	return users, nil
}

func (ur *UserRepository) DeleteAll() error {
	query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tableUser)
	_, err := ur.storage.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
