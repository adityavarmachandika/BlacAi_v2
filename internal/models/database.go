package models

import (
	"time"

	"github.com/google/uuid"
)

type UserDetails struct {

	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email string `gorm:"unique;not null"`
	PhoneNumber string
	FirstName string `gorm:"not null"`
	LastName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthProviderDetails struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserId uuid.UUID `gorm:"type:uuid"`
	ProviderId *string `gorm:"unique"`
	HashedPassword string
	ProviderName string
    User   UserDetails `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}