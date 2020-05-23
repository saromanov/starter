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
	testhelper.RemoveContentFromRoot(dirName)
}

func getContentFromFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
