package main

import (
	"fmt"
	"gist_file_sync/dir_scanner"
	"gopkg.in/yaml.v2"
	"log"
)

type T struct {
	syncfiles []string `yaml:"syncfiles"`
}

func loadYamlConfig() {
	t := T{}
	content := dir_scanner.ReadFileContent("./config.yaml")
	err := yaml.Unmarshal(content, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
}

func main() {
	loadYamlConfig()
	//githubToken := os.Getenv("GITHUB_GIST_TOKEN")
	//
	//ctx := context.Background()
	//ts := oauth2.StaticTokenSource(
	//	&oauth2.Token{AccessToken: githubToken},
	//)
	//tc := oauth2.NewClient(ctx, ts)
	//
	//client := github.NewClient(tc)
	//
	//// list all repositories for the authenticated user
	////gists, _, _ := client.Gists.List(ctx, "", nil)
	//
	//content := dir_scanner.ReadFileContent()
	//
	//createResult, _, err := client.Gists.Create(ctx, &github.Gist{
	//	Description: github.String("gist created by go"),
	//	Public:      github.Bool(false),
	//	Files: map[github.GistFilename]github.GistFile{
	//		github.GistFilename("main.go"): github.GistFile{
	//			Content: github.String(content),
	//		},
	//	},
	//})
	//if err != nil {
	//	fmt.Printf("Error creating gist: %v", err)
	//	return
	//}
	//
	//fmt.Println(createResult)
}
