package entryfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
