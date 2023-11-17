package userinfo

import "time"

type AccountStatus string

const (
	ACTIVE     AccountStatus = "ACTIVE"
	DEACTIVATE AccountStatus = "DEACTIVATE"
	BANNED     AccountStatus = "BANNED"
)

type Account struct {
	Id        uint32
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
}
