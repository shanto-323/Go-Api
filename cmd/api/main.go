package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
  "github.com/shanto-323/Go-Api/internal/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
  var r *chi.Mux = chi.NewRouter()
  handlers.Handler(r)

  fmt.Println("Starting Go Api Services...")

  err := http.ListenAndServe("localhost:8000",r)

  if err != nil {
    log.Error(err)
  }
}
