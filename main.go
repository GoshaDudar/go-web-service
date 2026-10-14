package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type HealthResponse struct {
	Status string `json:"status"`
	AppMode string `json:"app_mode"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appMode := os.Getenv("APP_MODE")
	if appMode == "" {
		appMode = "development"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "Go-сервис версии 1.1 работает в Docker :)")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		
		response := HealthResponse {
			Status: "ok",
			AppMode: appMode,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("не удалось создать ответ: %v", err)
		}
	})

	address := ":" + port
	log.Printf("сервис запущен на порту %s", port)
	log.Printf("режим приложения: %s", appMode)

	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatal(err)
	}
}