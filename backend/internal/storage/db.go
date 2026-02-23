package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/GordenArcher/multidownloader/backend/internal/downloader"
)

var historyFile = filepath.Join("downloads", "history.json")

func SaveHistory(video *downloader.VideoInfo) error {
	var history []*downloader.VideoInfo

	_ = os.MkdirAll("downloads", 0755)

	if data, err := os.ReadFile(historyFile); err == nil {
		_ = json.Unmarshal(data, &history)
	}

	history = append(history, video)

	data, _ := json.MarshalIndent(history, "", "  ")
	return os.WriteFile(historyFile, data, 0644)
}

func GetHistory() ([]*downloader.VideoInfo, error) {
	var history []*downloader.VideoInfo

	if data, err := os.ReadFile(historyFile); err == nil {
		_ = json.Unmarshal(data, &history)
	}
	return history, nil
}
