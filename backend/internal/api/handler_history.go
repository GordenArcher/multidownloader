package api

import (
	"encoding/json"
	"net/http"

	"github.com/GordenArcher/multidownloader/backend/internal/storage"
)

func handleHistory(w http.ResponseWriter, r *http.Request) {
	history, err := storage.GetHistory()
	if err != nil {
		http.Error(w, "could not load history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}
