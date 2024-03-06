package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/database"
	"log"
	"net/http"
	"strconv"
)

type ApiHandler struct {
}

func (a ApiHandler) Leaderboard(w http.ResponseWriter, r *http.Request) {
	songIdParam := r.URL.Query().Get("songId")
	difficultyParam := r.URL.Query().Get("difficulty")
	songId, err := strconv.ParseUint(songIdParam, 10, 16)
	if err != nil {
		http.Error(w, "Invalid songId", http.StatusBadRequest)
		return
	}
	difficulty, err := strconv.ParseUint(difficultyParam, 10, 8)
	if err != nil {
		http.Error(w, "Invalid difficulty", http.StatusBadRequest)
		return
	}
	leaderboard, err := database.GetLeaderboard(uint(songId), uint(difficulty))
	if err != nil {
		http.Error(w, "Error getting leaderboard", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboard)
}

func (a ApiHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
	if err != nil {
		http.Error(w, "Invalid baid", http.StatusBadRequest)
		return
	}
	profile, err := database.GetPublicProfile(uint(id))
	if err != nil {
		http.Error(w, "Error getting profile", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
