package main

import (
	"fmt"
	"github.com/abcxyz/pkg/buildinfo"
)

func main() {
	version := buildinfo.Version()
	fmt.Println(version)
}
