package main

type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Score    float64 `json:"score"`
}

type RecommendationResponse struct {
	UserID string `json:"user_id"`
	Items  []Item `json:"items"`
}
