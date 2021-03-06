package cmd

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"strings"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp source destination",
	Short: "Copy labels from [source] to [destination]",
	Run: func(cmd *cobra.Command, args []string) {
		c, ctx := getClient()
		if strings.HasPrefix(args[0], "github://") && strings.HasPrefix(args[1], "github://") {
			repoToRepo(ctx, c, args)
			return
		}
		if strings.HasPrefix(args[0], "github://") {
			repoToLocal(ctx, c, args)
			return
		}
		if strings.HasPrefix(args[1], "github://") {
			localToRepo(ctx, c, args)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(cpCmd)
}

func localToRepo(ctx context.Context, c *github.Client, args []string) {
	fmt.Println("Copying from local to remote")
	owner, repo := splitRepo(args[1])
	path := args[0]
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	labels := UnmarshalIntoLabels(file)
	for _, e := range labels {
		_, _, err := c.Issues.CreateLabel(ctx, owner, repo, &e)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func repoToRepo(ctx context.Context, c *github.Client, args []string) {
	fmt.Println("Copying from remote to remote")
	src_owner, src_repo := splitRepo(args[0])
	dest_owner, dest_repo := splitRepo(args[1])
	src_labels, _, err := c.Issues.ListLabels(ctx, src_owner, src_repo, nil)
	if err != nil {
		log.Fatal(err)
	}

	dest_labels, _, err := c.Issues.ListLabels(ctx, dest_owner, dest_repo, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range src_labels {
		_, _, err := c.Issues.CreateLabel(ctx, dest_owner, dest_repo, e)
		if err != nil {
			if containsLabel(dest_labels, "Name", *e.Name) {
				fmt.Printf("Skipping existing label: \"%s\": %s.\n", *e.Name, *e.Color)
			} else {
				log.Fatal(err)
			}
		}
	}
}

func repoToLocal(ctx context.Context, c *github.Client, args []string) {
	fmt.Println("Copying from remote to local")
	owner, repo := splitRepo(args[0])
	//fmt.Printf("Owner: %s, Repo: %s", owner, repo)
	labels, _, err := c.Issues.ListLabels(ctx, owner, repo, nil)
	if err != nil {
		panic(err)
	}

	labers := Labers(labels)
	y, err := labers.MarshalYAML()
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(y))
	err = ioutil.WriteFile(args[1], []byte(y), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func splitRepo(s string) (string, string) {
	s = strings.TrimPrefix(s, "github://")
	split := strings.Split(s, "/")
	return split[0], split[1]
}
