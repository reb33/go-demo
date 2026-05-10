package account

import (
	"demo/password/output"
	"encoding/json"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []*Account `json:"accounts"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type VaultWithDB struct {
	Vault
	db Db
}

func NewVault(db Db) *VaultWithDB {
	data, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []*Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(data, &vault)
	if err != nil {
		// color.Red("Ошибка при чтении файла: %v", err)
		output.PrintError("Ошибка при чтении файла: " + err.Error())
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []*Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDB) AddAccount(account *Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.WriteToFile()
}

func (vault *VaultWithDB) WriteToFile() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		// color.Red("Ошибка при записи файла: %v", err)
		output.PrintError("Ошибка при записи файла: " + err.Error())
	}
	vault.db.Write(data)
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDB) FindAccounts(url string) []*Account {
	var accounts []*Account
	for _, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			accounts = append(accounts, acc)
		}
	}
	return accounts
}

func (vault *VaultWithDB) FindAccountWithPosition(url string) (*Account, int) {
	for i, acc := range vault.Accounts {
		if strings.Contains(acc.Url, url) {
			return acc, i
		}
	}
	return nil, -1
}

func (vault *VaultWithDB) DelAccount(url string) *Account {
	acc, pos := vault.FindAccountWithPosition(url)
	if acc == nil {
		return nil
	}
	vault.Accounts = append(vault.Accounts[:pos], vault.Accounts[pos+1:]...)
	vault.WriteToFile()
	return acc
}
