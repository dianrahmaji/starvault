package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

func enqueueUploadJob(liveStream LiveStream) error {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	defer client.Close()

	payload, err := json.Marshal(liveStream)
	fmt.Println("payload", payload)

	if err != nil {
		return err
	}

	task := asynq.NewTask("video:upload", payload)

	_, err = client.Enqueue(
		task,
		asynq.TaskID(liveStream.Metadata.Folder),
		asynq.Queue("default"),
		asynq.Retention(1*time.Minute),
	)

	if err != nil {
		return fmt.Errorf("failed to enqueue upload task: %w", err)
	}

	return nil
}

func startAsynqWorker() {
	redisOpt := asynq.RedisClientOpt{Addr: "localhost:6379"}
	srv := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 5,
		Queues: map[string]int{
			"default": 1,
		},
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc("video:upload", handleUploadToYouTube)

	if err := srv.Run(mux); err != nil {
		log.Fatal("Asynq worker error:", err)
	}
}

func startAsynqMonitoring() {
	mon := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitor",
		RedisConnOpt: asynq.RedisClientOpt{Addr: "localhost:6379"},
	})

	http.Handle("/monitor/", mon)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
