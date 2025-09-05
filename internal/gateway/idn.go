package gateway

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/dianrahmaji/star-vault/internal/port"
)

const query = `
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

type IDNGateway struct {
	apiURL string
}

type GraphQLQueryVariables struct {
	Page int `json:"page"`
}

type GraphQLRequest struct {
	Query     string                `json:"query"`
	Variables GraphQLQueryVariables `json:"variables"`
}

type GraphQLGetLiveStreamsData struct {
	Livestreams []port.LivestreamDTO `json:"getLivestreams"`
}

type GraphQLGetLiveStreamsResponse struct {
	Data GraphQLGetLiveStreamsData `json:"data"`
}

func NewIDNGateway(apiURL string) *IDNGateway {
	return &IDNGateway{apiURL: apiURL}
}

func (g *IDNGateway) FetchLivestreams() ([]port.LivestreamDTO, error) {
	page := 1
	livestreams := []port.LivestreamDTO{}

	for {
		variables := GraphQLQueryVariables{Page: page}

		payload := GraphQLRequest{
			Query:     query,
			Variables: variables,
		}

		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		res, err := http.Post(g.apiURL, "application/json", bytes.NewBuffer(jsonPayload))
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		var resp GraphQLGetLiveStreamsResponse

		if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
			return nil, err
		}

		if len(resp.Data.Livestreams) == 0 {
			break
		}

		for i := range resp.Data.Livestreams {
			livestreams = append(livestreams, resp.Data.Livestreams[i])
		}

		page++
	}

	return livestreams, nil
}
