package link_test

import (
	"adv_demo/configs"
	"adv_demo/internal/link"
	"adv_demo/pkg/db"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var repo *link.LinkRepository

func TestMain(m *testing.M) {
	// Инициализация перед всеми тестами
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("не удалось загрузить .env", err)
	}
	conf := &configs.Config{
		Db: configs.DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: configs.AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
	db := db.NewDb(conf)

	repo = link.NewLinkRepository(db)

	// Запуск всех тестов
	m.Run()  //обязательная команда, а то тесты не запустяться
}

func TestRepositoryCheckHashNotFound(t *testing.T) {
	isExist, err := repo.IsHashExist("123")
	if err != nil {
		t.Error(err)	
	}
	if isExist {
		t.Error("hash должен быть не найден")
	}
}
