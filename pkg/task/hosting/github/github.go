// Package github defines interactions with Github
package github

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/saromanov/starter/pkg/models"
	"golang.org/x/oauth2"
)

// Github defines representation with github
type Github struct {
	client *github.Client
	conf   *models.Project
}

// New provides initialization of the Github client
func New(c *models.Project) *Github {
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
func (g *Github) CreateRepo() error {
	_, _, err := g.client.Repositories.Create(context.Background(), "", &github.Repository{
		Name:        &g.conf.Name,
		Description: &g.conf.HostingDescription,
	})
	if err != nil {
		return fmt.Errorf("unable to create repository: %v", err)
	}
	return nil
}

// Do provides execution of the task
func (g *Github) Do() error {
	return g.CreateRepo()
}

func (g *Github) String() string {
	return "hosting-github"
}
