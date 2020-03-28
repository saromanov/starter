// Package github defines interactions with Github
package github

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/saromanov/starter/pkg/config"
	"golang.org/x/oauth2"
)

// Github defines representation with github
type Github struct {
	client *github.Client
	conf   *config.Config
}

// New provides initialization of the Github client
func New(c *config.Config) *Github {
	return &Github{
		conf:   c,
		client: makeClient(),
	}
}

// makeClient provides making of Github client
func makeClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

// CreateRepo provides creating of the repo
func (g *Github) CreateRepo(author, name string) error {
	_, _, err := g.client.Repositories.Create(context.Background(), author, &github.Repository{
		Name: &name,
	})
	if err != nil {
		return fmt.Errorf("unable to create repository")
	}
	return nil
}
