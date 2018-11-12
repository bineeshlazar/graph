package graph

// Node represents
type Node interface {
	ID() string
}

// Edge represents a connection between two nodes
type Edge interface {
	From() Node
	To() Node
	GetWeight() float32
}

// Graph represents a graph
type Graph interface {
	AddNode(node Node) bool
	AddEdge(from, to Node, weight float32) bool
	RemoveNode(node Node) error
	RemoveEdge(edge Edge) error
	GetNode(ID string) (Node, error)
	GetEdge(from, to string) (Edge, error)
	GetNodeCount() int
	GetEdges(node Node) ([]Edge, error)
}
