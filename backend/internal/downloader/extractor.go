package downloader

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

type VideoInfo struct {
	Platform     string    `json:"platform"`
	Title        string    `json:"title"`
	URL          string    `json:"url"`
	FilePath     string    `json:"file_path"`
	DownloadedAt time.Time `json:"downloaded_at"`
}

func Download(platform, url string) (*VideoInfo, error) {
	fileName := fmt.Sprintf("%d.mp4", time.Now().Unix())
	outputPath := filepath.Join("downloads", fileName)

	cmd := exec.Command("yt-dlp", "-o", outputPath, url)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("download failed: %v", err)
	}

	return &VideoInfo{
		Platform:     platform,
		Title:        "Unknown (metadata extraction later)",
		URL:          url,
		FilePath:     outputPath,
		DownloadedAt: time.Now(),
	}, nil
}
