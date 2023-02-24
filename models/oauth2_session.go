package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Oauth2Session struct {
	ID        string    `gorm:"primaryKey;type:uuid;index"`
	UserId    string    `gorm:"index"`
	User      User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsActive  bool      `gorm:"column:is_active;default:true"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp with time zone;"`
}

func (Oauth2Session) TableName() string {
	return "oauth2_session"
}

func (session *Oauth2Session) BeforeCreate(tx *gorm.DB) error {
	session.ID = uuid.NewV4().String()
	return nil
}
