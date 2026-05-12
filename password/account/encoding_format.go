package account

import (
	"bytes"
	"demo/password/encrypter"
	"time"
)

// Код нерабочий, ширует неправильно

type EncodingFormat struct{
	Enc *encrypter.Encrypter
}

func (e *EncodingFormat) ToBytes(vault *Vault) ([]byte, error){
	data := []byte{}
	for i, account := range vault.Accounts {
		if i > 0 {
			data = append(data, "\n"...)
		}
		createdAt, err := account.CreatedAt.MarshalText()
		if err != nil {
			return nil, err
		}
		updatedAt, err := account.UpdatedAt.MarshalText()
		if err != nil {
			return nil, err
		}
		accountLine := bytes.Join([][]byte{
			[]byte(account.Login),
			[]byte(account.Url),
			[]byte(createdAt),
			[]byte(updatedAt),
			e.Enc.Encrypt([]byte(account.Password)),
		}, []byte(","))
		data = append(data, accountLine...)
	}
	return data, nil
}

func (e *EncodingFormat) FromBytes(data []byte) (*Vault, error){
	vault := Vault{} // empty 
	for line := range bytes.SplitSeq(data, []byte("\n")) {
		values := bytes.Split(line, []byte(","))
		createdAt, err := time.Parse(time.RFC3339Nano, string(values[2]))
		if err != nil {
			return nil, err
		}
		updatedAt, err := time.Parse(time.RFC3339Nano, string(values[3]))
		if err != nil {
			return nil, err
		}
		account := Account{
			Login: string(values[0]),
			Url: string(values[1]),
			Password: string(e.Enc.Decrypt(values[4])),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		vault.Accounts = append(vault.Accounts, &account)
	}
	return &vault, nil
}