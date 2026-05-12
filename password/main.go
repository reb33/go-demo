package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountsURL,
	"3": findAccountsLogin,
	"4": deleteAccount,
}

func menuCounter() func() { // демонстрация closure
	i := 0
	return func() {
		i++
		fmt.Println(i)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти env файл")
	}
	format := &account.JsonFormat{}
	// vault := account.NewVault(files.NewFileDB("account.json"), format, encrypter.NewEncrypter())
	// format := &account.EncodingFormat{
	// 	Enc: encrypter.NewEncrypter(),
	// }
	vault := account.NewVault(files.NewFileDB("vault.txt"), format, encrypter.NewEncrypter())
	counter := menuCounter()
mainloop:
	for {
		counter()
		promt := promtData(
			"1. Создать аккаунт",
			"2. Найти аккаунт по url",
			"3. Найти аккаунт по login",
			"4. Удалить аккаунт",
			"5. Выход",
			"",
		)
		menuFunc := menu[promt]
		if menuFunc == nil {
			break mainloop
		}
		menuFunc(vault)
	}

}

func findAccountsURL(vault *account.VaultWithDB) {
	findAccounts(vault, "url", func(acc *account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
}
func findAccountsLogin(vault *account.VaultWithDB) {
	findAccounts(vault, "login", func(acc *account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
}

func findAccounts(vault *account.VaultWithDB, field string, checker func(*account.Account, string) bool) {
	value := promtData("Введите " + field + " для поиска аккаунта")

	accounts := vault.FindAccounts(value, checker)
	if accounts == nil {
		color.Black("Нет аккаунта с %s %s", field, value)
	}
	for i, acc := range accounts {
		fmt.Print(i+1, ". ")
		acc.Output()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := promtData("Введите url для удаления аккаунта")
	acc := vault.DelAccount(url)
	if acc == nil {
		// color.Black("Нет аккаунта с url %s", url)
		output.PrintError("Нет аккаунта с url" + url)
		return
	}
	fmt.Print("Удален аккаунт: ")
	color.Magenta("%s %s %s", acc.Login, acc.Password, acc.Url)
}

func createAccount(vault *account.VaultWithDB) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAcc, err := account.NewAccount(login, password, url)
	if err != nil {
		// fmt.Println(err)
		output.PrintError(err)
	}
	myAcc.Output()
	vault.AddAccount(myAcc)
}

func promtData(prompts ...any) string {
	var input string
	last := len(prompts) - 1
	for i, prompt := range prompts {
		if i == last {
			fmt.Print(prompt, ": ")
			continue
		}
		fmt.Println(prompt)
	}
	fmt.Scanln(&input)
	return input
}
