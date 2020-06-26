package verbump

import (
	"fmt"
)

// Bump bumps one or more of the version numbers with the amount indicated
// in the respective number fields. Negative numbers are allowed for version
// decrease as long as the version doesn't go below 0.
// The resulting version string will be returned
func Bump(version string, major, minor, patch int) (string, error) {
	v, err := Parse(version)
	if err != nil {
		return "", fmt.Errorf("error parsing version (%s): %w", version, err)
	}
	v.Major += major
	v.Minor += minor
	v.Patch += patch
	err = v.Validate()
	if err != nil {
		return "", fmt.Errorf("failed to bump version: %w", err)
	}

	return v.String(), nil
}
