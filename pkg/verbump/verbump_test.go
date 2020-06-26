package verbump

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerBump(t *testing.T) {
	ver, err := Bump("0.0.0", 1, 0, 0)
	assert.NoError(t, err)
	assert.Equal(t, "1.0.0", ver)

	ver, err = Bump("1.0.0", 0, 0, 0)
	assert.NoError(t, err)
	assert.Equal(t, "1.0.0", ver)

	ver, err = Bump("0.0.0", 0, 0, 0)
	assert.NoError(t, err)
	assert.Equal(t, "0.0.0", ver)

	ver, err = Bump("1.1.1", -1, -1, -1)
	assert.NoError(t, err)
	assert.Equal(t, "0.0.0", ver)

	ver, err = Bump("0.0.0", -1, -1, -1)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrNegativeVersion))

	ver, err = Bump("0.0.0", -1, 1, 1)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrNegativeVersion))

	ver, err = Bump("0.0.0", 1, -1, 1)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrNegativeVersion))

	ver, err = Bump("0.0.0", 1, 1, -1)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrNegativeVersion))

	ver, err = Bump("0.0.0-preRelease+build", 1, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, "1.1.1-preRelease+build", ver)

	ver, err = Bump("0.0.0+build", 1, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, "1.1.1+build", ver)

	ver, err = Bump("0.0.0-preRelease", 1, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, "1.1.1-preRelease", ver)
}
