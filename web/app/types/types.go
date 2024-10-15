package types

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name       string
	GivenName  string
	FamilyName string
	Nickname   string
	Picture    string
	UpdatedAt  string
}
