package entities

import "github.com/google/uuid"

type Products struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Brand       string    `json:"brand"`
	Visible     bool      `json:"visible"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Weight      float64   `json:"weight"`
}
