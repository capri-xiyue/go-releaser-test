package main

import (
	"fmt"
	"github.com/abcxyz/pkg/buildinfo"
	"runtime/debug"
)

func main() {
	version := buildinfo.Version()
	fmt.Println(version)
	info, _ := debug.ReadBuildInfo()
	fmt.Println(info)
}
