package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RadioManage struct {
	gorm.Model
	ClassId 	int				`gorm:"type: varchar(128); not null; column:class_id"`
	RadioName 	string			`gorm:"type: integer; not null; column:radio_name"`
	RadioTime 	time.Time		`gorm:"type: integer; not null; column:radio_time"`
	RadioStu 	int				`gorm:"type: integer; not null; column:radio_stu"`
	RadioVie 	int				`gorm:"type: integer; not null; column:radio_vie"`
	RadioIcon 	string			`gorm:"type: varchar(512); not null; column:radio_icon"`
	RadioStatus int				`gorm:"type: integer; not null; column:radio_status"`
}

type RadioSerializer struct {
	ID       uint      	`json:"id"`
	CreateAt time.Time 	`json:"create_at"`
	UpdateAt time.Time 	`json:"update_at"`
	ClassID  int 		`json:"class_id"`
	RadioName string 	`json:"radio_name"`
	RadioTime time.Time `json:"radio_time"`
	RadioStu int 		`json:"radio_stu"`
	RadioVie int 		`json:"radio_vie"`
	RadioIcon string 	`json:"radio_icon"`
	RadioStatus int  	`json:"radio_status"`
}

func (r *RadioManage) Serializer() RadioSerializer {
	return RadioSerializer{
		ID:			r.ID,
		CreateAt:   r.CreatedAt,
		UpdateAt:   r.UpdatedAt,
		ClassID:	r.ClassId,
		RadioName:	r.RadioName,
		RadioTime:	r.RadioTime,
		RadioStu:	r.RadioStu,
		RadioVie:	r.RadioVie,
		RadioIcon:	r.RadioIcon,
		RadioStatus:r.RadioStatus,
	}
}
