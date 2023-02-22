package main

import (
	"fmt"
	"net/http"
	"try_go_test/request_api"
)

func main() {
	var github = request_api.NewGithub("https://api.github.com", &http.Client{})
	var repoList, err = github.GetUserRepos("shushuTona")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, repo := range repoList {
		fmt.Println(repo.Name)
	}
}
