package main

type stringFlag struct {
	set   bool
	value string
}

func (sf *stringFlag) Set(v string) error {
	sf.value = v
	sf.set = true
	return nil
}

func (sf *stringFlag) String() string {
	return sf.value
}
