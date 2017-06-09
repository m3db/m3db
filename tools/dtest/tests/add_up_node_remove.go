package dtests

import (
	"github.com/spf13/cobra"

	"github.com/m3db/m3cluster/shard"
	"github.com/m3db/m3db/tools/dtest/harness"
	xclock "github.com/m3db/m3x/clock"
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
	if err := globalArgs.Validate(); err != nil {
		printUsage(cmd)
		return
	}

	logger := newLogger(cmd)
	dt := harness.New(globalArgs, logger)
	defer dt.Close()

	nodes := dt.Nodes()
	numNodes := len(nodes) - 1 // leaving one spare
	testCluster := dt.Cluster()

	logger.Infof("setting up cluster")
	setupNodes, err := testCluster.Setup(numNodes)
	panicIfErr(err, "unable to setup cluster")
	logger.Infof("setup cluster with %d nodes", numNodes)

	logger.Infof("seeding nodes with initial data")
	panicIfErr(dt.Seed(setupNodes), "unable to seed nodes")
	logger.Infof("seeded nodes")

	logger.Infof("starting cluster")
	panicIfErr(testCluster.Start(), "unable to start nodes")
	logger.Infof("started cluster with %d nodes", numNodes)

	logger.Infof("waiting until all instances are bootstrapped")
	panicIfErr(dt.WaitUntilAllBootstrapped(setupNodes), "unable to bootstrap all nodes")
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
