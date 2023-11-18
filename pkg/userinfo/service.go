package userinfo

import (
	"errors"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
	"time"
)

type Service interface {
	Register(account *Account) error
	Login(string, string) error
	UpdateAccountInfo(*Account) error
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) Register(account *Account) error {
	if _, err := s.repository.GetAccountInfo(account.Username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return s.repository.CreateAccount(&Account{
				Username:  account.Username,
				Password:  account.Password,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Status:    string(ACTIVE),
			})
		}
		return err
	}
	return xerrors.Errorf("username existed")
}

func (s *serviceImpl) Login(username, password string) error {

	account, err := s.repository.GetAccountInfo(username)
	if err != nil {
		return err
	}
	if account.Password != password {
		return xerrors.New("error login: incorrect password")
	}
	return nil
}

func (s *serviceImpl) UpdateAccountInfo(account *Account) error {
	account.UpdatedAt = time.Now()
	return s.repository.UpdateAccount(account)
}
