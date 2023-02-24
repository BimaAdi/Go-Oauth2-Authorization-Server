package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Oauth2Session struct {
	ID           string    `gorm:"primaryKey;type:uuid;index"`
	UserId       string    `gorm:"column:user_id;index;"`
	User         User      `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ClientID     string    `gorm:"column:client_id;not null;"`
	ClientSecret string    `gorm:"column:client_secret;not null;"`
	IsActive     bool      `gorm:"column:is_active;default:true"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp with time zone;"`
}

func (Oauth2Session) TableName() string {
	return "oauth2_session"
}

func (session *Oauth2Session) BeforeCreate(tx *gorm.DB) error {
	session.ID = uuid.NewV4().String()
	return nil
}
