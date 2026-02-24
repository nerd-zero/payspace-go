package payspace

import (
	_ "embed"
	"strings"
)

//go:embed VERSION
var version string

// Version is the current version of the payspace-go SDK.
var Version = strings.TrimSpace(version)
