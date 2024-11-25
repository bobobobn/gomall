package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"unique_index;type:varchar(255) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "users"
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var u User

	return &u, db.Where("email =?", email).First(&u).Error
}
