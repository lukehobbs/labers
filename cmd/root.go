package cmd

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
	"net/url"
	"os"
)

type Labers []*github.Label

var RootCmd = &cobra.Command{
	Use:   "labers",
	Short: "A CLI for managing Github issue labels.",
	Long: `Labers is a CLI that allows you to save, copy, and transfer
Github issue labels between repositories using the Github API and YAML.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
}

func getClient() (*github.Client, context.Context) {
	dotEnv, err := godotenv.Read("labers.env")
	if err != nil {
		fmt.Println("Unable to locate environment file. You can configure environments by running \"labers init\".")
		os.Exit(1)
	}
	ctx := context.Background()
	tok := oauth2.Token{AccessToken: dotEnv["GITHUB_TOKEN"]}
	oauthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&tok))
	client := github.NewClient(oauthClient)
	if dotEnv["GITHUB_URL"] != "" {
		client.BaseURL = &url.URL{Scheme: "https", Host: dotEnv["GITHUB_URL"], Path: "api/v3"}
	}
	return client, ctx
}

func (l Labers) MarshalYAML() ([]byte, error) {
	m := make(map[string]interface{})
	for _, v := range l {
		m[*v.Name] = *v.Color
	}
	return yaml.Marshal(m)
}

func UnmarshalIntoLabels(file []byte) []github.Label {
	var m map[string]string
	err := yaml.Unmarshal(file, &m)
	if err != nil {
		panic(err)
	}
	var labels []github.Label
	for k := range m {
		name := k
		color := m[k]
		labels = append(labels, github.Label{Name: &name, Color: &color})
	}
	return labels
}

func containsLabel(labels []*github.Label, name string, value string) bool {
	for _, e := range labels {
		if *e.Name == value {
			return true
		}
	}
	return false
}
