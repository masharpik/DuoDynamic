package repository

import (
	"fmt"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
)

type Repository struct {
	Users     *[]ds.User
	Cities    *[]ds.City
	UserToken *map[uint]string
}

func New() *Repository {
	var userDS []ds.User
	var cityDS []ds.City
	tokenDS := make(map[uint]string)
	r := Repository{&userDS, &cityDS, &tokenDS}
	return &r
}

func (r *Repository) AddUser(user *ds.User) error {
	userId := len(*r.Users)
	user.UserId = uint(userId)

	if err := r.CheckUsername(user.Username); err != nil {

		return err
	}

	if err := r.CheckEmail(user.Email); err != nil {

		return err
	}

	if err := CheckAge(user.Age); err != nil {
		return err
	}

	usersDB := *r.Users
	usersDB = append(usersDB, *user)
	r.Users = &usersDB

	return nil
}

func (r *Repository) SaveToken(userId uint, token string) {
	tokenUser := *r.UserToken
	tokenUser[userId] = token
	r.UserToken = &tokenUser
}

func (r *Repository) LoginEmail(emailInp string, passwordInp string) (userId uint, err error) {
	var userPassword string

	for _, user := range *r.Users {
		if user.Email == emailInp {
			userPassword = user.Password
			userId = user.UserId
			break
		}
	}
	switch userPassword {
	case "":
		err = fmt.Errorf("cant find user with such email")
		return
	case passwordInp:
		return
	}

	err = fmt.Errorf("password is not correct")
	return
}

func (r *Repository) LoginUsername(usernameInp string, passwordInp string) (userId uint, err error) {
	var userPassword string
	for _, u := range *r.Users {
		if u.Username == usernameInp {
			userPassword = u.Password
			userId = u.UserId
		}
	}

	switch userPassword {
	case "":
		err = fmt.Errorf("cant find user with such username")
		return
	case passwordInp:
		return
	}

	err = fmt.Errorf("password is not correct")
	return
}

func (r *Repository) DeleteToken(token string) error {
	var userId uint
	flagFound := false
	for indexUser, tokenDS := range *r.UserToken {
		if tokenDS == token {
			userId = indexUser
			flagFound = true
			break
		}
	}

	if !flagFound {
		return fmt.Errorf("UnAuthorised")
	}

	delete(*r.UserToken, userId)
	return nil
}
