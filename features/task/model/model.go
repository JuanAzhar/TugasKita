package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:varchar(50);primaryKey;not null" json:"id"`
	AdminId     string
	Title       string `gorm:"not null" json:"title"`
	Description string 
	Point       int  
	Message     string 
	Status      string `gorm:"type:enum('Aktif', 'Melewati Tenggat');default:'Aktif'"`
	Start_date  time.Time
	End_date    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
