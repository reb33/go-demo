package account

import (
	"demo/password/encrypter"
	"demo/password/output"
	"strings"
	"time"

	"github.com/fatih/color"
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

type Format interface {
	ToBytes(*Vault) ([]byte, error)
	FromBytes([]byte) (*Vault, error)
}



type Vault struct {
	Accounts  []*Account `json:"accounts"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type VaultWithDB struct {
	Vault
	db Db
	Format
	enc *encrypter.Encrypter
}

func NewVault(db Db, format Format, enc *encrypter.Encrypter) *VaultWithDB {
	encryptedData, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []*Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
			Format: format,
			enc: enc,
		}
	}
	data := enc.Decrypt(encryptedData)
	vault, err := format.FromBytes(data)
	color.Cyan("Найдено %d аккаунтов", len(vault.Accounts))
	if err != nil {
		// color.Red("Ошибка при чтении файла: %v", err)
		output.PrintError("Ошибка при чтении файла: " + err.Error())
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []*Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
			Format: format,
			enc: enc,
		}
	}
	return &VaultWithDB{
		Vault: *vault,
		db:    db,
		Format: format,
		enc: enc,
	}
}

func (vault *VaultWithDB) AddAccount(account *Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.WriteToFile()
}

func (vault *VaultWithDB) WriteToFile() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Format.ToBytes(&vault.Vault)
	if err != nil {
		// color.Red("Ошибка при записи файла: %v", err)
		output.PrintError("Ошибка форматирования данных: " + err.Error())
	}
	encryptedData := vault.enc.Encrypt(data)
	vault.db.Write(encryptedData)
}

// func (vault *Vault) ToBytes() ([]byte, error) {
// 	file, err := json.Marshal(vault)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return file, nil
// }

func (vault *VaultWithDB) FindAccounts(str string, checker func(*Account, string) bool) []*Account {
	var accounts []*Account
	for _, acc := range vault.Accounts {
		if checker(acc, str) {
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
