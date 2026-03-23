package request

type CreateAvailabilityRequest struct {
	CoachID   uint   `json:"coach_id" validate:"required"`
	Day       string `json:"day" validate:"required"`
	StartTime string `json:"start_time" validate:"required"`
	EndTime   string `json:"end_time" validate:"required"`
}
