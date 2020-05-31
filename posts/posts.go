package posts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Posts struct {
	endpoint   string
	httpClient *http.Client //wir definieren hier einen eigenen HttpClient
}

const (
	seaEndpoint    = "http://sa-bonn.ddnss.de:3000"
	defaultTimeout = 10 * time.Second
)

func New(endpoint string) *Posts {
	return &Posts{
		endpoint: endpoint,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func NewWithSEA() *Posts {
	return New(seaEndpoint)
}

func (p *Posts) loadPosts() ([]RemotePost, error) {
	var remotePosts []RemotePost
	var err error

	// http.Get("url") -> würde auch gehen, Quick&Dirty
	req, err := http.NewRequest(http.MethodGet, p.endpoint+"/posts", nil)
	if err != nil {
		return remotePosts, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("accept-encoding", "application/json")

	//http.DefaultClient.Do(req) -> besser selber definieren
	res, err := p.httpClient.Do(req)
	if err != nil {
		return remotePosts, fmt.Errorf("failed to execute request: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= http.StatusBadRequest {
		return remotePosts, fmt.Errorf("remote server returned status: %d", res.StatusCode)
	}

	respData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return remotePosts, fmt.Errorf("failed to load body: %w", err)
	}

	err = json.Unmarshal(respData, &remotePosts)
	if err != nil {
		return remotePosts, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return remotePosts, nil

}
