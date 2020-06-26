package verbump

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Version contains version information
type Version struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string
	Build      string
}

// ErrNegativeVersion negative version number error
var ErrNegativeVersion = fmt.Errorf("negative version number")

// ErrNotAlphanum non alphanumeric string error
var ErrNotAlphanum = fmt.Errorf("non alphanumeric [A-Z] [a-z] [0-9]")

// ErrTooManyLines too many lines in the string
var ErrTooManyLines = fmt.Errorf("too many new-lines")

var strCheckRegExp = regexp.MustCompile("^[A-Za-z0-9]+$")

// Make makes a version object from the version components and validates
// the input
func Make(major, minor, patch int, preRelease, build string) (Version, error) {
	v := Version{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		PreRelease: preRelease,
		Build:      build,
	}
	err := v.Validate()
	if err != nil {
		return v, err
	}
	return v, nil
}

// Parse parses a version string and extracts it's components
func Parse(version string) (Version, error) {
	v := Version{}
	if strings.Count(version, "\n") > 1 {
		return v, fmt.Errorf("invalid version string: %w", ErrTooManyLines)
	}
	version = strings.TrimSpace(version)
	buildSep := strings.Split(version, "+")
	buildSepLen := len(buildSep)
	if buildSepLen > 2 {
		return v, fmt.Errorf("invalid build string (%s)", buildSep[1:])
	}
	if buildSepLen == 2 {
		v.Build = buildSep[1]
	}

	preReleaseSep := strings.Split(buildSep[0], "-")
	preReleaseSepLen := len(preReleaseSep)
	if preReleaseSepLen > 2 {
		return v, fmt.Errorf("invalid pre release string (%s)", preReleaseSep[1:])
	}
	if preReleaseSepLen == 2 {
		v.PreRelease = preReleaseSep[1]
	}

	verSep := strings.Split(preReleaseSep[0], ".")
	verSepLen := len(verSep)
	if verSepLen > 3 {
		return v, fmt.Errorf("too many dot-separated version fields")
	}
	var err error
	if verSepLen >= 1 && verSep[0] != "" {
		v.Major, err = strconv.Atoi(verSep[0])
		if err != nil {
			return v, fmt.Errorf("error parsing major version number")
		}
	}
	if verSepLen >= 2 {
		v.Minor, err = strconv.Atoi(verSep[1])
		if err != nil {
			return v, fmt.Errorf("error parsing minor version number")
		}
	}
	if verSepLen >= 3 {
		v.Patch, err = strconv.Atoi(verSep[2])
		if err != nil {
			return v, fmt.Errorf("error parsing patch version number")
		}
	}
	err = v.Validate()
	if err != nil {
		return v, err
	}
	return v, nil
}

// Make makes a version string from it's components
func (v Version) String() string {
	ver := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.PreRelease != "" {
		ver += "-" + v.PreRelease
	}
	if v.Build != "" {
		ver += "+" + v.Build
	}
	return ver
}

// Validate validates the contents of Version
func (v Version) Validate() error {
	if v.Major < 0 {
		return fmt.Errorf("major version (%d): %w", v.Major, ErrNegativeVersion)
	}
	if v.Minor < 0 {
		return fmt.Errorf("minor version (%d): %w", v.Minor, ErrNegativeVersion)
	}
	if v.Patch < 0 {
		return fmt.Errorf("patch version (%d): %w", v.Patch, ErrNegativeVersion)
	}
	if v.PreRelease != "" && !strCheckRegExp.MatchString(v.PreRelease) {
		return fmt.Errorf("pre release string (%s): %w", v.PreRelease, ErrNotAlphanum)
	}
	if v.Build != "" && !strCheckRegExp.MatchString(v.Build) {
		return fmt.Errorf("build string (%s): %w", v.Build, ErrNotAlphanum)
	}
	return nil
}
