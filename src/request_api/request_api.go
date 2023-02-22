package request_api

import (
	"encoding/json"
	"errors"
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

type ErrorResponse struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}

type Github struct {
	Client *http.Client
	Domain string
}

func NewGithub(domain string, client *http.Client) *Github {
	return &Github{Client: client, Domain: domain}
}

func (github *Github) GetUserRepos(userName string) ([]GithubRepository, error) {
	var url = fmt.Sprintf("%s/users/%s/repos", github.Domain, userName)

	var req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var res, clientErr = github.Client.Do(req)
	if clientErr != nil {
		return nil, clientErr
	}
	defer res.Body.Close()

	var body, readErr = io.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	if res.StatusCode != 200 {
		var errRes ErrorResponse
		var unmarshalErr = json.Unmarshal(body, &errRes)
		if unmarshalErr != nil {
			return nil, unmarshalErr
		}

		return nil, errors.New(errRes.Message)
	}

	var grl []GithubRepository
	var unmarshalErr = json.Unmarshal(body, &grl)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return grl, nil
}
