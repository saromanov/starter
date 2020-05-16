package readme

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/saromanov/starter/pkg/models"
	"github.com/stretchr/testify/assert"
)

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
	return os.Remove(dir)
}

func TestDo(t *testing.T) {
	p := New(nil)
	assert.Nil(t, p.Do())

	p = New(&models.Project{
		Name: "foobar",
		Type: models.Library,
	})
	assert.NoError(t, p.Do())
	removeContentFromRoot("foobar")
}
