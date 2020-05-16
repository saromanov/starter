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
	data := map[string]func(*strings.Builder, *models.Project) error{
		"goreportcard": goreportcardBadge,
		"godoc":        godocBadge,
	}

	for _, b := range p.Badges {
		badge, ok := data[b]
		if !ok {
			continue
		}

		if err := badge(builder, p); err != nil {
			return fmt.Errorf("unable to generate badge: %v", err)
		}
	}
	builder.WriteString("\n")
	return nil
}

func godocBadge(builder *strings.Builder, p *models.Project) error {
	line := fmt.Sprintf("[![GoDoc](https://godoc.org/github.com/%s/%s?status.png)](https://godoc.org/github.com/%s/%s)", p.Username, p.Name)
	return writeBadge(builder, "godoc", line)
}

func goreportcardBadge(builder *strings.Builder, p *models.Project) error {
	line := fmt.Sprintf("[![Go Report Card](https://goreportcard.com/badge/github.com/%s/%s)](https://goreportcard.com/report/github.com/%s/%s)", p.Username, p.Name)
	return writeBadge(builder, "godoc", line)
}

func writeBadge(builder *strings.Builder, name, link string) error {
	if _, err := builder.WriteString(link); err != nil {
		return fmt.Errorf("unable to generate badges %s: %v", name, err)
	}

	return nil
}
