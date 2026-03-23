package models

import "time"

type Booking struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	CoachID   uint      `json:"coach_id"`
	StartTime time.Time `json:"start_time"`

	// Prevent double booking
	// Unique constraint
	_ struct{} `gorm:"uniqueIndex:idx_booking"`
}
