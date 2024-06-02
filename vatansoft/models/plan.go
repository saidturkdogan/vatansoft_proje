package models

import "time"

type Plan struct {
	ID        uint      `gorm:"primaryKey"`
	StudentID uint      `json:"student_id"`
	Day       string    `json:"day"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
