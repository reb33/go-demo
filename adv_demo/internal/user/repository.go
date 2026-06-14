package user

import (
	"adv_demo/pkg/db"

	"gorm.io/gorm"
)

type UserDB struct {
	gorm.Model
	User
}

func (UserDB) TableName() string {
	return "users"
}

type UserRepository struct {
	*db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Create(user *User) (*UserDB, error) {
	userDB := &UserDB{
		User: *user,
	}
	result := r.Db.Create(userDB)
	if result.Error != nil {
		return nil, result.Error
	}
	return userDB, nil
}

func (r *UserRepository) FirstByEmail(email string) (*UserDB, error) {
	var userDB UserDB
	result := r.Db.First(&userDB, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userDB, nil
}

func (r *UserRepository) FindByEmail(email string) (*UserDB, bool, error) {
	var userDB UserDB
	result := r.Db.Limit(1).Find(&userDB, "email = ?", email)
	if result.Error != nil {
		return nil, false, result.Error
	}
	return &userDB, result.RowsAffected>0, nil
}
