package readme

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/saromanov/starter/pkg/models"
	"github.com/saromanov/starter/pkg/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	dirName := "foobar"
	testhelper.MakeRootDir(dirName)
	p := New(nil)
	assert.Nil(t, p.Do())

	p = New(&models.Project{
		Name: dirName,
		Type: models.Library,
	})
	assert.NoError(t, p.Do())
	data, err := getContentFromFile(fmt.Sprintf("%s/README.md", dirName))
	assert.NoError(t, err)
	assert.Equal(t, "# foobar\n\n### Author\n\n### LICENCE\n", string(data))
	testhelper.RemoveContentFromRoot(dirName)
}

func TestDoBadges(t *testing.T) {
	dirName := "foobar"
	testhelper.MakeRootDir(dirName)
	p := New(&models.Project{
		Name:   dirName,
		Type:   models.Library,
		Badges: []string{"goreportcard"},
	})
	assert.NoError(t, p.Do())
	data, err := getContentFromFile(fmt.Sprintf("%s/README.md", dirName))
	assert.NoError(t, err)
	assert.Equal(t, "# foobar\n\n### Author\n\n### LICENCE\n", string(data))

	p = New(&models.Project{
		Name:     dirName,
		Type:     models.Library,
		Username: "foob",
		Badges:   []string{"godoc", "goreportcard"},
	})
	assert.NoError(t, p.Do())
	data, err = getContentFromFile(fmt.Sprintf("%s/README.md", dirName))
	assert.NoError(t, err)
	assert.Equal(t, "# foobar\n\n[![GoDoc](https://godoc.org/github.com/foob/foobar?status.png)](https://godoc.org/github.com/foob/foobar)[![Go Report Card](https://goreportcard.com/badge/github.com/foob/foobar)](https://goreportcard.com/report/github.com/foob/foobar)\n### Author\n\n### LICENCE\n", string(data))

	p = New(&models.Project{
		Name:     dirName,
		Type:     models.Library,
		Username: "foob",
		Badges:   []string{"godoc"},
	})
	assert.NoError(t, p.Do())
	data, err = getContentFromFile(fmt.Sprintf("%s/README.md", dirName))
	assert.NoError(t, err)
	assert.Equal(t, "# foobar\n\n[![GoDoc](https://godoc.org/github.com/foob/foobar?status.png)](https://godoc.org/github.com/foob/foobar)\n### Author\n\n### LICENCE\n", string(data))

	p = New(&models.Project{
		Name:     dirName,
		Type:     models.Library,
		Username: "foob",
		Badges:   []string{"anything"},
	})
	assert.NoError(t, p.Do())
	data, err = getContentFromFile(fmt.Sprintf("%s/README.md", dirName))
	assert.NoError(t, err)
	assert.Equal(t, "# foobar\n\n### Author\n\n### LICENCE\n", string(data))
	testhelper.RemoveContentFromRoot(dirName)
}

func getContentFromFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
