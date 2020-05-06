package main

import (
	"errors"
	"strings"
)

type user struct {
	UserName string `json:user_name`
	PassWord string `json:-`
}

var userList = []user{
	user{UserName: "user1", PassWord: "pass1"},
	user{UserName: "user2", PassWord: "pass2"},
	user{UserName: "user3", PassWord: "pass3"},
}

func isUserValid(userName, passWord string) bool {
	for _, u := range userList {
		if u.UserName == userName && u.PassWord == passWord {
			return true
		}
	}
	return false
}

func registerNewUser(userName, passWord string) (*user, error) {
	if strings.TrimSpace(passWord) == "" {
		return nil, errors.New("password can't be empty")
	} else if !isUsernameAvailable(userName) {
		return nil, errors.New("username is not available")
	}

	u := user{UserName: userName, PassWord: passWord}
	userList = append(userList, u)
	return &u, nil
}

func isUsernameAvailable(userName string) bool {
	for _, u := range userList {
		if u.UserName == userName {
			return false
		}
	}
	return true
}
