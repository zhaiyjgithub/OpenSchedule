package userDao

import (
	"OpenSchedule/model/userModel"
	"gorm.io/gorm/clause"
)

func (d *Dao) GetUserInsurance(userID int) ([]userModel.UserInsurance, error) {
	var ins []userModel.UserInsurance
	db := d.db.Where("user_id = ?", userID).Find(&ins)
	return ins, db.Error
}

func (d *Dao) UpdateUserInsurance(ins []userModel.UserInsurance) error {
	for _, in := range ins {
		db := d.db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"plan_id", "member_id", "photo"}),
		}).Create(&in)
		if db.Error != nil {
			return db.Error
		}
	}
	return nil
}
