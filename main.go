package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := initDB()

	if err != nil {
		panic(err)
	}

	c := startCronJob(db)

	go startAsynqWorker(db)
	go startHttpServer(db)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	c.Stop()

	resetIsRecordingStatus(db)
}
