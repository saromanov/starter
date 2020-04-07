package dirs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var rootTestDir = "./tmp"

func TestMakeRootDir(t *testing.T) {
	assert.NoError(t, makeRootDir(rootTestDir))
}
