package dirs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/saromanov/starter/pkg/models"
	"github.com/stretchr/testify/assert"
)

var rootTestDir = "./tmp"

func removeContentFromRoot(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func TestMakeRootDir(t *testing.T) {
	assert.NoError(t, makeRootDir(rootTestDir))
	assert.NoError(t, removeContentFromRoot(rootTestDir))
}

func TestDirs(t *testing.T) {
	p := New(&models.Project{
		Name:    rootTestDir,
		SubDirs: []string{"abc", "foobar"},
	})
	assert.NoError(t, p.Do())
	assert.NoError(t, New(&models.Project{Name: "acccc"}).Do())
	assert.NoError(t, removeContentFromRoot(rootTestDir))
	assert.Error(t, New(&models.Project{}).Do())
}
