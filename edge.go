package graph

import (
	"fmt"
)

// Edge represents a connection between two nodes
type edge struct {
	from   Node
	to     Node
	weight float32
}

//NewEdge returns a Edge interface
func NewEdge(from, to Node, weight float32) Edge {
	return &edge{
		from:   from,
		to:     to,
		weight: weight,
	}
}

func (e *edge) From() Node {
	return e.from
}

func (e *edge) To() Node {
	return e.to
}

func (e *edge) Weight() float32 {
	return e.weight
}

func (e *edge) SetWeight(w float32) {
	e.weight = w
}

func (e *edge) String() string {
	return fmt.Sprintf("(%s)--%f-->(%s)", e.from.ID(), e.weight, e.to.ID())
}
