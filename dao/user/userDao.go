package user

import (
	"OpenSchedule/model/user"
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

func (d *Dao) CreateSubUser(subUser user.SubUsers) error {
	db := d.mainEngine.Create(&subUser)
	return db.Error
}

func (d *Dao) GetSubUsers(userId int) []user.SubUsers {
	var subUsers []user.SubUsers
	_ = d.mainEngine.Model(&user.SubUsers{}).Where("user_id = ?", userId).Find(&subUsers)
	return subUsers
}

func (d *Dao) DeleteSubUser(subUserID int) error {
	u := user.SubUsers{ID: subUserID}
	db := d.mainEngine.Delete(&u)
	return db.Error
}

func (d *Dao) UpdateSubUserPhone(subUserId int, phone string) error {
	u := user.SubUsers{
		ID: subUserId,
	}
	db := d.mainEngine.Model(&u).Update("phone", phone)
	return db.Error
}
