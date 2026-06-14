package user_test

import (
	"adv_demo/configs"
	"adv_demo/internal/user"
	"adv_demo/pkg/db"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var repo *user.UserRepository

func TestMain(m *testing.M) {
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

	repo = user.NewUserRepository(db)

	// Запуск всех тестов
	m.Run()
}

func TestRepositoryFindByEmailNotFound(t *testing.T) {
	user, isFound, err := repo.FindByEmail("test@test.ru")
	fmt.Println(user, isFound, err)
}
