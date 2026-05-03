package account

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+~`")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(login, pasword, urlString string) (*Account, error) {
	if login == "" {
		return nil, fmt.Errorf("логин не может быть пустым")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("некорректный URL: %w", err)
	}
	acc := Account{
		Login:     login,
		Password:  pasword,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if acc.Password == "" {
		acc.generatePassword(10)
	}
	return &acc, nil
}

func (acc *Account) OutputPassword() {
	color.Cyan("%s %s %s", acc.Login, acc.Password, acc.Url)
	fmt.Println(acc)
}

func (acc *Account) ToBytes() ([]byte, error){
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) generatePassword(lenght int) {
	pass := make([]rune, lenght)
	for i := range lenght {
		pass[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(pass)
}
