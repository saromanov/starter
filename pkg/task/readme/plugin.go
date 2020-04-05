package readme

import (
	"fmt"
	"strings"

	"github.com/saromanov/starter/pkg/models"
)

// Plugin defines method for additional data for README
type Plugin func(*strings.Builder, *models.Project) error

func description(b *strings.Builder, p *models.Project) error {
	desc := func(des string) string {
		if des == "" {
			return "\n"
		}
		return fmt.Sprintf("%s\n", des)
	}(p.Description)
	_, err := b.WriteString(desc)
	return err
}

func addBadges(builder *strings.Builder, p *models.Project) error {
	data := map[string]string{
		"goreportcard": "[![Go Report Card](https://goreportcard.com/badge/github.com/%s/%s)](https://goreportcard.com/report/github.com/%s/%s)",
		"godoc":        "[![GoDoc](https://godoc.org/github.com/%s/%s?status.png)](https://godoc.org/github.com/%s/%s)",
	}

	for _, b := range p.Badges {
		link, ok := data[b]
		if !ok {
			continue
		}
		if _, err := builder.WriteString(fmt.Sprintf(link, p.Username, p.Name, p.Username, p.Name)); err != nil {
			return fmt.Errorf("unable to generate badges: %v", err)
		}
	}
	builder.WriteString("\n")
	return nil
}
