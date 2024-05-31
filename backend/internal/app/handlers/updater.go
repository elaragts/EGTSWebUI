package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/updater"
	"net/http"
	"strconv"
)

type UpdaterHandler struct{}

func (a UpdaterHandler) GetUpdaterVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(updater.UpdaterVars.Version))
}

func (a UpdaterHandler) Releases(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	release, ok := updater.UpdaterVars.Releases[name]
	if !ok {
		http.NotFound(w, r)
		return
	}

	var req struct {
		Verison  string `json:"version"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Password != release.Password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if req.Verison == release.Version {
		http.Error(w, "Already up to date", http.StatusNotModified)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var ret struct {
		Version       string `json:"version"`
		URI           string `json:"uri"`
		Name          string `json:"name"`
		DeleteCabinet bool   `json:"deleteCabinet"`
	}
	ret.Version = release.Version
	ret.URI = release.URI
	ret.Name = release.Name
	currVer, err := strconv.ParseFloat(req.Verison, 64)
	if err != nil {
		http.Error(w, "Invalid version number", http.StatusBadRequest)
		return
	}
	cabinetVer, err := strconv.ParseFloat(release.Cabinet, 64)
	if err != nil {
		http.Error(w, "Cabinet Version Error (Contact TPS Admins)", http.StatusBadRequest)
		return
	}
	if currVer < cabinetVer {
		ret.DeleteCabinet = true
	} else {
		ret.DeleteCabinet = false
	}
	json.NewEncoder(w).Encode(ret)
}
