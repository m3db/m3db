// +build integration

package integration

import (
	 _ "github.com/m3db/m3db/client"
	 _ "github.com/m3db/m3db/clock"
	 _ "github.com/m3db/m3db/digest"
	 _ "github.com/m3db/m3db/encoding"
	 _ "github.com/m3db/m3db/encoding/m3tsz"
	 _ "github.com/m3db/m3db/encoding/testgen"
	 _ "github.com/m3db/m3db/environment"
	 _ "github.com/m3db/m3db/generated/mocks"
	 _ "github.com/m3db/m3db/generated/proto"
	 _ "github.com/m3db/m3db/generated/proto/index"
	 _ "github.com/m3db/m3db/generated/proto/namespace"
	 _ "github.com/m3db/m3db/generated/proto/pagetoken"
	 _ "github.com/m3db/m3db/generated/thrift"
	 _ "github.com/m3db/m3db/generated/thrift/rpc"
	 _ "github.com/m3db/m3db/kvconfig"
	 _ "github.com/m3db/m3db/network/server"
	 _ "github.com/m3db/m3db/network/server/httpjson"
	 _ "github.com/m3db/m3db/network/server/httpjson/cluster"
	 _ "github.com/m3db/m3db/network/server/httpjson/node"
	 _ "github.com/m3db/m3db/network/server/tchannelthrift"
	 _ "github.com/m3db/m3db/network/server/tchannelthrift/cluster"
	 _ "github.com/m3db/m3db/network/server/tchannelthrift/convert"
	 _ "github.com/m3db/m3db/network/server/tchannelthrift/errors"
	 _ "github.com/m3db/m3db/network/server/tchannelthrift/node"
	 _ "github.com/m3db/m3db/network/server/tchannelthrift/node/channel"
	 _ "github.com/m3db/m3db/persist"
	 _ "github.com/m3db/m3db/persist/fs"
	 _ "github.com/m3db/m3db/persist/fs/clone"
	 _ "github.com/m3db/m3db/persist/fs/commitlog"
	 _ "github.com/m3db/m3db/persist/fs/msgpack"
	 _ "github.com/m3db/m3db/persist/schema"
	 _ "github.com/m3db/m3db/ratelimit"
	 _ "github.com/m3db/m3db/retention"
	 _ "github.com/m3db/m3db/runtime"
	 _ "github.com/m3db/m3db/serialize"
	 _ "github.com/m3db/m3db/services/m3dbnode/config"
	 _ "github.com/m3db/m3db/services/m3dbnode/server"
	 _ "github.com/m3db/m3db/sharding"
	 _ "github.com/m3db/m3db/storage"
	 _ "github.com/m3db/m3db/storage/block"
	 _ "github.com/m3db/m3db/storage/bootstrap"
	 _ "github.com/m3db/m3db/storage/bootstrap/bootstrapper"
	 _ "github.com/m3db/m3db/storage/bootstrap/bootstrapper/commitlog"
	 _ "github.com/m3db/m3db/storage/bootstrap/bootstrapper/fs"
	 _ "github.com/m3db/m3db/storage/bootstrap/bootstrapper/peers"
	 _ "github.com/m3db/m3db/storage/bootstrap/result"
	 _ "github.com/m3db/m3db/storage/cluster"
	 _ "github.com/m3db/m3db/storage/index"
	 _ "github.com/m3db/m3db/storage/index/convert"
	 _ "github.com/m3db/m3db/storage/namespace"
	 _ "github.com/m3db/m3db/storage/repair"
	 _ "github.com/m3db/m3db/storage/series"
	 _ "github.com/m3db/m3db/tools"
	 _ "github.com/m3db/m3db/tools/dtest/config"
	 _ "github.com/m3db/m3db/tools/dtest/harness"
	 _ "github.com/m3db/m3db/tools/dtest/tests"
	 _ "github.com/m3db/m3db/tools/dtest/util"
	 _ "github.com/m3db/m3db/tools/dtest/util/seed"
	 _ "github.com/m3db/m3db/topology"
	 _ "github.com/m3db/m3db/topology/testutil"
	 _ "github.com/m3db/m3db/ts"
	 _ "github.com/m3db/m3db/x/m3em/convert"
	 _ "github.com/m3db/m3db/x/m3em/node"
	 _ "github.com/m3db/m3db/x/metrics"
	 _ "github.com/m3db/m3db/x/mmap"
	 _ "github.com/m3db/m3db/x/tchannel"
	 _ "github.com/m3db/m3db/x/test"
	 _ "github.com/m3db/m3db/x/xcounter"
	 _ "github.com/m3db/m3db/x/xio"
	 _ "github.com/m3db/m3db/x/xpool"
)

