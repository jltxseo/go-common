package config

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kelseyhightower/envconfig"
)

// PersisterType is the type of persister to use.
type PersisterType int

const (
	// PersisterTypeInvalid is an invalid persister value
	PersisterTypeInvalid PersisterType = iota

	// PersisterTypeNone is a persister that does nothing but return default values
	PersisterTypeNone

	// PersisterTypePostgresql is a persister that uses PostgreSQL as the backend
	PersisterTypePostgresql
)

var (
	// PersisterNameToType maps valid persister names to the types above
	PersisterNameToType = map[string]PersisterType{
		"none":       PersisterTypeNone,
		"postgresql": PersisterTypePostgresql,
	}
)

const (
	usageFormat = `The %v is configured via environment vars only. The following environment variables can be used:
{{range .}}
{{usage_key .}}
  description: {{usage_description .}}
  type:        {{usage_type .}}
  default:     {{usage_default .}}
  required:    {{usage_required .}}
{{end}}
`
)

// OutputUsage is a generic function to output a list of available environment vars
// based on the given config struct.
func OutputUsage(configStruct interface{}, appName string, envarPrefix string) {
	tabs := tabwriter.NewWriter(os.Stdout, 1, 0, 4, ' ', 0)
	usageFormat := fmt.Sprintf(usageFormat, appName)
	_ = envconfig.Usagef(envarPrefix, configStruct, tabs, usageFormat) // nolint: gosec
	_ = tabs.Flush()                                                   // nolint: gosec
}
