package services

type MovieListRequest struct {
	SearchDate string `json:"searchDate" binding:"required"`
}
