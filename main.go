package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	githubToken := os.Getenv("GITHUB_GIST_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	gists, _, _ := client.Gists.List(ctx, "", nil)

	createResult, r, err := client.Gists.Create(ctx, &github.Gist{
		Description: github.String("gist created by go"),
		Public:      github.Bool(false),
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename("main.go"): github.GistFile{
				Content: github.String("package main\n\nfunc main() {\n\tprintln(\"Hello world!\")\n}"),
			},
		},
	})
	if err != nil {
		return
	}

	fmt.Println(gists, createResult, r, err)
}
