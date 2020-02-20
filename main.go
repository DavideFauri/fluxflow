package main

import (
	"github.com/DavideFauri/fluxflow/cmd"
	"github.com/DavideFauri/fluxflow/versioning"
)

// Deliberately uninitialized, see versioning package.
var version string

func main() {
	versioning.Set(version)

	cmd.Execute()
}
