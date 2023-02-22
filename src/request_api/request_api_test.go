package request_api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewMockServer(f http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestGetUserRepos_OK(t *testing.T) {
	var testHandler = func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/users/shushuTona/repos" {
			t.Errorf("Faild request path, want %s, result %s", "/users/shushuTona/repos", req.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var repo = map[string]interface{}{
			"id":         461129021,
			"name":       "go_practice",
			"full_name":  "shushuTona/go_practice",
			"private":    false,
			"html_url":   "https://github.com/shushuTona",
			"created_at": "2022-02-19T08:17:40Z",
			"updated_at": "2022-02-19T17:16:14Z",
			"pushed_at":  "2022-03-03T18:51:24Z",
		}
		var body []map[string]interface{}
		body = append(body, repo)

		var b, _ = json.Marshal(body)
		w.Write(b)
	}

	var server = NewMockServer(testHandler)
	defer server.Close()

	var client = server.Client()
	var github = NewGithub(server.URL, client)
	var repoList, err = github.GetUserRepos("shushuTona")

	if err != nil {
		t.Errorf("Err %s", err.Error())
	}

	if repoList[0].Id != 461129021 {
		t.Errorf("Faild Id, want %d, result %d", 461129021, repoList[0].Id)
	}

	if repoList[0].Name != "go_practice" {
		t.Errorf("Faild Name, want %s, result %s", "go_practice", repoList[0].Name)
	}

	if repoList[0].FullName != "shushuTona/go_practice" {
		t.Errorf("Faild FullName, want %s, result %s", "shushuTona/go_practice", repoList[0].FullName)
	}

	if repoList[0].Private != false {
		t.Errorf("Faild Private, want %t, result %t", false, repoList[0].Private)
	}

	if repoList[0].HtmlUrl != "https://github.com/shushuTona" {
		t.Errorf("Faild HtmlUrl, want %s, result %s", "https://github.com/shushuTona", repoList[0].HtmlUrl)
	}

	if repoList[0].CreatedAt != "2022-02-19T08:17:40Z" {
		t.Errorf("Faild CreatedAt, want %s, result %s", "2022-02-19T08:17:40Z", repoList[0].CreatedAt)
	}

	if repoList[0].UpdatedAt != "2022-02-19T17:16:14Z" {
		t.Errorf("Faild UpdatedAt, want %s, result %s", "2022-02-19T17:16:14Z", repoList[0].UpdatedAt)
	}

	if repoList[0].PushedAt != "2022-03-03T18:51:24Z" {
		t.Errorf("Faild PushedAt, want %s, result %s", "2022-03-03T18:51:24Z", repoList[0].PushedAt)
	}
}

func TestGetUserRepos_Faild_NotFound(t *testing.T) {
	var testHandler = func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/users/not_exsist_user_name/repos" {
			t.Errorf("Faild request path, want %s, result %s", "/users/not_exsist_user_name/repos", req.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		var body = make(map[string]string)
		body["message"] = "Not Found"
		body["documentation_url"] = "https://docs.github.com/rest/reference/repos#list-repositories-for-a-user"

		var b, _ = json.Marshal(body)
		w.Write(b)
	}

	var server = NewMockServer(testHandler)
	defer server.Close()

	var client = server.Client()
	var github = NewGithub(server.URL, client)
	var repoList, err = github.GetUserRepos("not_exsist_user_name")

	if repoList != nil {
		t.Errorf("Faild : repoList is not nil, want %v, result %v", nil, repoList)
	}

	if err.Error() != "Not Found" {
		t.Errorf("Faild error message, want %s, result %s", "Not Found", err.Error())
	}
}
