package gitremote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	p := New(nil)
	assert.Nil(t, p.Do())
}
