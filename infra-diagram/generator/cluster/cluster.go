package cluster

import (
	"github.com/bengtrj/cfcr-cluster-diagram/infra-diagram/generator/deployment"
	"fmt"
	"strings"
	"sort"
	)

type Cluster struct {
	Name    string
	Masters []Node
	Workers []Node
}

type Node struct {
	Id   string
	Name string
	Type string
	Jobs []Job
}

type Job struct {
	Id           string
	ParentId     string
	Name         string
	Position     Position
	CssClasses   []string
	IsK8sProcess bool
}

type Position struct {
	X int
	Y int
}

type LayoutPosition struct {
	InitialX          int
	InitialY          int
	ParentIncrementX  int
	ParentIncrementY  int
	SiblingIncrementX int
	SiblingIncrementY int
}

var greatestKnowPosition = Position{
	X: 0,
	Y: 0,
}

func NewCluster(d deployment.Deployment) (Cluster, error) {

	cluster := Cluster{
		Name: d.Name,
	}

	for _, n := range d.Nodes {

		switch n.Name {
		case "master":
			cluster.Masters = getNodes(n, LayoutPosition{
				InitialX:          0,
				InitialY:          0,
				ParentIncrementX:  0,
				ParentIncrementY:  280,
				SiblingIncrementX: 100,
				SiblingIncrementY: 0,
			})
		case "worker":
			cluster.Workers = getNodes(n, LayoutPosition{
				InitialX:          greatestKnowPosition.X + 150,
				InitialY:          0,
				ParentIncrementX:  0,
				ParentIncrementY:  280,
				SiblingIncrementX: 100,
				SiblingIncrementY: 0,
			})
		}

	}

	return cluster, nil

}

func getNodes(n deployment.Node, layoutPosition LayoutPosition) []Node {
	var nodes []Node
	for i := 0; i < n.Instances; i++ {
		id := fmt.Sprintf("%s-%d", n.Name, i+1)
		nodes = append(nodes, Node{
			Id:   id,
			Type: n.Name,
			Name: fmt.Sprintf("%s %d", n.Name, i+1),
			Jobs: getJobs(i, id, n, layoutPosition),
		})
	}
	return nodes
}

func getJobs(parentIndex int, parentId string, node deployment.Node, layoutPosition LayoutPosition) []Job {
	js := node.Jobs
	var clusterJobs []Job
	initialPosition := Position{
		X: layoutPosition.InitialX,
		Y: layoutPosition.InitialY,
	}
	for i, job := range js {
		if shouldInclude(job) {
			newJob := Job{
				Id:           fmt.Sprintf("%s-%d-%s", job.Name, i+1, parentId),
				Name:         strings.TrimPrefix(job.Name, "kube-"),
				ParentId:     parentId,
				CssClasses:   []string{fmt.Sprintf("%s-process", node.Name), job.Name},
				IsK8sProcess: strings.HasPrefix(job.Name, "kube"),
			}
			if newJob.IsK8sProcess {
				newJob.CssClasses = append(newJob.CssClasses, "k8s-process")
			}
			clusterJobs = append(clusterJobs, newJob)
		}
	}

	sortJobs(clusterJobs)
	setupPositions(clusterJobs, initialPosition, layoutPosition, parentIndex)

	return clusterJobs
}

func shouldInclude(job deployment.Job) bool {
	switch job.Name {
	case "apply-specs", "smoke-tests", "secure-var-vcap", "cloud-provider", "bpm", "kubernetes-dependencies", "kubernetes-roles":
		return false
	default:
		return true
	}
}

func sortJobs(clusterJobs []Job) {
	sort.Slice(clusterJobs, func(i, j int) bool {
		if clusterJobs[i].IsK8sProcess && !clusterJobs[j].IsK8sProcess {
			return true
		}
		if clusterJobs[i].IsK8sProcess && clusterJobs[j].IsK8sProcess || !clusterJobs[i].IsK8sProcess && !clusterJobs[j].IsK8sProcess {
			return clusterJobs[i].Name < clusterJobs[j].Name
		}
		return false
	})
}

func setupPositions(clusterJobs []Job, initialPosition Position, layoutPosition LayoutPosition, parentIndex int) {
	indexInRow := 0
	rowNum := 0
	lastPosition := Position{
		X: initialPosition.X,
		Y: initialPosition.Y,
	}
	for i := range clusterJobs {

		if !clusterJobs[i].IsK8sProcess && rowNum == 0 {
			rowNum++
			indexInRow = 0
			lastPosition.X = initialPosition.X
			lastPosition.Y = initialPosition.Y + 100
		}

		clusterJobs[i].Position = layoutPosition.calculateNextPosition(parentIndex, indexInRow, lastPosition)
		greatestKnowPosition.X = max(greatestKnowPosition.X, clusterJobs[i].Position.X)
		greatestKnowPosition.Y = max(greatestKnowPosition.Y, clusterJobs[i].Position.Y)
		indexInRow++
	}
}

func max(i int, j int) int {
	if i < j {
		return j
	}
	return i
}

func (l LayoutPosition) calculateNextPosition(parentIndex int, childIndex int, previous Position) Position {
	return Position{
		X: previous.X + (l.ParentIncrementX * parentIndex) + (l.SiblingIncrementX * childIndex),
		Y: previous.Y + (l.ParentIncrementY * parentIndex) + (l.SiblingIncrementY * childIndex),
	}
}


