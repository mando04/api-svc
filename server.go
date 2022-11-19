package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func Run(router chi.Router){
	log.Infoln("starting listener on :8443...")
	log.Fatal(http.ListenAndServeTLS(":8443", "ssl/localhost.crt", "ssl/localhost.key", router))
}
