package Models

import "time"

type Model struct {
	ID        uint      `json:"id" gorm:"unique;autoIncrement;primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
