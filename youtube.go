package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/hibiken/asynq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)


func getClient(ctx context.Context) *http.Client {
	clientID := os.Getenv("YOUTUBE_CLIENT_ID")
	clientSecret := os.Getenv("YOUTUBE_CLIENT_SECRET")
	refreshToken := os.Getenv("YOUTUBE_REFRESH_TOKEN")
	accessToken := os.Getenv("YOUTUBE_ACCESS_TOKEN")

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{youtube.YoutubeUploadScope},
	}

	token := &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return conf.Client(ctx, token)
}

type ProgressReader struct {
	file       *os.File
	read       int64
	totalSize  int64
	lastReport int
}

func (p *ProgressReader) Read(buf []byte) (int, error) {
	n, err := p.file.Read(buf)
	if n > 0 {
		p.read += int64(n)
		percent := int(float64(p.read) / float64(p.totalSize) * 100)
		fmt.Printf("Uploading... %d%%\n", percent)
		p.lastReport = percent
	}
	return n, err
}

func removeFiles(path string) {
	err := os.RemoveAll(path)

	if err != nil {
		fmt.Println("Failed to remove folder")
	}
}

func createYouTubeService(ctx context.Context) (*youtube.Service, error) {
	client := getClient(ctx)

	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))

	if err != nil {
		return nil, err
	}

	return service, nil
}

func createLiveStreamTitle(liveStream LiveStream) string {
	creator := liveStream.Creator.Name
	title := liveStream.Title
	timestamp := time.Now().Format("02 Jan 2006 15:04")

	return fmt.Sprintf("%s - %s | %s", creator, title, timestamp)
}

func uploadVideo(service *youtube.Service, liveStream LiveStream) (string, error) {
	filePath := filepath.Join(liveStream.Metadata.Folder, "result.mp4")
	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return "", err

	}

	fmt.Printf("uploading %s to youtube\n", liveStream)

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:      createLiveStreamTitle(liveStream),
			CategoryId: "24",
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: "private",
		},
	}

	call := service.Videos.Insert([]string{"snippet", "status"}, upload)

	progress := &ProgressReader{
		file:       file,
		totalSize:  fi.Size(),
		lastReport: -1,
	}
	response, err := call.Media(progress).Do()

	if err != nil {
		return "", err
	}

	return response.Id, nil

}

func setThumbnail(service *youtube.Service, id string, liveStream LiveStream) error {
	imagePath := filepath.Join(liveStream.Metadata.Folder, "thumbnail.jpg")
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = service.Thumbnails.Set(id).Media(file).Do()

	return err
}

func createPlaylist(service *youtube.Service, title string) (string, error) {
	resp, err := service.Playlists.List([]string{"snippet"}).Mine(true).Do()

	if err != nil {
		return "", err
	}

	for _, item := range resp.Items {
		if item.Snippet.Title == title {
			return item.Id, nil
		}
	}

	pl := &youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{Title: title},
		Status:  &youtube.PlaylistStatus{PrivacyStatus: "private"},
	}

	newPl, err := service.Playlists.Insert([]string{"snippet", "status"}, pl).Do()
	if err != nil {
		return "", err
	}

	return newPl.Id, nil
}

func addVideoToPlaylist(service *youtube.Service, videoId string, playlistId string) error {
	item := &youtube.PlaylistItem{
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: playlistId,
			ResourceId: &youtube.ResourceId{
				Kind:    "youtube#video",
				VideoId: videoId,
			},
		},
	}

	_, err := service.PlaylistItems.Insert([]string{"snippet"}, item).Do()

	return err
}

func handleUploadToYouTube(ctx context.Context, t *asynq.Task) error {
	var liveStream LiveStream
	if err := json.Unmarshal(t.Payload(), &liveStream); err != nil {
		return err
	}

	service, err := createYouTubeService(ctx)

	videoId, err := uploadVideo(service, liveStream)

	if err != nil {
		return fmt.Errorf("error creating YouTube client: %w", err)
	}

	err = setThumbnail(service, videoId, liveStream)

	playlistId, err := createPlaylist(service, liveStream.Creator.Name)

	if err != nil {
		return fmt.Errorf("error creating playlist: %w", err)
	}

	err = addVideoToPlaylist(service, videoId, playlistId)

	if err != nil {
		return fmt.Errorf("error adding video to playlist: %w", err)
	}

	removeFiles(liveStream.Metadata.Folder)

	fmt.Printf("Video uploaded successfully!")

	return nil
}
