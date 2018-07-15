package layout

type ClusterLayoutManager struct {
	ClusterIndex int
	MasterNodesManager MasterNodesManager
	WorkerNodesManager WorkerNodesManager
}

type MasterNodesManager struct {
	CurrentIndex int
}

type WorkerNodesManager struct {
	CurrentIndex int
}

func (c ClusterLayoutManager) add(nodeName string) {
	switch nodeName {
		case "master": c.MasterNodesManager.CurrentIndex++
		case "worker": c.WorkerNodesManager.CurrentIndex++
	}
}
