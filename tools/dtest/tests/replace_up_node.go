package dtests

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/m3db/m3db/tools/dtest/harness"
	"github.com/m3db/m3db/tools/dtest/util"
	"github.com/m3db/m3db/x/m3em/convert"
	m3emnode "github.com/m3db/m3db/x/m3em/node"
	xclock "github.com/m3db/m3x/clock"
)

var (
	replaceUpNodeTestCmd = &cobra.Command{
		Use:   "replace_up_node",
		Short: "Run a dtest where a node that is UP, is replaced from the cluster. Node is left UP.",
		Long:  "",
		Example: `
TODO(prateek): write up`,
		Run: replaceUpNodeDTest,
	}
)

func replaceUpNodeDTest(cmd *cobra.Command, args []string) {
	if err := globalArgs.Validate(); err != nil {
		printUsage(cmd)
		return
	}

	logger := newLogger(cmd)
	dt := harness.New(globalArgs, logger)
	defer dt.Close()

	nodes := dt.Nodes()
	numNodes := len(nodes) - 1 // leaving spare to replace with
	testCluster := dt.Cluster()

	setupNodes, err := testCluster.Setup(numNodes)
	panicIfErr(err, "unable to setup cluster")
	logger.Infof("setup cluster with %d nodes", numNodes)

	panicIfErr(testCluster.Start(), "unable to start nodes")
	logger.Infof("started cluster with %d nodes", numNodes)

	m3dbnodes, err := convert.AsM3DBNodes(setupNodes)
	panicIfErr(err, "unable to cast to m3dbnodes")

	logger.Infof("waiting until all instances are bootstrapped")
	watcher := util.NewNodesWatcher(m3dbnodes, logger, defaultBootstrapStatusReportingInterval)
	allBootstrapped := watcher.WaitUntilAll(m3emnode.Node.Bootstrapped, dt.BootstrapTimeout())
	panicIf(!allBootstrapped, fmt.Sprintf("unable to bootstrap all nodes, err = %v", watcher.PendingAsError()))
	logger.Infof("all nodes bootstrapped successfully!")

	// replace first node from the cluster
	logger.Infof("replacing node")
	replaceNode := setupNodes[0]
	newNodes, err := testCluster.ReplaceNode(replaceNode)
	panicIfErr(err, "unable to replace node")
	logger.Infof("replaced node: %s", replaceNode.ID())

	// start added nodes
	for _, n := range newNodes {
		panicIfErr(n.Start(), "unable to start node")
	}

	// wait until all shards are marked available again
	logger.Infof("waiting till all shards are available")
	allAvailable := xclock.WaitUntil(dt.AllShardsAvailable, dt.BootstrapTimeout())
	panicIf(!allAvailable, "all shards not available")
	logger.Infof("all shards available!")
}
