package account

import (
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+~`")

type account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
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
		login:    login,
		password: pasword,
		url:      urlString,
	}
	if acc.password == "" {
		acc.generatePassword(10)
	}
	return &acc, nil
}

func NewAccountWithTimeStamp(login, pasword, urlString string) (*AccountWithTimeStamp, error) {
	acc, err := newAccount(login, pasword, urlString)
	if err != nil {
		return nil, err
	}

	accWithTimeStamp := AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account:   *acc,
	}

	if accWithTimeStamp.password == "" {
		accWithTimeStamp.generatePassword(10)
	}
	return &accWithTimeStamp, nil
}

func (acc *AccountWithTimeStamp) OutputPassword() {
	color.Cyan("%s %s %s", acc.login, acc.password, acc.url)
	fmt.Println(acc)
}

func (acc *account) generatePassword(lenght int) {
	pass := make([]rune, lenght)
	for i := range lenght {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(pass)
}
