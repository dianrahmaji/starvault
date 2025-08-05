package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gorm.io/gorm"
)

type GraphQLQueryVariables struct {
	Page int `json:"page"`
}

type GraphQLRequest struct {
	Query     string                `json:"query"`
	Variables GraphQLQueryVariables `json:"variables"`
}

type GraphQLGetLiveStreamsData struct {
	GetLiveStreams []LiveStream `json:"getLivestreams"`
}

type GraphQLGetLiveStreamsResponse struct {
	Data GraphQLGetLiveStreamsData `json:"data"`
}

func createLiveStreamFolderName(liveStream LiveStream) (string, error) {
	baseDir := "./recordings/temp"

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", err
	}

	creator := sanitizeStr(liveStream.Creator.Name)
	title := sanitizeStr(liveStream.Title)
	timestamp := time.Now().Format("20060102-150405")

	finalDir := filepath.Join(baseDir, fmt.Sprintf("%s_%s_%s", creator, title, timestamp))

	if err := os.MkdirAll(finalDir, 0755); err != nil {
		return "", err
	}

	return finalDir, nil
}

func fetchLiveStreams(page int) ([]LiveStream, error) {
	url := "https://api.idn.app/graphql"
	query := `
		query ($page: Int) {
			getLivestreams(category: "all", page: $page) {
				title
				image_url
				playback_url
				status
				creator {
					name
					uuid
				}
			}
		}
	`
	variables := GraphQLQueryVariables{
		Page: page,
	}

	payload := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))

	if err != nil {
		return nil, err
	}

	var resp GraphQLGetLiveStreamsResponse

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return resp.Data.GetLiveStreams, nil
}

func getLiveStreams(db *gorm.DB, userMap map[string]User) ([]LiveStream, error) {
	_liveStreams := []LiveStream{}
	liveStreams := []LiveStream{}

	page := 1

	for {
		data, err := fetchLiveStreams(page)

		if err != nil {
			return nil, err
		}

		if len(data) == 0 {
			break
		}

		for _, stream := range data {
			_liveStreams = append(_liveStreams, stream)

			if stream.Status != StatusLive {
				continue
			}

			user, exists := userMap[stream.Creator.UUID]

			if !exists {
				continue
			}

			if user.Name == nil || *user.Name != stream.Creator.Name {
				db.Model(&user).Update("Name", stream.Creator.Name)
			}

			folder, err := createLiveStreamFolderName(stream)

			if err != nil {
				return nil, err
			}

			stream.Metadata.Folder = folder

			if !user.IsLive {
				db.Model(&user).Update("IsLive", true)
			}

			liveStreams = append(liveStreams, stream)
		}

		page++
	}

	// fmt.Println("all liveStreams:", _liveStreams)

	for _, liveStream := range _liveStreams {
		fmt.Println(liveStream.Creator.Name)
		fmt.Println(liveStream.Creator.UUID)
		fmt.Println(liveStream.ImageUrl)
	}

	return liveStreams, nil
}
