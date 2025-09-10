package main

import (
	"fmt"
)

func main() {
	sv := NewSemanticVersion(1,2,3)
	sv.IncrementMajor()
	sv.IncrementMinor()
	sv.IncrementMinor()
	sv.IncrementPatch()
	sv.IncrementPatch()
	sv.IncrementMajor()
	fmt.Println(sv.String())
}