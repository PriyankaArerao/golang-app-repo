package models

type Availability struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	CoachID   uint   `json:"coach_id"`
	Day       string `json:"day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
