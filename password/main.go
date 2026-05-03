package main

import (
	"fmt"
	"demo/password/account"
)

func main() {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAcc, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAcc.OutputPassword()
}

func promtData(prompt string) string {
	var input string
	fmt.Print(prompt, ": ")
	fmt.Scanln(&input)
	return input
}
