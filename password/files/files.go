package files

import (
	"fmt"
	"os"
)

type FileDB struct {
	filename string
}

func NewFileDB(name string) *FileDB {
	return &FileDB{
		filename: name,
	}
}

func (db *FileDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *FileDB) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
