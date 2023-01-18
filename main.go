package main

import (
	"fmt"
	"tekton-test/version"
)

func main() {
	fmt.Printf("version.Version: %v\n", version.Version)
	fmt.Printf("version.BuildTime: %v\n", version.BuildTime)
	fmt.Printf("version.GitCommit: %v\n", version.GitCommit)
}
