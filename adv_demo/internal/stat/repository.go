package stat

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type StatDB struct {
	gorm.Model
	LinkId uint
	Clicks int
	Data datatypes.Date
}

func (StatDB) TableName() string {
	return "stats"
}
