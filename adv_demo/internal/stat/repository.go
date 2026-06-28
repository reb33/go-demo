package stat

import (
	"adv_demo/pkg/db"
	"log"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type StatDB struct {
	gorm.Model
	LinkId uint
	Clicks int
	Data   datatypes.Date
}

func (StatDB) TableName() string {
	return "stats"
}

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{Db: db}
}

func (r *StatRepository) AddClick(linkId uint) {
	// если нет статы за сегодня по ссылке, то создаем ее
	// если есть, то инкрементируем
	var stat StatDB
	currentDate := datatypes.Date(time.Now())
	result := r.Db.Find(&stat, "link_id = ? AND data = ?", linkId, currentDate)
	if result.RowsAffected == 0 {
		stat = StatDB{
			LinkId: linkId,
			Clicks: 1,
			Data:   currentDate,
		}
		r.Db.Create(&stat)
	} else {
		stat.Clicks += 1
		r.Db.Save(&stat)
	}
	log.Printf("link_id %d, clicks %d\n", linkId, stat.Clicks)
}
