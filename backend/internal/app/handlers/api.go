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

func (a ApiHandler) Stats(w http.ResponseWriter, r *http.Request) {

	stats, err := database.GetStats()
	if err != nil {
		http.Error(w, "Error getting stats", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
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
		return
	} else if len(profileOptions.Title) > 200 {
		http.Error(w, "Title Length Exceeds 200", http.StatusBadRequest)
		return
	} else if profileOptions.Language < 0 || profileOptions.Language > 4 {
		http.Error(w, "Invalid Language (0-4)", http.StatusBadRequest)
		return
	} else if profileOptions.TitlePlateId < 0 || profileOptions.TitlePlateId > 7 {
		http.Error(w, "Invalid Title Plate Id (0-7)", http.StatusBadRequest)
		return
	} else if profileOptions.AchievementDisplayDifficulty < 0 || profileOptions.AchievementDisplayDifficulty > 5 {
		http.Error(w, "Invalid Achievement Panel Difficulty (0-5)", http.StatusBadRequest)
		return
	} else if profileOptions.DifficultySettingCourse < 0 || profileOptions.DifficultySettingCourse > 6 {
		http.Error(w, "Invalid Course Difficulty Setting (0-6)", http.StatusBadRequest)
		return
	} else if profileOptions.DifficultySettingStar < 0 || profileOptions.DifficultySettingStar > 11 {
		http.Error(w, "Invalid Star Difficulty Setting (0-11)", http.StatusBadRequest)
		return
	} else if profileOptions.DifficultySettingSort < 0 || profileOptions.DifficultySettingSort > 5 {
		http.Error(w, "Invalid Sort Difficulty Setting (0-5)", http.StatusBadRequest)
		return
	}

	// costumeOptions data filtering
	costumeOptions := profileSettings.CostumeOptions
	if costumeOptions.CurrentBody < 0 || (costumeOptions.CurrentBody >= uint(len(pkg.Datatable.Body)) && len(pkg.Datatable.Body) > 0) {
		errorMessage := fmt.Sprintf("Invalid Body (0-%d)", len(pkg.Datatable.Body))
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	} else if costumeOptions.CurrentFace < 0 || (costumeOptions.CurrentFace >= uint(len(pkg.Datatable.Face)) && len(pkg.Datatable.Face) > 0) {
		errorMessage := fmt.Sprintf("Invalid Face (0-%d)", len(pkg.Datatable.Face))
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	} else if costumeOptions.CurrentHead < 0 || (costumeOptions.CurrentHead >= uint(len(pkg.Datatable.Head)) && len(pkg.Datatable.Head) > 0) {
		errorMessage := fmt.Sprintf("Invalid Head (0-%d)", len(pkg.Datatable.Head))
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	} else if costumeOptions.CurrentKigurumi < 0 || (costumeOptions.CurrentKigurumi >= uint(len(pkg.Datatable.Kigurumi)) && len(pkg.Datatable.Kigurumi) > 0) {
		errorMessage := fmt.Sprintf("Invalid Kigurumi (0-%d)", len(pkg.Datatable.Kigurumi))
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	} else if costumeOptions.CurrentPuchi < 0 || (costumeOptions.CurrentPuchi >= uint(len(pkg.Datatable.Puchi)) && len(pkg.Datatable.Puchi) > 0) {
		errorMessage := fmt.Sprintf("Invalid Puchi (0-%d)", len(pkg.Datatable.Puchi))
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	} else if costumeOptions.ColorBody < 0 || costumeOptions.ColorBody > 62 {
		http.Error(w, "Invalid Body Color (0-62)", http.StatusBadRequest)
		return
	} else if costumeOptions.ColorFace < 0 || costumeOptions.ColorFace > 62 {
		http.Error(w, "Invalid Face Color (0-62)", http.StatusBadRequest)
		return
	} else if costumeOptions.ColorLimb < 0 || costumeOptions.ColorLimb > 62 {
		http.Error(w, "Invalid Limb Color (0-62)", http.StatusBadRequest)
		return
	}

	// songOptions data filtering
	songOptions := profileSettings.SongOptions
	if songOptions.SpeedId < 0 || songOptions.SpeedId > 14 {
		http.Error(w, "Invalid Speed Id (0-14)", http.StatusBadRequest)
		return
	} else if songOptions.RandomId < 0 || songOptions.RandomId > 2 {
		http.Error(w, "Invalid Random Id (0-2)", http.StatusBadRequest)
		return
	} else if songOptions.SelectedToneId < 0 || songOptions.SelectedToneId > 19 {
		http.Error(w, "Invalid Tone Id (0-19)", http.StatusBadRequest)
		return
	} else if songOptions.NotesPosition < -5 || songOptions.NotesPosition > 5 {
		http.Error(w, "Invalid Notes Position (-5 to 5)", http.StatusBadRequest)
		return
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

func (a ApiHandler) GetAccessCodes(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}

	accessCodes, err := database.GetAccessCodes(id)
	if err != nil {
		http.Error(w, "Error getting access codes", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// IMPORTANT - sending access codes as a SEPARATE array, NOT an object that contains an array
	json.NewEncoder(w).Encode(accessCodes)
}

func (a ApiHandler) AddAccessCode(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}

	var newAccessCode model.AccessCode
	err = json.NewDecoder(r.Body).Decode(&newAccessCode)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// length of new access code between 1-100
	if len(newAccessCode.AccessCode) < 1 || len(newAccessCode.AccessCode) > 100 {
		http.Error(w, "New access code must be between 8 and 100 characters", http.StatusBadRequest)
		return
	}

	// Checking that new access code is unique/not in use
	_, found, err := database.GetBaidFromAccessCode(newAccessCode.AccessCode)
	if err != nil {
		http.Error(w, "Error checking access code validity", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if found {
		http.Error(w, "Invalid Access Code", http.StatusBadRequest)
		return
	}

	// adding new access code to db
	err = database.AddAccessCode(id, newAccessCode.AccessCode)
	if err != nil {
		http.Error(w, "Error adding access code", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a ApiHandler) DeleteAccessCode(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}

	var deleteAccessCode model.AccessCode
	err = json.NewDecoder(r.Body).Decode(&deleteAccessCode)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Checking that the access code being deleted actually belongs to the current user
	ownerId, found, err := database.GetBaidFromAccessCode(deleteAccessCode.AccessCode)
	if err != nil {
		http.Error(w, "Error checking access code validity", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if !found || id != ownerId {
		http.Error(w, "Invalid Access Code", http.StatusBadRequest)
		return
	}

	// User MUST have at least 1 access code
	currentAccessCodes, err := database.GetAccessCodes(id)
	if err != nil {
		http.Error(w, "Error getting access code count", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if len(currentAccessCodes) == 1 {
		http.Error(w, "You must have at least one access code", http.StatusBadRequest)
		return
	}

	// deleting access code from db
	err = database.DeleteAccessCode(id, deleteAccessCode.AccessCode) // sending baid just in case
	if err != nil {
		http.Error(w, "Error deleting access code", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a ApiHandler) GetFavouritedSongs(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}

	favouritedSongs, err := database.GetFavouritedSongs(id)
	if err != nil {
		http.Error(w, "Error getting favourited songs", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// IMPORTANT - sending favourited songs as a SEPARATE array, NOT an object that contains an array
	json.NewEncoder(w).Encode(favouritedSongs)
}

func (a ApiHandler) AddFavouritedSong(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}

	var newFavouritedSong model.FavouritedSong
	err = json.NewDecoder(r.Body).Decode(&newFavouritedSong)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// check that song ID links to a valid song
	_, exists := pkg.Datatable.Song[newFavouritedSong.SongId]
	if !exists {
		http.Error(w, "Invalid Song ID", http.StatusBadRequest)
		return
	}

	// check that user doesn't already have the song favourited
	favouritedSongs, err := database.GetFavouritedSongs(id)
	if err != nil {
		http.Error(w, "Error checking favourited songs", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	for _, songId := range favouritedSongs {
		if newFavouritedSong.SongId != songId {
			continue
		}
		http.Error(w, "Song already favourited", http.StatusBadRequest)
		return
	}

	// adding new song to db
	err = database.AddFavouritedSong(id, newFavouritedSong.SongId)
	if err != nil {
		http.Error(w, "Error adding favourited song", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a ApiHandler) DeleteFavouritedSong(w http.ResponseWriter, r *http.Request) {
	id, err := verifyClientBaid(w, r)
	if err != nil {
		return
	}

	var deleteFavouritedSong model.FavouritedSong
	err = json.NewDecoder(r.Body).Decode(&deleteFavouritedSong)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// check that song ID links to a valid song
	_, exists := pkg.Datatable.Song[deleteFavouritedSong.SongId]
	if !exists {
		http.Error(w, "Invalid Song ID", http.StatusBadRequest)
		return
	}

	// check that user actually has the song favourited
	favouritedSongs, err := database.GetFavouritedSongs(id)
	if err != nil {
		http.Error(w, "Error checking favourited songs", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	songFavourited := false
	for _, songId := range favouritedSongs {
		if deleteFavouritedSong.SongId == songId { // song found in fav song list
			songFavourited = true
			break
		}
	}
	if !songFavourited { // song wasn't found in fav song list
		http.Error(w, "Invalid Song ID: Song already not favourited", http.StatusBadRequest)
		return
	}

	// adding new song to db
	err = database.DeleteFavouritedSong(id, deleteFavouritedSong.SongId)
	if err != nil {
		http.Error(w, "Error deleting favourited song", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
