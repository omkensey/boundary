
package targets

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/targets"
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

	
	extraCmdVars
	
}

var flagsMap = map[string][]string{
	
	
	"read": {"id"},
	
	
	
	
	
	
	"delete": {"id"},
	
	
	
	
	
	
	"list": {"scope-id"},
	
	

	
	foo
	
}

func (c *Command) Synopsis() string {
	
	if extra := c.extraSynopsisFunc(); extra != "" {
		return extra
	} 
	
	return common.SynopsisFunc(c.Func, "target")
}

