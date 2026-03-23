package response

type BookingResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	CoachID   uint   `json:"coach_id"`
	StartTime string `json:"start_time"`
}
