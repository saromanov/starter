package makefile

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

func TestDo(t *testing.T) {
	p := New(&models.Project{
		Name:     rootTestDir,
		Makefile: "../../../assets/makefile-default",
	})
	assert.NoError(t, p.Do())

	assert.Nil(t, New(&models.Project{
		Name: rootTestDir,
	}).Do())

	assert.NoError(t, New(&models.Project{
		Name:     rootTestDir,
		Makefile: "../../../assets/makefile-default",
	}).Do())
}
