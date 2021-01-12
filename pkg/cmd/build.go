package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/saromanov/starter/pkg/config"
	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/project"
)

// Build provides building of the tree structure for project
func Build(projectFlag string) error {
	pr := &models.Project{}
	pr.Type = models.StrToProjectType(projectFlag)
	err := consoleRead(pr)
	if err != nil {
		return errors.Wrap(err, "unable to read data from console")
	}
	if err := project.Build(pr); err != nil {
		return fmt.Errorf("unable to build project: %v", err)
	}

	return nil
}

// BuildFromConfig provides building of the project from config
func BuildFromConfig(configPath string) error {
	pr := config.DefaultConfig().ToModel()
	if configPath != "" {
		cfg, err := config.Load(configPath)
		if err != nil {
			return errors.Wrap(err, "unable to load config")
		}
		pr = cfg.ToModel()
	}
	if err := project.Build(pr); err != nil {
		return fmt.Errorf("unable to build project: %v", err)
	}
	return nil
}

func consoleRead(p *models.Project) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Add definition of the project(%v)\n", p.Type.String())
	fmt.Println("---------------------")
	name, err := readLine(reader, "Name of the project")
	if err != nil {
		return fmt.Errorf("unable to read user input: %v", err)
	}
	if err := validateName(name); err != nil {
		return fmt.Errorf("unable to validate name: %v", err)
	}
	p.Name = strings.ToLower(name[:len(name)-1])

	author, err := readLine(reader, "Author of the project")
	if err != nil {
		return fmt.Errorf("unable to read user input: %v", err)
	}
	if author == "" {
		return errNoAuthor
	}
	p.Author = author

	gitPath, err := readLine(reader, "Git path(optional)")
	if err != nil {
		return fmt.Errorf("unable to read git path: %v", err)
	}
	p.GitPath = gitPath

	entryFile, err := readLine(reader, "Entry file(optional)")
	if err != nil {
		return fmt.Errorf("unable to read git path: %v", err)
	}

	p.EntryFile = entryFile[:len(entryFile)-1]

	hostingDescription, err := readRepoConfig(reader)
	if err != nil {
		return fmt.Errorf("unable to get repo name")
	}
	if hostingDescription != "" {
		p.HostingDescription = hostingDescription[:len(hostingDescription)-1]
	}

	username, err := readLine(reader, "Enter Github username(optional)")
	if err != nil {
		return fmt.Errorf("unable to read badge line: %v", err)
	}
	p.Username = username[:len(username)-1]

	if p.Username != "" {
		badges, err := readLine(reader, "Entry badges to README. Separate it with commas. Supported: godoc,goreportcard")
		if err != nil {
			return fmt.Errorf("unable to read badge line: %v", err)
		}
		badgesRaw := badges[:len(badges)-1]
		p.Badges = strings.Split(badgesRaw, ",")
	}
	ciRaw, err := readLine(reader, "Enter CI providers(optional). Supported (github)")
	if err != nil {
		return fmt.Errorf("unable to read badge line: %v", err)
	}
	ciRaw = ciRaw[:len(ciRaw)-1]
	if len(ciRaw) > 0 {
		p.CI = toCIActions(strings.Split(ciRaw, ","))
	}
	return nil
}

// validateName provides validating of the project name
func validateName(name string) error {
	if name == "" {
		return errNoName
	}
	if len(name) < 3 || len(name) > 30 {
		return fmt.Errorf("unable to validate name of the project length")
	}
	return nil
}

func readLine(reader *bufio.Reader, name string) (string, error) {
	fmt.Println(name)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("unable to read line: %v", err)
	}
	return line, nil
}

func toCIActions(data []string) []models.CIActionType {
	result := make([]models.CIActionType, len(data))
	for i, d := range data {
		result[i] = models.ToCIActionType(d)
	}
	return result
}

// readRepoConfig provides configuration for creating repo name
func readRepoConfig(reader *bufio.Reader) (string, error) {
	data, err := readLine(reader, "Create repo at Github y/n")
	if err != nil {
		return "", fmt.Errorf("unable to read user input: %v", err)
	}
	if data != "y\n" {
		return "", nil
	}
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("GITHUB_TOKEN environment variable is not defined. ")
		token, err := readLine(reader, "GITHUB_TOKEN:")
		if err != nil {
			return "", fmt.Errorf("unable to get GITHUB_TOKEN")
		}
		os.Setenv("GITHUB_TOKEN", token[:len(token)-1])
		return "", nil
	}
	description, err := readLine(reader, "Enter description")
	if err != nil {
		return "", fmt.Errorf("unable to get github author")
	}
	return description, nil
}
