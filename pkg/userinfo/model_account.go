package userinfo

import "time"

type AccountStatus string

const (
	ACTIVE     AccountStatus = "ACTIVE"
	DEACTIVATE AccountStatus = "DEACTIVATE"
	BANNED     AccountStatus = "BANNED"
)

type Account struct {
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
}
