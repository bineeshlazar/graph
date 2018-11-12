package graph

import (
	"errors"
	"sync"
)

type nodeInfo struct {
	Node

	//map of edges indexed by far node ID
	edges map[string]Edge
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

func (g *graph) AddNode(node Node) (bool, error) {
	n := g.GetNode(node.ID())
	if n != nil {
		logger.Printf("Node %s already present\n", node.ID())
		return false, nil
	}
	g.lock.Lock()
	g.nodes[node.ID()] = &nodeInfo{Node: node, edges: make(map[string]Edge)}
	g.lock.Unlock()
	return true, nil
}

func (g *graph) AddEdge(from, to string, weight float32) (bool, error) {

	var e Edge
	newEdge := false

	e, err := g.GetEdge(from, to)
	if err != nil {
		logger.Println(err.Error())
		return false, err
	}

	fromNode := g.GetNode(from)
	toNode := g.GetNode(to)

	if e == nil {
		newEdge = true
		e = NewEdge(fromNode, toNode, weight)
		g.lock.Lock()
		g.nodes[from].edges[to] = e
		g.lock.Unlock()
	} else {
		e.SetWeight(weight)
	}

	return newEdge, nil
}

func (g *graph) RemoveNode(ID string) error {
	n := g.GetNode(ID)
	if n == nil {
		err := errors.New("RemoveNode: could not find node")
		logger.Println(err)
		return err
	}

	g.lock.Lock()
	delete(g.nodes, ID)
	g.lock.Unlock()
	return nil
}

func (g *graph) RemoveEdge(from, to string) error {

	e, err := g.GetEdge(from, to)
	if err != nil {
		logger.Println(err.Error())
		return err
	}

	if e == nil {
		err := errors.New("RemoveEdge: could not find edge")
		logger.Println(err.Error())
		return err
	}

	g.lock.Lock()
	delete(g.nodes[from].edges, to)
	g.lock.Unlock()
	return nil
}

func (g *graph) GetNode(ID string) Node {
	g.lock.RLock()
	n, found := g.nodes[ID]
	g.lock.RUnlock()
	if !found {
		return nil
	}
	return n
}

func (g *graph) GetEdge(from, to string) (Edge, error) {

	fromNode := g.GetNode(from)
	toNode := g.GetNode(to)
	if (fromNode == nil) || (toNode == nil) {
		err := errors.New("Could not find nodes")
		logger.Println(err.Error())
		return nil, err
	}

	info, ok := fromNode.(*nodeInfo)
	if !ok {
		err := errors.New("Invalid type for node")
		logger.Println(err.Error())
		return nil, err
	}

	g.lock.RLock()
	e, _ := info.edges[to]
	g.lock.RUnlock()
	return e, nil
}

func (g *graph) GetNodeCount() int {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return len(g.nodes)
}

func (g *graph) GetEdges(ID string) (map[string]Edge, error) {
	node := g.GetNode(ID)
	if node == nil {
		err := errors.New("Could not find node")
		logger.Println(err.Error())
		return nil, err
	}

	info, ok := node.(*nodeInfo)
	if !ok {
		err := errors.New("Invalid type for node")
		logger.Println(err.Error())
		return nil, err
	}

	return info.edges, nil
}

func (g *graph) GetNodes() []Node {
	nodes := make([]Node, 0, len(g.nodes))
	for _, n := range g.nodes {
		nodes = append(nodes, n)
	}
	return nodes
}
