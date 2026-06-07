package link

import (
	"adv_demo/pkg/db"
	"errors"

	"gorm.io/gorm"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

var (
	ErrNotFound = errors.New("record not found")
)

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.First(&link, "hash = ?", hash)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, result.Error
	}
	return &link, nil
}
