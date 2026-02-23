package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/GordenArcher/multidownloader/backend/internal/downloader"
	"github.com/GordenArcher/multidownloader/backend/internal/storage"
)

type DownloadRequest struct {
	URL string `json:"url"`
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	var req DownloadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	url := req.URL
	if url == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	var platform string
	switch {
	case strings.Contains(url, "youtube.com"), strings.Contains(url, "youtu.be"):
		platform = "youtube"
	case strings.Contains(url, "tiktok.com"):
		platform = "tiktok"
	case strings.Contains(url, "instagram.com"):
		platform = "instagram"
	case strings.Contains(url, "facebook.com"):
		platform = "facebook"
	case strings.Contains(url, "x.com"), strings.Contains(url, "twitter.com"):
		platform = "twitter"
	default:
		http.Error(w, "unsupported platform", http.StatusBadRequest)
		return
	}

	videoInfo, err := downloader.Download(platform, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := storage.SaveHistory(videoInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(videoInfo)
}
