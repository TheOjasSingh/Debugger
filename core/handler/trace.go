package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"debugr/internal/trace"
)

func TraceHandler(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil || pid <= 0 {
		http.Error(w, "Invalid PID", http.StatusBadRequest)
		return
	}

	stack := trace.CollectStack(pid)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"stack": stack,
	})
}
