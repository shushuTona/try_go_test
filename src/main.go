package main

import (
	"fmt"
	"try_go_test/request_api"
)

func main() {
	var repoList, err = request_api.GetGithubRepos("shushuTona")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, repo := range repoList {
		fmt.Println(repo.Name)
	}
}
