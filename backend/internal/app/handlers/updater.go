package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/updater"
	"net/http"
	"strconv"
	"strings"
)

type UpdaterHandler struct{}

func (a UpdaterHandler) GetUpdaterVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(updater.UpdaterVars.Version))
}
func (a UpdaterHandler) GetQuickDownload(w http.ResponseWriter, r *http.Request) {
	url := updater.UpdaterVars.ManualQuickDownloadURI
	http.Redirect(w, r, url, http.StatusFound)
}

func (a UpdaterHandler) GetFullDownload(w http.ResponseWriter, r *http.Request) {
	url := updater.UpdaterVars.ManualFullDownloadURI
	http.Redirect(w, r, url, http.StatusFound)
}

func (a UpdaterHandler) Releases(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	release, ok := updater.UpdaterVars.Releases[name]
	if !ok {
		http.NotFound(w, r)
		return
	}

	var req struct {
		Version  string `json:"version"`
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
	if req.Version == fmt.Sprintf("%s$%s", name, release.Version) {
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
	ret.Version = fmt.Sprintf("%s$%s", name, release.Version)
	//	ret.URI = release.URI
	ret.Name = release.Name
	parts := strings.Split(req.Version, "$") // "latest$1.0" -> ["latest", "1.0"]
	if len(parts) != 2 || parts[0] != release.Name {
		ret.URI = release.FullDownloadURI
		ret.DeleteCabinet = true
		json.NewEncoder(w).Encode(ret)
		return
	}
	currVer, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		http.Error(w, "Invalid version number", http.StatusBadRequest)
		return
	}
	minimumQuickDownloadVer, err := strconv.ParseFloat(release.MinimumQuickDownloadVer, 64)
	if err != nil {
		http.Error(w, "Minimum Quick Download Version Error (Contact EGTS Admins)", http.StatusBadRequest)
		return
	}
	cabinetVer, err := strconv.ParseFloat(release.Cabinet, 64)
	if err != nil {
		http.Error(w, "Cabinet Version Error (Contact EGTS Admins)", http.StatusBadRequest)
		return
	}
	if currVer < minimumQuickDownloadVer {
		ret.URI = release.FullDownloadURI
	} else {
		ret.URI = release.QuickDownloadURI
	}
	if currVer < cabinetVer {
		ret.DeleteCabinet = true
	} else {
		ret.DeleteCabinet = false
	}
	json.NewEncoder(w).Encode(ret)
}
