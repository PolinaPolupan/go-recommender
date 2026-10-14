package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/recommendations", recommendationsHandler)

	fmt.Println("Server listening on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func recommendationsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	limit := 10

	if value := r.URL.Query().Get("limit"); value != "" {
		parsedLimit, err := strconv.Atoi(value)
		if err != nil || parsedLimit <= 0 {
			http.Error(w, "invalid limit", http.StatusBadRequest)
			return
		}

		limit = parsedLimit
	}

	if limit > 100 {
		http.Error(w, "limit must be <= 100", http.StatusBadRequest)
		return
	}

	history := user_history[userID]

	recommendations := recommend(items, history)

	if limit < len(recommendations) {
		recommendations = recommendations[:limit]
	}

	response := RecommendationResponse{
		UserID: userID,
		Items:  recommendations,
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
