package gox

import (
	"strings"

	"dd"
)

// Clsx plugin
// Conditional classes rendering
type Clsx map[string]bool

func (c Clsx) Node() Node {
	return Class(c.join())
}

func (c Clsx) join() string {
	result := make([]string, 0)
	for classes, use := range c {
		if !use {
			continue
		}
		dd.Print(classes)
		result = append(result, classes)
	}
	return strings.Join(result, " ")
}
