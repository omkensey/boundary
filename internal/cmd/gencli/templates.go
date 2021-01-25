package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func fillTemplates() {
	for pkg, data := range inputStructs {
		outBuf := new(bytes.Buffer)

		if err := cmdTemplate.Execute(outBuf, data); err != nil {
			fmt.Printf("error executing struct template for resource %s: %v\n", pkg, err)
			os.Exit(1)
		}

		outFile, err := filepath.Abs(fmt.Sprintf("%s/%s/%s.gen.go", os.Getenv("CLI_GEN_BASEPATH"), pkg, data.ResourceType))
		if err != nil {
			fmt.Printf("error opening file for package %s: %v\n", pkg, err)
			os.Exit(1)
		}
		outDir := filepath.Dir(outFile)
		if _, err := os.Stat(outDir); os.IsNotExist(err) {
			_ = os.Mkdir(outDir, os.ModePerm)
		}
		if err := ioutil.WriteFile(outFile, outBuf.Bytes(), 0o644); err != nil {
			fmt.Printf("error writing file %q: %v\n", outFile, err)
			os.Exit(1)
		}
	}
}

var cmdTemplate = template.Must(template.New("").Parse(`
package {{ .ResourceType }}s

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/boundary/api"
	"{{ .PkgPath }}"
	"github.com/hashicorp/boundary/internal/cmd/base"
	"github.com/hashicorp/boundary/internal/cmd/common"
	"github.com/hashicorp/boundary/internal/types/resource"
	"github.com/hashicorp/boundary/sdk/strutil"
	"github.com/mitchellh/cli"
	"github.com/posener/complete"
)

var (
	_ cli.Command             = (*Command)(nil)
	_ cli.CommandAutocomplete = (*Command)(nil)
)

type Command struct {
	*base.Command

	Func string

	{{ if .HasExtraCommandVars }}
	extraCmdVars
	{{ end }}
}

var flagsMap = map[string][]string{
	{{ range $i, $action := .StdActions }}
	{{ if eq $action "read" }}
	"read": {"id"},
	{{ end }}
	{{ if eq $action "delete" }}
	"delete": {"id"},
	{{ end }}
	{{ if eq $action "list" }}
	"list": {"scope-id"},
	{{ end }}
	{{ end }}

	{{ range $action, $flags .ExtraFlagsMap }}
	
	{{ end }}
}

func (c *Command) Synopsis() string {
	{{ if .HasExtraSynopsisFunc }}
	if extra := c.extraSynopsisFunc(); extra != "" {
		return extra
	} 
	{{ end }}
	return common.SynopsisFunc(c.Func, "{{ .ResourceType }}")
}

`))
