/**
 * @author zhaiyuanji
 * @date 2022年04月08日 10:40 上午
 */
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

func (d *Dao) CreateUser(user user.Users) error {
	existUser := d.GetUserByEmail(user.Email)
	if existUser.Email == user.Email {
		errText := user.Email + " is exist."
		return errors.New(errText)
	}
	db := d.mainEngine.Create(&user)
	return db.Error
}

func (d *Dao) GetUserByEmail(email string) user.Users {
	var user user.Users
	d.mainEngine.Where("email = ?", email).First(&user)
	return user
}
