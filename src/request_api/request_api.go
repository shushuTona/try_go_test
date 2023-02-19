package request_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GithubRepository struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	FullName  string `json:"full_name"`
	Private   bool   `json:"private"`
	HtmlUrl   string `json:"html_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	PushedAt  string `json:"pushed_at"`
}

func GetGithubRepos(userName string) ([]GithubRepository, error) {
	var url = fmt.Sprintf("https://api.github.com/users/%s/repos", userName)

	var req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var client = &http.Client{}
	var res, clientErr = client.Do(req)
	if clientErr != nil {
		return nil, clientErr
	}
	defer res.Body.Close()

	var body, readErr = io.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	var grl []GithubRepository
	var unmarshalErr = json.Unmarshal(body, &grl)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return grl, nil
}
