/**
 * @author zhaiyuanji
 * @date 2022年03月31日 11:46 下午
 */
package doctorModel

import "time"

// Insurances [...]
type Insurances struct {
	ID        uint32    `gorm:"primaryKey;column:id;type:int(11) unsigned;not null" json:"-"`
	Npi       int64     `gorm:"index:index_npi;column:npi;type:int(20)" json:"-"`
	Name      string    `gorm:"column:name;type:varchar(255)" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"-"`
}

func (s *Insurances) TableName() string {
	return "insurances"
}
