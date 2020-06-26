# verbump
Bumps a semantic version number to a new number.

Useful for increasing the version number in build scripts (for example).

```
Options:
  -in string
        read version from file (-- for stdin))
  -major int
        major version increment
  -minor int
        minor version increment
  -out string
        write version to file (-- for stdout))
  -patch int
        patch version increment
  -print
        print resulting version when done
  -version string
        use version string
```

## Installation
From source (to the go bin folder):
```
git clone github.com/patrlind/verbump
cd verbump
go install
```

## Examples
Create a version file:  
`verbump -version "1.0.0" -out VERSION`

Bump patch version number up by one:  
`verbump -patch 1 -in VERSION -out VERSION`

Decrease the patch version number down by two: 
`verbump -patch -2 -in VERSION -out VERSION`

Pipe a version number thru the program:  
`echo "1.2.3" | verbump -major 1 -in -- -out --` (prints `2.2.3`)
