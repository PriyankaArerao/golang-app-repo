package request

type CreateBookingRequest struct {
	UserID  uint   `json:"user_id" validate:"required"`
	CoachID uint   `json:"coach_id" validate:"required"`
	Time    string `json:"datetime" validate:"required"`
}
