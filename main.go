package main

import (
	"github.com/hxrgjo/gitlab-ci/cmd"
)

var (
	VERSION string
	COMMIT  string
)

func main() {

	// get build param
	cmd.VERSION = VERSION
	cmd.COMMIT = COMMIT

	cmd.Execute()
}

func Sum(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
