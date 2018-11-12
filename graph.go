package graph

import (
	"sync"
)

type nodeInfo struct {
	node  Node
	edges []Edge
}

type graph struct {
	nodes map[string]*nodeInfo
	lock  sync.RWMutex
}

// NewGraph returns a new graph
func NewGraph() Graph {
	g := &graph{}
	g.nodes = make(map[string]*nodeInfo)
	return g
}
