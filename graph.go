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
	if _, found := g.nodes[node.ID()]; found {
		logger.Printf("Node %s already present\n", node.ID())
		return false, nil
	}
	g.nodes[node.ID()] = &nodeInfo{Node: node, edges: make(map[string]Edge)}
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
		g.nodes[from].edges[to] = e
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

	delete(g.nodes, ID)
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

	delete(g.nodes[from].edges, to)
	return nil
}

func (g *graph) GetNode(ID string) Node {
	n, found := g.nodes[ID]
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

	e, _ := info.edges[to]
	return e, nil
}

func (g *graph) GetNodeCount() int {
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
