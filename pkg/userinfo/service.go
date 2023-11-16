package userinfo

import (
	"golang.org/x/xerrors"
	"time"
)

type Service interface {
	Register(string, string) error
	Login(string, string) error
	UpdateAccountInfo(Account) error
}

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{
		repository: repository,
	}
}

func (s *serviceImpl) Register(username, password string) error {
	account := &Account{
		Username:  username,
		Password:  HashString(password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    string(ACTIVE),
	}
	return s.repository.CreateAccount(account)
}

func (s *serviceImpl) Login(username, password string) error {

	account, err := s.repository.GetAccountInfo(username)
	if err != nil {
		return err
	}
	if account.Password != HashString(password) {
		return xerrors.New("error login: incorrect password")
	}
	return nil
}

func (s *serviceImpl) UpdateAccountInfo(account Account) error {
	account.UpdatedAt = time.Now()
	return s.repository.UpdateAccount(&account)
}
