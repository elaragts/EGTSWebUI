package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/database"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/model"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/pkg"
	"log"
	"math"
	"net/http"
	"strconv"
)

type ApiHandler struct {
}

func verifyClientBaid(w http.ResponseWriter, r *http.Request) (uint, error) {
	accountBaid := uint(math.Round(r.Context().Value("baid").(float64)))
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 16)
	if err != nil {
		http.Error(w, "Invalid baid", http.StatusBadRequest)
		return 0, err
	}
	if accountBaid != uint(id) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return 0, errors.New("unauthorized")
	}
	return uint(id), nil
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

func (a ApiHandler) Datatable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pkg.Datatable)
}

func (a ApiHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}

	profile, err := database.GetPublicProfile(id)
	if err != nil {
		http.Error(w, "Error getting profile", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func (a ApiHandler) GetProfileOptions(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}

	profileOptions, err := database.GetProfileOptions(id)
	if err != nil {
		http.Error(w, "Error getting profile options", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// stored in auth db so need an extra query
	customTitleOn, err := database.GetCustomTitleOn(id)
	if err != nil {
		http.Error(w, "Error getting customTitleOn setting", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	profileOptions.CustomTitleOn = customTitleOn

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profileOptions)
}

func (a ApiHandler) GetCostumeOptions(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}
	costumeOptions, err := database.GetCostumeOptions(id)
	if err != nil {
		http.Error(w, "Error getting costume options", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(costumeOptions)
}

func (a ApiHandler) GetSongOptions(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}
	songOptions, err := database.GetSongOptions(id)
	if err != nil {
		http.Error(w, "Error getting song options", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songOptions)
}

func (a ApiHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}

	var profileSettings model.ProfileSettings
	err = json.NewDecoder(r.Body).Decode(&profileSettings)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// profileOptions data filtering
	profileOptions := profileSettings.ProfileOptions
	if len(profileOptions.MyDonName) < 1 || len(profileOptions.MyDonName) > 100 {
		http.Error(w, "Invalid Name Length (1-100)", http.StatusBadRequest)
	} else if len(profileOptions.Title) > 200 {
		http.Error(w, "Title Length Exceeds 200", http.StatusBadRequest)
	} else if profileOptions.Language < 0 || profileOptions.Language > 4 {
		http.Error(w, "Invalid Language (0-4)", http.StatusBadRequest)
	} else if profileOptions.TitlePlateId < 0 || profileOptions.TitlePlateId > 7 {
		http.Error(w, "Invalid Title Plate Id (0-7)", http.StatusBadRequest)
	} else if profileOptions.AchievementDisplayDifficulty < 0 || profileOptions.AchievementDisplayDifficulty > 5 {
		http.Error(w, "Invalid Achievement Panel Difficulty (0-5)", http.StatusBadRequest)
	} else if profileOptions.DifficultySettingCourse < 0 || profileOptions.DifficultySettingCourse > 6 {
		http.Error(w, "Invalid Course Difficulty Setting (0-6)", http.StatusBadRequest)
	} else if profileOptions.DifficultySettingStar < 0 || profileOptions.DifficultySettingStar > 11 {
		http.Error(w, "Invalid Star Difficulty Setting (0-11)", http.StatusBadRequest)
	} else if profileOptions.DifficultySettingSort < 0 || profileOptions.DifficultySettingSort > 5 {
		http.Error(w, "Invalid Sort Difficulty Setting (0-5)", http.StatusBadRequest)
	}

	// costumeOptions data filtering
	costumeOptions := profileSettings.CostumeOptions
	if costumeOptions.CurrentBody < 0 || (costumeOptions.CurrentBody >= uint(len(pkg.Datatable.Body)) && len(pkg.Datatable.Body) > 0) {
		errorMessage := fmt.Sprintf("Invalid Body (0-%d)", len(pkg.Datatable.Body))
		http.Error(w, errorMessage, http.StatusBadRequest)
	} else if costumeOptions.CurrentFace < 0 || (costumeOptions.CurrentFace >= uint(len(pkg.Datatable.Face)) && len(pkg.Datatable.Face) > 0) {
		errorMessage := fmt.Sprintf("Invalid Face (0-%d)", len(pkg.Datatable.Face))
		http.Error(w, errorMessage, http.StatusBadRequest)
	} else if costumeOptions.CurrentHead < 0 || (costumeOptions.CurrentHead >= uint(len(pkg.Datatable.Head)) && len(pkg.Datatable.Head) > 0) {
		errorMessage := fmt.Sprintf("Invalid Head (0-%d)", len(pkg.Datatable.Head))
		http.Error(w, errorMessage, http.StatusBadRequest)
	} else if costumeOptions.CurrentKigurumi < 0 || (costumeOptions.CurrentKigurumi >= uint(len(pkg.Datatable.Kigurumi)) && len(pkg.Datatable.Kigurumi) > 0) {
		errorMessage := fmt.Sprintf("Invalid Kigurumi (0-%d)", len(pkg.Datatable.Kigurumi))
		http.Error(w, errorMessage, http.StatusBadRequest)
	} else if costumeOptions.CurrentPuchi < 0 || (costumeOptions.CurrentPuchi >= uint(len(pkg.Datatable.Puchi)) && len(pkg.Datatable.Puchi) > 0) {
		errorMessage := fmt.Sprintf("Invalid Puchi (0-%d)", len(pkg.Datatable.Puchi))
		http.Error(w, errorMessage, http.StatusBadRequest)
	} else if costumeOptions.ColorBody < 0 || costumeOptions.ColorBody > 62 {
		http.Error(w, "Invalid Body Color (0-62)", http.StatusBadRequest)
	} else if costumeOptions.ColorFace < 0 || costumeOptions.ColorFace > 62 {
		http.Error(w, "Invalid Face Color (0-62)", http.StatusBadRequest)
	} else if costumeOptions.ColorLimb < 0 || costumeOptions.ColorLimb > 62 {
		http.Error(w, "Invalid Limb Color (0-62)", http.StatusBadRequest)
	}

	// songOptions data filtering
	songOptions := profileSettings.SongOptions
	if songOptions.SpeedId < 0 || songOptions.SpeedId > 14 {
		http.Error(w, "Invalid Speed Id (0-14)", http.StatusBadRequest)
	} else if songOptions.RandomId < 0 || songOptions.RandomId > 2 {
		http.Error(w, "Invalid Random Id (0-2)", http.StatusBadRequest)
	} else if songOptions.SelectedToneId < 0 || songOptions.SelectedToneId > 19 {
		http.Error(w, "Invalid Tone Id (0-19)", http.StatusBadRequest)
	} else if songOptions.NotesPosition < -5 || songOptions.NotesPosition > 5 {
		http.Error(w, "Invalid Notes Position (-5 to 5)", http.StatusBadRequest)
	}

	err = database.UpdateUser(id, profileSettings)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// customTitleOn is stored in the auth db
	err = database.UpdateCustomTitleOn(id, profileOptions.CustomTitleOn)
	if err != nil {
		http.Error(w, "Error updating customTitleOn", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
