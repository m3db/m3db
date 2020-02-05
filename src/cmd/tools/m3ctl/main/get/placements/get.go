package placements

import (
	"fmt"

	"github.com/m3db/m3/src/cmd/tools/m3ctl/main/client"
	"github.com/m3db/m3/src/cmd/tools/m3ctl/main/globalopts"
	"github.com/m3db/m3/src/cmd/tools/m3ctl/main/placements"
)

func doGet(s *placementVals, globals globalopts.GlobalOpts) error {
	url := fmt.Sprintf("%s%s", globals.Endpoint, placements.DefaultPath)
	return client.DoGet(url, client.Dumper, globals.Zap)
}
