package userinfo

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	GetAccountInfo(string) (*Account, error)
}

type repoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repoImpl{
		db: db,
	}
}

func (r *repoImpl) CreateAccount(account *Account) error {
	return r.db.Create(account).Error
}

func (r *repoImpl) UpdateAccount(account *Account) error {
	return r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(account).Error
}

func (r *repoImpl) GetAccountInfo(username string) (*Account, error) {
	account := &Account{}
	err := r.db.Model(&Account{}).Where("username = ?", username).First(account).Error
	return account, err
}
