/**
 * @author zhaiyuanji
 * @date 2022年06月11日 8:37 上午
 */
package user

// SubUsers [...]
// SubUsers [...]
type SubUsers struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(11);not null" json:"-"`
	FirstName string `gorm:"column:first_name;type:varchar(50)" json:"firstName"`
	LastName  string `gorm:"column:last_name;type:varchar(50)" json:"lastName"`
	Email     string `gorm:"column:email;type:varchar(50)" json:"email"`
	Phone     string `gorm:"column:phone;type:varchar(20)" json:"phone"`
	Birthday  string `gorm:"column:birthday;type:char(10)" json:"birthday"`
	Gender    string `gorm:"column:gender;type:char(1)" json:"gender"`
	UserID    int    `gorm:"column:user_id;type:int(11)" json:"userId"`
	IsLegal   bool   `gorm:"column:is_legal;type:tinyint(1)" json:"isLegal"`
}

func (d *SubUsers) TableName() string {
	return "sub_users"
}
