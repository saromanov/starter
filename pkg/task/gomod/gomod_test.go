package gomod

import (
	"errors"
	"os"
	"testing"

	"github.com/saromanov/starter/pkg/models"
	"github.com/stretchr/testify/assert"
)

const tmpDir = "gomodtest"

func checkFile(path string) error {
	if _, err := os.Stat(path + "/" + "go.mod"); err != nil {
		return errors.New("go.mod file is not defined")
	}
	return nil
}

func TestDo(t *testing.T) {
	p := &models.Project{Name: tmpDir}
	d := New(p)
	assert.NoError(t, d.Do())
	assert.NoError(t, checkFile(tmpDir))
}
