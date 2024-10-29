package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	Id        uint `json:"Id" gorm:"primarykey"` //主键
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
