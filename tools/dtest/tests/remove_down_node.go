package dtests

import (
	"fmt"

	"github.com/m3db/m3db/tools/dtest/harness"
	"github.com/m3db/m3db/tools/dtest/util"
	m3emnode "github.com/m3db/m3db/x/m3em/node"

	xclock "github.com/m3db/m3x/clock"
	"github.com/spf13/cobra"
)

var removeDownNodeTestCmd = &cobra.Command{
	Use:   "remove_down_node",
	Short: "Run a dtest where a node that is down, is removed from the cluster. Node is left down.",
	Long:  "",
	Example: `
TODO(prateek): write up`,
	Run: removeDownNodeDTest,
}

func removeDownNodeDTest(cmd *cobra.Command, args []string) {
	if err := globalArgs.Validate(); err != nil {
		printUsage(cmd)
		return
	}

	logger := newLogger(cmd)
	dt := harness.New(globalArgs, logger)
	dt.SetClusterOptions(dt.ClusterOptions().
		SetNodeListener(util.NewPanicListener()))
	defer dt.Close()

	nodes := dt.Nodes()
	numNodes := len(nodes)
	testCluster := dt.Cluster()

	setupNodes, err := testCluster.Setup(numNodes)
	panicIfErr(err, "unable to setup cluster")
	logger.Infof("setup cluster with %d nodes", numNodes)

	panicIfErr(testCluster.Start(), "unable to start nodes")
	logger.Infof("started cluster with %d nodes", numNodes)

	logger.Infof("waiting until all instances are bootstrapped")
	watcher := util.NewNodesWatcher(nodes, logger, defaultBootstrapStatusReportingInterval)
	allBootstrapped := watcher.WaitUntilAll(m3emnode.Node.Bootstrapped, dt.BootstrapTimeout())
	panicIf(!allBootstrapped, fmt.Sprintf("unable to bootstrap all nodes, err = %v", watcher.PendingAsError()))
	logger.Infof("all nodes bootstrapped successfully!")

	// stop first node in the cluster
	removeNode := setupNodes[0]
	logger.Infof("bringing node down: %v", removeNode.String())
	panicIfErr(removeNode.Stop(), "unable to stop node")
	logger.Infof("node is now down")

	// remove from cluster
	logger.Infof("removing node")
	panicIfErr(testCluster.RemoveNode(removeNode), "unable to remove node")
	logger.Infof("removed node")

	// wait until all shards are marked available again
	logger.Infof("waiting till all shards are available")
	allAvailable := xclock.WaitUntil(dt.AllShardsAvailable, dt.BootstrapTimeout())
	panicIf(!allAvailable, "all shards not available")
	logger.Infof("all shards available!")
}
