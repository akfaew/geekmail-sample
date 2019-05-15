package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	geekmail "github.com/geekmail/go-geekmail"
)

func main() {
	apiToken := os.Getenv("GEEKMAIL_APITOKEN")
	if apiToken == "" {
		fmt.Fprintf(os.Stderr, "GEEKMAIL_APITOKEN must be set before running this example\n")
		os.Exit(1)
	}
	CreateDraftFromGitHub(apiToken)
	CreateDraftFromFile(apiToken)
}

func CreateDraftFromGitHub(apiToken string) {
	conf := &geekmail.Conf{
		GitHubAuth: geekmail.GitHubAuth{
			Repository: "github.com/geekmail/geekmail-sample",
			Secret:     "password123", // Secret from geekmail.yaml found inside the repository
		},
		APIToken: apiToken,
	}

	vars := map[string]string{
		"To":   "John Doe <john@example.com>",
		"Name": "John",
	}

	draft := &geekmail.DraftCreate{
		TemplatePath: "templates/example.template", // Path to the template in the GitHub repository
		Vars:         vars,
		Labels:       []string{"GeekMail"},
	}

	client := geekmail.NewClient(&http.Client{}, conf)
	if resp, err := client.Draft.Create(context.Background(), draft); err != nil {
		fmt.Printf("Cannot communicate with the API: err=%v\n", err)
	} else {
		fmt.Printf("GeekMail response: %+v\n", resp)
	}
}

func CreateDraftFromFile(apiToken string) {
	conf := &geekmail.Conf{
		APIToken: apiToken,
	}

	vars := map[string]string{
		"To":   "John Doe <john@example.com>",
		"Name": "John",
	}

	// Read in a template from file
	contents, err := ioutil.ReadFile("../../templates/example.template")
	if err != nil {
		panic(err)
	}

	draft := &geekmail.DraftCreate{
		Template: string(contents), // template contents
		Vars:     vars,
		Labels:   []string{"GeekMail"},
	}

	client := geekmail.NewClient(&http.Client{}, conf)
	if resp, err := client.Draft.Create(context.Background(), draft); err != nil {
		fmt.Printf("Cannot communicate with the API: err=%v\n", err)
	} else {
		fmt.Printf("GeekMail response: %+v\n", resp)
	}
}
