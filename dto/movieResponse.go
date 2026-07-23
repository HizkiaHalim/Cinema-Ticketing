package dto

type MovieDetailResponse struct {
	ID          uint               `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Poster      string             `json:"poster"`
	Showtimes   []ShowtimeResponse `json:"showtimes"`
}

type ShowtimeResponse struct {
	ID       uint   `json:"id"`
	Time     string `json:"time"`
	SeatLeft int    `json:"seat_left"`
	IsFull   bool   `json:"is_full"`
}
