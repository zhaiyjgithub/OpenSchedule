/**
 * @author zhaiyuanji
 * @date 2022年03月31日 2:26 上午
 */
package doctor

import "time"

// Clinicals [...]
type	Clinicals	struct {
	ID	uint32	`gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"-"`
	Npi	int	`gorm:"index:index_npi;column:npi;type:int(20)" json:"-"`
	Name	string	`gorm:"column:name;type:varchar(255)" json:"name"`
	Desc	string	`gorm:"column:desc;type:varchar(255)" json:"desc"`
	CreatedAt	time.Time	`gorm:"column:created_at;type:datetime" json:"-"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;type:datetime" json:"-"`
}

func (s *Clinicals) TableName() string {
	return "clinicals"
}