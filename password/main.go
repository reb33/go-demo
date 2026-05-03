package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	createAccount()
}

func createAccount() {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	myAcc.OutputPassword()
	file, err := myAcc.ToBytes()
	if err != nil {
		fmt.Println("Не не удалось преобоазовать в JSON", err)
		return
	}
	fmt.Println(string(file))
	files.WriteFile(file, "account.json")
}

func promtData(prompt string) string {
	var input string
	fmt.Print(prompt, ": ")
	fmt.Scanln(&input)
	return input
}
