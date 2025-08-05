package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func spawnRecordings(db *gorm.DB, userMap map[string]User, liveStreams []LiveStream) {
	for _, liveStream := range liveStreams {
		user, exists := userMap[liveStream.Creator.UUID]

		if !exists || user.IsRecording {
			continue
		}

		if err := db.Model(&user).Update("IsRecording", true).Error; err != nil {
			log.Println("Failed to set IsRecording = true for", user.ID, err)
			continue
		}

		downloadThumbnail(liveStream)
		go startRecording(db, liveStream)
	}
}

func downloadThumbnail(liveStream LiveStream) {
	resp, err := http.Get(liveStream.ImageUrl)

	if err != nil {
		fmt.Println("Failed to download thumbnail")
	}

	defer resp.Body.Close()

	outputPath := filepath.Join(liveStream.Metadata.Folder, "thumbnail.jpg")
	fmt.Println(outputPath)

	outputFile, err := os.Create(outputPath)

	if err != nil {
		fmt.Println("Failed to create output file")
	}

	defer outputFile.Close()

	_, err = io.Copy(outputFile, resp.Body)

	if err != nil {
		fmt.Println("Failed to save the thumbnail")
	}
}

func recordLiveStreams(db *gorm.DB) {
	userMap, err := createUserMap(db)

	if err != nil {
		return
	}

	liveStreams, err := getLiveStreams(db, userMap)

	if err != nil {
		return
	}

	updateLiveStreamStatus(db, userMap, liveStreams)

	spawnRecordings(db, userMap, liveStreams)
}

func startCronJob(db *gorm.DB) *cron.Cron {
	c := cron.New()

	recordLiveStreams(db)
	c.AddFunc("@every 1m", func() {
		recordLiveStreams(db)
	})

	c.Start()

	return c
}
