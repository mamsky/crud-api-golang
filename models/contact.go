package models

import "time"

type Contact struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateContactRequest struct {
	Name   string `json:"name" validate:"required" example:"buki9"`
	Gender string `json:"gender" validate:"required,oneof=male female" example:"female"`
	Phone  string `json:"phone" validate:"required" example:"6281234567892"`
	Email  string `json:"email" validate:"required,email" example:"fulans@email.com"`
}
