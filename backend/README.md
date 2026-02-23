

#  MultiDownloader Backend

A powerful **multi-platform video downloader** built with **Go** — supporting YouTube, TikTok, Instagram, Facebook, X (Twitter), and more.

No templates, no external APIs — just a clean **REST API** that lets you fetch and download videos directly using [`yt-dlp`](https://github.com/yt-dlp/yt-dlp).

---

## Features

- Download videos from **multiple platforms**
-  Unified API (`/download`, `/history`)
-  Keeps a local **download history**
-  Platform detection (YouTube, TikTok, Instagram, Facebook, Twitter/X)
-  Extensible structure for future extractors
-  No third-party API keys required

---

## Folder Structure

```

backend/
│
├── cmd/
│   └── server/main.go           # App entrypoint
│
├── internal/
│   ├── api/                     # REST routes and handlers
│   ├── downloader/              # Download and extraction logic
│   ├── storage/                 # History persistence
│   
│   
│
├── downloads/                   # Saved videos + history.json
│
├── go.mod
├── go.sum
└── README.md

````

---

## Requirements

- **Go 1.21+**
- **yt-dlp** installed on your system  
  (Required for multi-platform extraction)

### Install `yt-dlp`

```bash
# macOS or Linux
pip install -U yt-dlp

# or use Homebrew
brew install yt-dlp
````

---

## Setup

```bash
# Clone the project
git clone https://github.com/GordenArcher/multidownloader.git
cd multidownloader

# Initialize Go modules
go mod tidy

# Run the server
go run ./cmd/server
```

Server will start on
 **[http://localhost:8080](http://localhost:8080)**

---

## API Endpoints

### Download Video

**POST** `/download`

Request body:

```json
{
  "url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
}
```

Response:

```json
{
  "platform": "youtube",
  "title": "Unknown (metadata extraction later)",
  "url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
  "file_path": "downloads/1730400542.mp4",
  "downloaded_at": "2025-10-31T01:22:00Z"
}
```

---

### View Download History

**GET** `/history`

Response:

```json
[
  {
    "platform": "youtube",
    "title": "Unknown",
    "url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
    "file_path": "downloads/1730400542.mp4",
    "downloaded_at": "2025-10-31T01:22:00Z"
  }
]
```

---

## Platform Support

| Platform                        | Supported      | Notes                           |
| ------------------------------- | -------------- | ------------------------------- |
| YouTube                         |               | Full metadata & video download  |
| TikTok                          |               | Works with or without watermark |
| Instagram                       |               | Reels, posts, stories           |
| Facebook                        |               | Public videos only              |
| X (Twitter)                     |               | Videos & GIFs                   |
| Reddit, Vimeo, SoundCloud, etc. |   Coming soon |                                 |

---

##  Next Steps

* [ ] Add `/metadata` route using `yt-dlp --dump-json`
* [ ] Add concurrent downloads
* [ ] Add file cleanup logic
* [ ] Optional: integrate a small web UI (React or Svelte)

---

##  Example Usage (via curl)

```bash
curl -X POST http://localhost:8080/download \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ"}'
```

---

## License

MIT License © 2025 Gorden
