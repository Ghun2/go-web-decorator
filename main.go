package main

import (
	"github.com/Ghun2/go-web-decorator/decoHandler"
	"github.com/Ghun2/go-web-decorator/myapp"
	"log"
	"net/http"
	"time"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[Logger1] Started")
	h.ServeHTTP(w, r)
	log.Println("[Logger1] Completed time: ", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[Logger2] Started")
	h.ServeHTTP(w, r)
	log.Println("[Logger2] Completed time: ", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	h := myapp.NewHandler()
	h = decoHandler.NewDecoHandler(h, logger)
	h = decoHandler.NewDecoHandler(h, logger2)
	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}



