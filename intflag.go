package main

import "strconv"

type intFlag struct {
	set   bool
	value int
}

func (f *intFlag) Set(v string) error {
	f.set = true
	i, err := strconv.ParseInt(v, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	f.value = int(i)
	return err
}

func (f *intFlag) String() string {
	return strconv.Itoa(f.value)
}
