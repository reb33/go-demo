package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []*Account `json:"accounts"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func NewVault() *Vault {
	data, err := files.ReadFile("account.json")
	if err != nil {
		return &Vault{
			Accounts:  []*Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(data, &vault)
	if err != nil {
		color.Red("Ошибка при чтении файла: %v", err)
		return &Vault{
			Accounts:  []*Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) AddAccount(account *Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.WriteToFile()
}

func (vault *Vault) WriteToFile() {
	vault.UpdatedAt = time.Now()
	data, err := json.Marshal(vault)
	if err != nil {
		color.Red("Ошибка при записи файла: %v", err)
	}
	files.WriteFile(data, "account.json")
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) FindAccounts(url string) []*Account {
	var accounts []*Account
	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			accounts = append(accounts, acc)
		}
	}
	return accounts
}

func (vault *Vault) FindAccountWithPosition(url string) (*Account, int) {
	for i, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			return acc, i
		}
	}
	return nil, -1
}

func (vault *Vault) DelAccount(url string) *Account{
	acc, pos := vault.FindAccountWithPosition(url)
	if acc == nil {
		return nil
	}
	vault.Accounts = append(vault.Accounts[:pos], vault.Accounts[pos+1:]...)
	vault.WriteToFile()
	return acc
}
