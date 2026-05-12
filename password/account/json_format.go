package account

import "encoding/json"

type JsonFormat struct {

}

func (j *JsonFormat) ToBytes(vault *Vault) ([]byte, error){
	return json.Marshal(vault)
}

func (j *JsonFormat) FromBytes(data []byte) (*Vault, error){
	var vault Vault
	err := json.Unmarshal(data, &vault)
	return &vault, err
}