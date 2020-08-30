package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v32/github"
)

func Sum(x int, y int) int {
	// Return sum of x and y
	return x + y
}

func getActionInput(name string) string {
	return os.Getenv("INPUT_" + strings.ToUpper(name))
}

func main() {
	fmt.Println(Sum(5, 5))

	client := github.NewClient(nil)

	// list public repositories for org "github"
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)

	if err != nil {
		fmt.Println(err)
	}

	for _, repo := range repos {
		fmt.Println(repo.GetURL())
	}
}
