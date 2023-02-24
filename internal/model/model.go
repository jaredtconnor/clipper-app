package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Define the structure of a webclipping object
type Clipping struct {
	ID       uuid.UUID `gorm:"type:uuid"`
	Title    string    `json:"title"`
	URL      string    `json:"url"`
	Contents string    `json:"description"`
}
