package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ClassManage struct {
	gorm.Model
	ClassName   string 		`gorm:"type: varchar(128); not null; column:class_name"`
	ClassTime   time.Time  	`gorm:"type: integer; not null; column:class_time"`
	ClassVis    int  		`gorm:"type: integer; not null; column:class_vis"`
	ClassVie    int  		`gorm:"type: integer; not null; column:class_vie"`
	ClassDesc   string 		`gorm:"type: varchar(3000); not null; column:class_desc"`
	ClassIcon   string 		`gorm:"type: varchar(512); not null; column:class_icon"`
	ClassStatus int    		`gorm:"type: integer; not null; column:class_status"`
}

type ClassSerializer struct {
	ID       			uint        `json:"id"`
	CreateAt 			time.Time   `json:"create_at"`
	UpdateAt 			time.Time   `json:"update_at"`
	ClassName 			string		`json:"class_name"`
	ClassTime 			time.Time	`json:"class_time"`
	ClassVis 			int			`json:"class_vis"`
	ClassVie 			int			`json:"class_vie"`
	ClassDesc 			string		`json:"class_desc"`
	ClassIcon 			string		`json:"class_icon"`
	ClassStatus 		int			`json:"class_status"`
}

func (c *ClassManage) Serializer() ClassSerializer {
	return ClassSerializer{
		ID:			c.ID,
		CreateAt:   c.CreatedAt,
		UpdateAt:   c.UpdatedAt,
		ClassName:  c.ClassName,
		ClassTime:  c.ClassTime,
		ClassVis:   c.ClassVis,
		ClassVie:   c.ClassVie,
		ClassDesc:  c.ClassDesc,
		ClassIcon:  c.ClassIcon,
		ClassStatus:c.ClassStatus,
	}
}