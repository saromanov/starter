package entryfile

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
		Name:      "foobar",
		Type:      models.Library,
		EntryFile: "cmd/main.go",
	})
	assert.NoError(t, p.Do())
	removeContentFromRoot("foobar")
}

func TestGetFileName(t *testing.T) {

	assert.Equal(t, "main.go", getFileName("main.go"))
	assert.Equal(t, "main.go", getFileName("main"))
	assert.Equal(t, "/cmd/main.go", getFileName("/cmd/main"))
	assert.Equal(t, "/cmd/main.go", getFileName("/cmd/main.go"))
	assert.Equal(t, "./cmd/main.go", getFileName("./cmd/main"))
	assert.Equal(t, "./cmd/main.go", getFileName("./cmd/main.go"))
	assert.Equal(t, "./cmd/abs/main.go", getFileName("./cmd/abs/main"))
	assert.Equal(t, "./cmd/abs/main.go", getFileName("./cmd/abs/main.go"))
}

func TestGetPackageName(t *testing.T) {
	assert.Equal(t, "cmd", getPackageName("/cmd/main.go"))
	assert.Equal(t, "", getPackageName("main.go"))
	assert.Equal(t, "data", getPackageName("/cmd/data/main.go"))
	assert.Equal(t, "data", getPackageName("./cmd/data/main.go"))
}
