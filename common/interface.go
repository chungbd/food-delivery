package common

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
}

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
