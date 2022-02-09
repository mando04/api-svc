package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("fatal error %s", err)))
	}
	var resp = map[string]string{
		"status":   "ok",
		"date":     time.Now().String(),
		"version":  BuildVersion,
		"env":      os.Getenv("ENV"),
		"hostname": hostname,
	}
	marshaledResp, _ := json.Marshal(resp)
	w.Write([]byte(marshaledResp))
}
