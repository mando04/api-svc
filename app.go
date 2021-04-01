package main

import (
	"encoding/json"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/status", func(rw http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		resp, _ := json.Marshal(map[string]string{"status": "ok", "hostname": hostname, "version": BuildVersion})
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(200)
		rw.Write(resp)
	})
}

func (a *App) Run(addr string) {
	log.Infoln("Listenin on http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
