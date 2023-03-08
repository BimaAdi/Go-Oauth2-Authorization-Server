package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Oauth2Token struct {
	ID        string    `gorm:"primaryKey;type:uuid;index"`
	UserId    string    `gorm:"column:user_id;index;"`
	User      User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Code      string    `gorm:"column:code;type:varchar;not null;index"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp with time zone;"`
}

func (Oauth2Token) TableName() string {
	return "oauth2_token"
}

func (session *Oauth2Token) BeforeCreate(tx *gorm.DB) error {
	session.ID = uuid.NewV4().String()
	return nil
}
