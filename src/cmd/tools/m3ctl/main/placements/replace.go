package placements

import (
	"fmt"
	"github.com/m3db/m3/src/cmd/tools/m3ctl/main/http"
	"github.com/m3db/m3/src/cmd/tools/m3ctl/main/yaml"
	"github.com/m3db/m3/src/query/generated/proto/admin"
)

func Replace(flags *PlacementArgs, endpoint string) {
	data := yaml.Load(flags.replaceFlag.Value[0], &admin.PlacementReplaceRequest{})
	url := fmt.Sprintf("%s%s%s", endpoint, defaultPath, "/replace")
	http.DoPost(url, data, http.DoDump)
	return
}
