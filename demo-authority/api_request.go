package demo_authority

import (
	"fmt"
	"net/url"
	"strconv"
)

type ApiRequest struct {
	baseURL   string
	token     string
	appId     string
	timestamp int64
}

func NewApiRequest(baseURL string, token string, appId string, timestamp int64) *ApiRequest {
	return &ApiRequest{
		baseURL:   baseURL,
		token:     token,
		appId:     appId,
		timestamp: timestamp,
	}
}

// BuildFromURL helper function to build ApiRequest from raw url.
func BuildFromURL(rawURL string) (*ApiRequest, error) {
	parse, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	query := parse.Query()

	parseInt, err := strconv.ParseInt(query["timestamp"][0], 10, 64)
	if err != nil {
		return nil, err
	}

	res := &ApiRequest{
		baseURL:   fmt.Sprintf("%s://%s", parse.Scheme, parse.Host),
		token:     query["token"][0],
		appId:     query["appId"][0],
		timestamp: parseInt,
	}

	return res, nil
}

func (r *ApiRequest) GetBaseUrl() string {
	return r.baseURL
}

func (r *ApiRequest) GetToken() string {
	return r.token
}

func (r *ApiRequest) GetAppId() string {
	return r.appId
}

func (r *ApiRequest) GetTimestamp() int64 {
	return r.timestamp
}
