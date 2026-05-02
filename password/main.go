package main

import (
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+~`")

type account struct {
	login string
	password string
	url string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func newAccount(login, pasword, urlString string) (*account, error) {
	if login == "" {
		return nil, fmt.Errorf("логин не может быть пустым")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("некорректный URL: %w", err)
	}
	acc := account{
		login: login,
		password: pasword,
		url: urlString,
	}
	if acc.password == "" {
		acc.generatePassword(10)
	}
	return &acc, nil
}

func newAccountWithTimeStamp(login, pasword, urlString string) (*accountWithTimeStamp, error) {
	acc, err := newAccount(login, pasword, urlString)
	if err != nil {
		return nil, err
	}
	
	accWithTimeStamp := accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: *acc,
	}

	if accWithTimeStamp.password == "" {
		accWithTimeStamp.generatePassword(10)
	}
	return &accWithTimeStamp, nil
}



func (acc *accountWithTimeStamp) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
	fmt.Println(acc)
}

func (acc *account) generatePassword(lenght int) {
	pass := make([]rune, lenght)
	for i := range lenght {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(pass)
}

func main() {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAcc, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAcc.outputPassword()
}

func promtData(prompt string) string {
	var input string
	fmt.Print(prompt, ": ")
	fmt.Scanln(&input)
	return input
}



