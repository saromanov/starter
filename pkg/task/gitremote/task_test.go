package gitremote

import (
	"testing"

	"github.com/saromanov/starter/pkg/models"
	"github.com/stretchr/testify/assert"
)

const remotePath = "https://github.com/saromanov/starter"

func TestDo(t *testing.T) {
	p := New(nil)
	assert.Nil(t, p.Do())

	p = New(&models.Project{
		Name: "foobar",
	})
	assert.Nil(t, p.Do())

	p = New(&models.Project{
		Name:    "foobar",
		GitPath: remotePath,
	})
	assert.NotNil(t, p.Do())
}
