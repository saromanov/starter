package dirs

import (
	"os"
	"path/filepath"
	"testing"

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
