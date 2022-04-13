package user

import (
	"OpenSchedule/src/model/user"
	"errors"
	"gorm.io/gorm"
)

type Dao struct {
	mainEngine *gorm.DB
}

func NewUserDao(engine *gorm.DB) *Dao {
	return &Dao{mainEngine: engine}
}

func (d *Dao) CreateUser(user user.User) (error, *user.User) {
	existUser := d.GetUserByEmail(user.Email)
	if existUser.Email == user.Email {
		errText := user.Email + " is exist."
		return errors.New(errText), nil
	}
	db := d.mainEngine.Create(&user)
	return db.Error, &user
}

func (d *Dao) GetUserByEmail(email string) user.User {
	var user user.User
	d.mainEngine.Where("email = ?", email).First(&user)
	return user
}
