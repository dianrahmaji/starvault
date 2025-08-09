package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
)

func convertTsToMp4(sourcePath string) (string, error) {
	resultPath := strings.TrimSuffix(sourcePath, filepath.Ext(sourcePath)) + ".mp4"

	cmd := exec.Command("ffmpeg",
		"-i", sourcePath,
		"-c", "copy",
		"-bsf:a", "aac_adtstoasc",
		"-y", resultPath,
	)

	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to convert TS to MP4: %w", err)
	}

	err = os.Remove(sourcePath)

	if err != nil {
		fmt.Println("Error removing file:", err)
	}
	fmt.Println("File removed successfully.")

	return resultPath, nil
}

func startRecording(db *gorm.DB, liveStream LiveStream) {
	defer func() {
		fmt.Println("Recording ended for:", liveStream.Creator.Name, liveStream.Creator.UUID)
		resetLiveStreamStatus(db, liveStream)
	}()

	outputPath := filepath.Join(liveStream.Metadata.Folder, "result.ts")

	cmd := exec.Command("ffmpeg",
		"-reconnect", "1",
		"-reconnect_streamed", "1",
		"-reconnect_delay_max", "2",
		"-i", liveStream.PlaybackUrl,
		// "-t", "5",
		"-c", "copy",
		"-f", "mpegts",
		"-y", outputPath,
	)

	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Run()

	if err != nil {
		fmt.Println("Error running ffmpeg:", err)
	}

	_, err = convertTsToMp4(outputPath)

	if err != nil {
		fmt.Println("Failed to convert TS to MP4:", err)
		return
	}

	if err := enqueueUploadJob(liveStream); err != nil {
		fmt.Println("Failed to enqueue upload task:", err)
	}
}
