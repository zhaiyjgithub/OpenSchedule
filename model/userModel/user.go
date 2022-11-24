/**
 * @author zhaiyuanji
 * @date 2022年04月08日 10:38 上午
 */
package userModel

import "time"

// User [...]
type User struct {
	ID            int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"id"`
	FirstName     string    `gorm:"column:first_name;type:varchar(50)" json:"firstName"`
	LastName      string    `gorm:"column:last_name;type:varchar(50)" json:"lastName"`
	Gender        string    `gorm:"column:gender;type:char(1)" json:"gender"`
	Email         string    `gorm:"column:email;type:varchar(100)" json:"email"`
	Birthday      string    `gorm:"column:birthday;type:varchar(10)" json:"birthday"`
	Phone         string    `gorm:"column:phone;type:varchar(20)" json:"phone"`
	StreetAddress string    `gorm:"column:street_address;type:text" json:"streetAddress"`
	City          string    `gorm:"column:city;type:varchar(50)" json:"city"`
	State         string    `gorm:"column:state;type:varchar(50)" json:"state"`
	Password      string    `gorm:"column:password;type:varchar(32)" json:"password"`
	Zip           string    `gorm:"column:zip;type:varchar(10)" json:"zip"`
	Suit          string    `gorm:"column:suit;type:varchar(100)" json:"suit"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime" json:"-"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime" json:"-"`
}

func (d *User) TableName() string {
	return "users"
}
