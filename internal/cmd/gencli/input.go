package main

import (
	"github.com/hashicorp/boundary/internal/types/resource"
)

var standardActions = map[string][]string{
	"read":   {"id"},
	"delete": {"id"},
	"list":   {"scope-id"},
}

type cmdInfo struct {
	// The type of the resource, e.g. "target"
	ResourceType string

	// The import path of the API package
	PkgPath string

	// Standard actions (with standard parameters) used by this resource
	StdActions []string

	// HasExtraCommandVars controls whether to generate an embedded struct with
	// extra command variables
	HasExtraCommandVars bool

	// HasExtraSynopsisFunc controls whether to generate code to look for an
	// extra synopsis function
	HasExtraSynopsisFunc bool

	// NonStdFlagsMap controls extra flags to add to the command
	ExtraFlagsMap map[string][]string
}

var inputStructs = map[string]*cmdInfo{
	"targets": {
		ResourceType:         resource.Target.String(),
		PkgPath:              "github.com/hashicorp/boundary/api/targets",
		StdActions:           []string{"read", "delete", "list"},
		HasExtraCommandVars:  true,
		HasExtraSynopsisFunc: true,
		ExtraFlagsMap: map[string][]string{
			"authorize-session": {"id", "host-id"},
			"add-host-sets":     {"id", "host-set", "version"},
			"remove-host-sets":  {"id", "host-set", "version"},
			"set-host-sets":     {"id", "host-set", "version"},
		},
	},
}
