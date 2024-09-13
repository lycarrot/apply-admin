package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	ID        uint `json:"ID" gorm:"primarykey"` //主键
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
