package graph

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "graph:", log.Ldate|log.Ltime|log.Lshortfile)

// Node represents
type Node interface {
	ID() string
}

// Edge represents a connection between two nodes
type Edge interface {
	From() Node
	To() Node
	Weight() float32
	SetWeight(float32)
	String() string
}

// Graph represents a graph,
type Graph interface {
	// AddNode adds a node. Returns false if node already present
	AddNode(node Node) (bool, error)

	// AddEdge adds a conection between two nodes.
	// Returns false if edge already present.
	// Weight will be modified if edge already present
	AddEdge(from, to string, weight float32) (bool, error)

	// RemoveNode  removes a node from graph.
	RemoveNode(ID string) error

	// RemoveEdge removes a edge
	RemoveEdge(from, to string) error

	//GetNode gets a node from graph
	GetNode(ID string) Node

	// GetEdge get a edge from graph
	GetEdge(from, to string) (Edge, error)

	// Get number of nodes
	GetNodeCount() int

	// Get map of edges going from node 'ID'
	GetEdges(ID string) (map[string]Edge, error)

	// Get all nodes in graph
	GetNodes() []Node
}
