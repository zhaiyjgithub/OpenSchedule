package userDao

import (
	"OpenSchedule/model/userModel"
	"errors"
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

func NewUserDao(engine *gorm.DB) *Dao {
	return &Dao{db: engine}
}

func (d *Dao) CreateUser(user userModel.User) (error, *userModel.User) {
	existUser := d.GetUserByEmail(user.Email)
	if existUser.Email == user.Email {
		errText := user.Email + " is exist."
		return errors.New(errText), nil
	}
	db := d.db.Create(&user)
	return db.Error, &user
}

func (d *Dao) GetUserByEmail(email string) userModel.User {
	var user userModel.User
	d.db.Where("email = ?", email).First(&user)
	return user
}

func (d *Dao) CreateSubUser(subUser userModel.SubUsers) error {
	db := d.db.Create(&subUser)
	return db.Error
}

func (d *Dao) GetSubUsers(userId int) []userModel.SubUsers {
	var subUsers []userModel.SubUsers
	_ = d.db.Model(&userModel.SubUsers{}).Where("user_id = ?", userId).Find(&subUsers)
	return subUsers
}

func (d *Dao) DeleteSubUser(subUserID int) error {
	u := userModel.SubUsers{ID: subUserID}
	db := d.db.Delete(&u)
	return db.Error
}

func (d *Dao) UpdateSubUserPhone(subUserId int, phone string) error {
	u := userModel.SubUsers{
		ID: subUserId,
	}
	db := d.db.Model(&u).Update("phone", phone)
	return db.Error
}

func (d *Dao) UpdateUserProfile(userProfile userModel.UserProfile) error {
	var fields []string
	u := userModel.User{
		ID: userProfile.UserID,
	}
	if userProfile.Email != nil {
		u.Email = *userProfile.Email
		fields = append(fields, "Email")
	}
	if userProfile.Phone != nil {
		u.Phone = *userProfile.Phone
		fields = append(fields, "Phone")
	}
	if userProfile.Birthday != nil {
		u.Birthday = *userProfile.Birthday
		fields = append(fields, "Birthday")
	}
	if userProfile.Gender != nil {
		u.Gender = *userProfile.Gender
		fields = append(fields, "Gender")
	}
	if userProfile.StreetAddress != nil {
		u.StreetAddress = *userProfile.StreetAddress
		fields = append(fields, "StreetAddress")
	}
	if userProfile.Suit != nil {
		u.Suit = *userProfile.Suit
		fields = append(fields, "Suit")
	}
	if userProfile.City != nil {
		u.City = *userProfile.City
		fields = append(fields, "City")
	}
	if userProfile.State != nil {
		u.State = *userProfile.State
		fields = append(fields, "State")
	}
	if userProfile.Zip != nil {
		u.Zip = *userProfile.Zip
		fields = append(fields, "Zip")
	}
	db := d.db.Model(&u).Select(fields).Updates(u)
	return db.Error
}
