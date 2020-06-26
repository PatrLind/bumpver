package verbump

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	v, err := Make(1, 1, 1, "", "")
	assert.NoError(t, err)
	assert.Equal(t, 1, v.Major)
	assert.Equal(t, 1, v.Minor)
	assert.Equal(t, 1, v.Patch)
	assert.Equal(t, "", v.PreRelease)
	assert.Equal(t, "", v.Build)
	assert.Equal(t, "1.1.1", v.String())

	_, err = Make(1, 1, 1, "-test", "")
	assert.Error(t, err)
	_, err = Make(1, 1, 1, "", "+test")
	assert.Error(t, err)
	_, err = Make(1, 1, 1, "!!", "!!")
	assert.Error(t, err)

	v, err = Parse("1")
	assert.NoError(t, err)
	assert.Equal(t, 1, v.Major)
	assert.Equal(t, 0, v.Minor)
	assert.Equal(t, 0, v.Patch)
	assert.Equal(t, "", v.PreRelease)
	assert.Equal(t, "", v.Build)
	assert.Equal(t, "1.0.0", v.String())

	v, err = Parse("1\n")
	assert.NoError(t, err)

	v, err = Parse("")
	assert.NoError(t, err)

	v, err = Parse("1\n\n")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrTooManyLines))

	v, err = Parse("10.11")
	assert.NoError(t, err)
	assert.Equal(t, 10, v.Major)
	assert.Equal(t, 11, v.Minor)
	assert.Equal(t, 0, v.Patch)
	assert.Equal(t, "", v.PreRelease)
	assert.Equal(t, "", v.Build)
	assert.Equal(t, "10.11.0", v.String())

	v, err = Parse("-10.11")
	assert.Error(t, err)

	v, err = Parse("0.0.0.0")
	assert.Error(t, err)
	v, err = Parse("0.0.0-preRelease.1")
	assert.Error(t, err)
	v, err = Parse("0.0.0-preRelease+build.1")
	assert.Error(t, err)

	v, err = Parse("0.0.0-preRelease+build")
	assert.NoError(t, err)
	assert.Equal(t, "preRelease", v.PreRelease)
	assert.Equal(t, "build", v.Build)
	assert.Equal(t, "0.0.0-preRelease+build", v.String())

	v, err = Parse("0.0.0-preRelease")
	assert.NoError(t, err)
	assert.Equal(t, "preRelease", v.PreRelease)
	assert.Equal(t, "", v.Build)
	assert.Equal(t, "0.0.0-preRelease", v.String())

	v, err = Parse("0.0.0+build")
	assert.NoError(t, err)
	assert.Equal(t, "", v.PreRelease)
	assert.Equal(t, "build", v.Build)
	assert.Equal(t, "0.0.0+build", v.String())
}
