package userModel

import "time"

type UserInsurance struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"id"`
	UserID    int       `gorm:"column:user_id;type:int" json:"userID"`
	Type      int8    `gorm:"column:type;type:tinyint" json:"type"`
	PlanID    string    `gorm:"column:plan_id;type:varchar(50)" json:"planID"`
	MemberID  string    `gorm:"column:member_id;type:varchar(100)" json:"memberID"`
	Photo     string    `gorm:"column:photo;type:varchar(100)" json:"photo"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"-"`
}

func (m *UserInsurance) TableName() string {
	return "user_insurances"
}
