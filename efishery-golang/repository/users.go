package repository

import (
	"efishery-golang/entity"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(DB *gorm.DB) User {
	return User{
		db: DB,
	}
}

func (repo User) FindByPhonePassword(phone string, password string) (entity.Users, error) {
	var data entity.Users
	err := repo.db.
		Where("phone", phone).
		Where("password", password).
		Find(&data).
		Error
	return data, err
}

func (repo User) FindByPhone(phone string) (entity.Users, error) {
	var data entity.Users
	err := repo.db.
		Where("phone", phone).
		Find(&data).
		Error
	return data, err
}

func (repo User) StoreUser(user entity.Users) (entity.Users, error) {
	err := repo.db.Create(&user).
		Error
	return user, err
}
