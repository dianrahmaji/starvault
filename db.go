package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createUserMap(db *gorm.DB) (map[string]User, error) {
	var users []User
	if err := db.Select("ID", "Name", "IsLive", "IsRecording").Debug().Find(&users).Error; err != nil {
		log.Println("[createUserMap]: failed to fetch users", err)
		return nil, err
	}

	userMap := make(map[string]User, len(users))

	for _, user := range users {
		userMap[user.ID] = user
	}

	return userMap, nil
}

func updateLiveStreamStatus(db *gorm.DB, userMap map[string]User, liveStreams []LiveStream) {
	liveStreamsUUIDSet := make(map[string]bool, len(liveStreams))

	for _, stream := range liveStreams {
		liveStreamsUUIDSet[stream.Creator.UUID] = true
	}

	for UUID, user := range userMap {
		if (user.IsLive || user.IsRecording) && !liveStreamsUUIDSet[UUID] {
			resetLiveStreamStatus(db, LiveStream{Creator: LiveStreamCreator{UUID: UUID}})
		}
	}
}

func resetLiveStreamStatus(db *gorm.DB, liveStream LiveStream) {
	err := db.Model(&User{ID: liveStream.Creator.UUID}).
		Select("IsLive", "IsRecording").
		Updates(User{
			IsLive:      false,
			IsRecording: false,
		}).Error

	if err != nil {
		log.Println("Failed to reset status for", liveStream.Creator.UUID, err)
	}
}

func resetIsRecordingStatus(db *gorm.DB) {
	if err := db.Model(&User{}).Where("is_recording = ?", true).Update("IsRecording", false).Error; err != nil {
		println("Failed to reset is_recording flags:", err.Error())
	}
}

func initDB() (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal("failed to migrate schema:", err)
	}

	users := []User{
		// {ID: "61d9d168-a875-498d-97e3-7126f2c7c208"},
		// {ID: "331b1f61-7819-4881-b6c5-a8554670ea87"},
		// {ID: "e6168cbe-c838-4b64-93bd-35371cd2441c"},
		// {ID: "741260d4-7fb5-4a84-aee7-109f6878f1cc"},
		// {ID: "396ea983-9aad-4460-8641-e67f2c871030"},
		// {ID: "396ea983-9aad-4460-8641-e67f2c871030"},
		// {ID: "cbc0bbe1-d6ea-4646-bfcf-24b3a642fa21"},
	}

	for _, u := range users {
		db.FirstOrCreate(&u, User{ID: u.ID})

	}

	return db, err
}
