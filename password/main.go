package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJsonDB("account.json"))
mainloop:
	for {
		promt := promtData([]string{"1. Создать аккаунт", "2. Найти аккаунт", "3. Удалить аккаунт", "4. Выход", ""})
		switch promt {
		case "1":
			vault = createAccount(vault)
		case "2":
			findAccounts(vault)
		case "3":
			vault = deleteAccount(vault)
		default:
			break mainloop
		}
	}

}

func findAccounts(vault *account.VaultWithDB){
	url := promtData([]string{"Введите url для поиска аккаунта"})
	
	accounts := vault.FindAccounts(url)
	if accounts == nil {
		color.Black("Нет аккаунта с url %s", url)
	}
	for i, acc := range accounts {
		fmt.Print(i+1, ". ")
		acc.Output()
	}
}

func deleteAccount(vault *account.VaultWithDB) *account.VaultWithDB{
	url := promtData([]string{"Введите url для удаления аккаунта"})
	acc := vault.DelAccount(url)
	if acc == nil {
		// color.Black("Нет аккаунта с url %s", url)
		output.PrintError("Нет аккаунта с url" + url)
		return vault
	}
	fmt.Print("Удален аккаунт: ")
	color.Magenta("%s %s %s", acc.Login, acc.Password, acc.Url)
	return vault
}

func createAccount(vault *account.VaultWithDB) (*account.VaultWithDB){
	login := promtData([]string{"Введите логин"})
	password := promtData([]string{"Введите пароль"})
	url := promtData([]string{"Введите URL"})

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		// fmt.Println(err)
		output.PrintError(err)
		return vault
	}
	myAcc.Output()
	vault.AddAccount(myAcc)
	return vault
}

func promtData[T any](prompts []T) string {
	var input string
	last := len(prompts) - 1
	for i, prompt := range prompts {
		if i == last{
			fmt.Print(prompt, ": ")
			continue
		}
		fmt.Println(prompt)
	}
	fmt.Scanln(&input)
	return input
}
