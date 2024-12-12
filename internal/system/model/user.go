package model

import "time"

type (
	User struct {
		Id          int64
		Name        string
		Password    string
		Role        string
		Phone       string
		Email       string
		Vip         int
		EnabledFlag int
		Created_at  time.Time
		Updated_at  time.Time
		Deleted_at  time.Time
	}
)
