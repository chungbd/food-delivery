package appcontext

import "gorm.io/gorm"

type ctx struct {
	mainDB    *gorm.DB
	secretKey string
}

func New(mainDB *gorm.DB, secretKey string) *ctx {
	return &ctx{mainDB: mainDB, secretKey: secretKey}
}

func (c ctx) GetMainDBConnection() *gorm.DB {
	return c.mainDB
}

func (c ctx) SecretKey() string {
	return c.secretKey
}

type tokenExpiry struct {
	atExp int
	rtExp int
}

func NewTokenConfig() tokenExpiry {
	return tokenExpiry{
		atExp: 60 * 60 * 24 * 7,
		rtExp: 60 * 60 * 24 * 7 * 2,
	}
}

func (tk tokenExpiry) GetAtExp() int {
	return tk.atExp
}

func (tk tokenExpiry) GetRtExp() int {
	return tk.rtExp
}
