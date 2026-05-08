package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJsonDB("account.json"))
mainloop:
	for {
		promt := getMenu()
		switch promt {
		case 1:
			vault = createAccount(vault)
		case 2:
			findAccounts(vault)
		case 3:
			vault = deleteAccount(vault)
		default:
			break mainloop
		}
	}

}

func getMenu() int {
	promt := 0
	fmt.Println()
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&promt)
	return promt
}

func findAccounts(vault *account.VaultWithDB){
	url := promtData("Введите url для поиска аккаунта")
	
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
	url := promtData("Введите url для удаления аккаунта")
	acc := vault.DelAccount(url)
	if acc == nil {
		color.Black("Нет аккаунта с url %s", url)
		return vault
	}
	fmt.Print("Удален аккаунт: ")
	color.Magenta("%s %s %s", acc.Login, acc.Password, acc.Url)
	return vault
}

func createAccount(vault *account.VaultWithDB) (*account.VaultWithDB){
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return vault
	}
	myAcc.Output()
	vault.AddAccount(myAcc)
	return vault
}

func promtData(prompt string) string {
	var input string
	fmt.Print(prompt, ": ")
	fmt.Scanln(&input)
	return input
}
