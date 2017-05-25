package dtests

import (
	"fmt"
	"os"

	"github.com/m3db/m3db/tools/dtest/harness"
	"github.com/m3db/m3db/tools/dtest/util"
	m3dbnode "github.com/m3db/m3db/x/m3em/node"

	"github.com/m3db/m3cluster/shard"
	xclock "github.com/m3db/m3x/clock"
	"github.com/m3db/m3x/log"
	"github.com/spf13/cobra"
)

var (
	addUpNodeRemoveTestCmd = &cobra.Command{
		Use:   "add_up_node_remove",
		Short: "Run a dtest where a node that is UP, is added to the cluster. Node is removed as it begins bootstrapping.",
		Long:  "",
		Example: `
TODO(prateek): write up`,
		Run: addUpNodeRemoveDTest,
	}
)

func addUpNodeRemoveDTest(cmd *cobra.Command, args []string) {
	if err := globalCLIOpts.Validate(); err != nil {
		printUsage(cmd)
		return
	}
	logger := xlog.NewLogger(os.Stdout)
	logger.Infof("============== %v  ==============", cmd.Name())

	dt := harness.New(globalCLIOpts, logger)
	dt.SetClusterOptions(dt.ClusterOptions().
		SetNodeListener(util.NewPanicListener()))
	defer dt.Close()

	nodes := dt.Nodes()
	numNodes := len(nodes) - 1 // leaving one spare
	testCluster := dt.Cluster()

	setupNodes, err := testCluster.Setup(numNodes)
	panicIfErr(err, "unable to setup cluster")
	logger.Infof("setup cluster with %d nodes", numNodes)

	panicIfErr(testCluster.Start(), "unable to start nodes")
	logger.Infof("started cluster with %d nodes", numNodes)

	m3dbnodes, err := util.AsM3DBNodes(setupNodes)
	panicIfErr(err, "unable to cast to m3dbnodes")

	logger.Infof("waiting until all instances are bootstrapped")
	watcher := util.NewM3DBNodesWatcher(m3dbnodes)
	allBootstrapped := watcher.WaitUntilAll(m3dbnode.Node.Bootstrapped, dt.BootstrapTimeout())
	panicIf(!allBootstrapped, fmt.Sprintf("unable to bootstrap all nodes, err = %v", watcher.PendingAsError()))
	logger.Infof("all nodes bootstrapped successfully!")

	// get a spare, ensure it's up and add to the cluster
	logger.Infof("adding spare to the cluster")
	spares := testCluster.SpareNodes()
	panicIf(len(spares) < 1, "no spares to add to the cluster")
	spare := spares[0]

	// start node
	logger.Infof("starting new node: %v", spare.ID())
	panicIfErr(spare.Start(), "unable to start node")
	logger.Infof("started node")

	// add to placement
	logger.Infof("adding node")
	panicIfErr(testCluster.AddSpecifiedNode(spare), "unable to add node")
	logger.Infof("added node")

	// NB(prateek): ideally we'd like to wait until the node begins bootstrapping, but we don't
	// have a way to capture that node status. The rpc endpoint in m3dbnode only captures bootstrap
	// status at the database level, and m3kv only captures state once a shard is marked as bootstrapped.
	// So here we wait until any shard is marked as bootstrapped before continuing.

	// wait until any shard is bootstrapped (i.e. marked available on new node)
	logger.Infof("waiting till any shards are bootstrapped on new node")
	timeout := dt.BootstrapTimeout() / 10
	anyBootstrapped := xclock.WaitUntil(func() bool { return dt.AnyInstanceShardHasState(spare.ID(), shard.Available) }, timeout)
	panicIf(!anyBootstrapped, "all shards not available")

	// remove the node once it has a shard available
	logger.Infof("node has at least 1 shard available. removing node")
	panicIfErr(testCluster.RemoveNode(spare), "unable to remove node")
	logger.Infof("removed node")
}
