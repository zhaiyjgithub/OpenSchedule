package doctorDao

import (
	"OpenSchedule/model/doctorModel"
	"errors"
)

func (d *Dao) CreateUser(user doctorModel.DoctorUser) error {
	var count int64
	db := d.mainEngine.Model(&doctorModel.DoctorUser{}).Where("email = ? AND npi = ?", user.Email, user.Npi).Count(&count)
	if count > 0 {
		return errors.New("email has bind with this npi")
	} else {
		db.Create(&user)
	}
	return db.Error
}

func (d *Dao) GetUser(email string, password string) (doctorModel.DoctorUser, error) {
	var user doctorModel.DoctorUser
	db := d.mainEngine.Model(&doctorModel.DoctorUser{}).Where("email = ? AND password = ?", email, password).First(&user)
	return user, db.Error
}
