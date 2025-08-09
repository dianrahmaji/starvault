package main

import (
	"log"
	"net/http"

	"gorm.io/gorm"
)

func startHttpServer(db *gorm.DB) {
	http.Handle("/monitor/", asynqMonitoringHandler())

	http.HandleFunc("/oauth/google", handleAuth)
	http.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
		handleOAuthCallback(w, r, db)
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
