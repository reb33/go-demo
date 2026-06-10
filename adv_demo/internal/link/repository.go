package link

import (
	"adv_demo/pkg/db"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repo *LinkRepository) IsHashExist(hash string) (bool, error) {
	var link Link
	result := repo.Database.Limit(1).Find(&link, "hash = ?", hash)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (repo *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := repo.Database.Limit(1).Find(&link, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

type UpdateLinkParams struct {
	ID   uint64
	Url  string
	Hash string
}

func (repo *LinkRepository) Update(params *UpdateLinkParams) (*Link, error) {
	link := &Link{
		Model: gorm.Model{ID: uint(params.ID)},
		Url:   params.Url,
		Hash:  params.Hash,
	}
	result := repo.Database.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(id uint) (*Link, bool, error) {
	var link Link
	result := repo.Database.Clauses(clause.Returning{}).Delete(&link, id)
	if result.Error != nil {
		return nil, false, result.Error
	}
	fmt.Println(result.RowsAffected)
	return &link, result.RowsAffected>0, nil
}
