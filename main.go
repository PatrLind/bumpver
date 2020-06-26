package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/patrlind/verbump/pkg/verbump"
)

func main() {
	var version, fn, outFn string
	var major, minor, patch int
	var print bool
	flag.StringVar(&version, "version", "", "use version string (ignore -in)")
	flag.StringVar(&fn, "in", "", "read version from file (-- for stdin))")
	flag.StringVar(&outFn, "out", "", "write version to file (-- for stdout))")
	flag.IntVar(&major, "major", 0, "major version increment")
	flag.IntVar(&minor, "minor", 0, "minor version increment")
	flag.IntVar(&patch, "patch", 0, "patch version increment")
	flag.BoolVar(&print, "print", false, "print resulting version when done")
	flag.Parse()

	if fn == "" && version == "" {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		os.Exit(2)
	}

	if fn != "" && version != "" {
		fmt.Fprintln(os.Stderr, "Please specify only one of -version or -in")
		os.Exit(2)
	}

	if fn != "" {
		var err error
		version, err = readVersion(fn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading version: %v\n", err)
			os.Exit(1)
		}
	}

	ver, err := verbump.Bump(version, major, minor, patch)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error bumping version: %v\n", err)
		os.Exit(1)
	}

	if outFn != "" {
		err = writeVersion(outFn, ver)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing version: %v\n", err)
			os.Exit(1)
		}
	}

	if print {
		fmt.Fprintln(os.Stdout, ver)
	}
}

func readVersion(fn string) (string, error) {
	if fn == "--" {
		r := bufio.NewReader(os.Stdin)
		ver, err := r.ReadString('\n')
		if err != nil {
			return "", err
		}
		return ver, nil
	}
	f, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	ver, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	_, err = r.ReadString('\n')
	if err == nil {
		return "", fmt.Errorf("too many lines, only one line allowed in a version file")
	}
	return string(ver), nil
}

func writeVersion(fn, ver string) error {
	if fn == "--" {
		_, err := fmt.Fprintln(os.Stdout, ver)
		if err != nil {
			return fmt.Errorf("error writing to stdout: %w", err)
		}
	} else {
		f, err := os.Create(fn)
		if err != nil {
			return fmt.Errorf("error opening file for writing: %w", err)
		}
		defer f.Close()
		_, err = fmt.Fprintln(f, ver)
		if err != nil {
			return fmt.Errorf("error writing to '%s': %w", fn, err)
		}
		err = f.Close()
		if err != nil {
			return fmt.Errorf("error closing file: %w", err)
		}
	}
	return nil
}
