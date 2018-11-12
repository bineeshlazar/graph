package graph

import (
	"fmt"
	"testing"
)

type node struct {
	uuid string
}

func (n *node) ID() string {
	return n.uuid
}

type env struct {
	test  *testing.T
	graph Graph
}

func (e *env) fill(count int) {

	for i := 1; i <= count; i++ {
		n := &node{fmt.Sprintf("Child%d", i)}
		e.graph.AddNode(n)
	}
	if count != e.graph.GetNodeCount() {
		e.test.Errorf("Node count mismatch. Expected %d, Actual %d", count, e.graph.GetNodeCount())
	}

	// Link them one after one
	for i := 1; i < count; i++ {
		from := fmt.Sprintf("Child%d", i)
		to := fmt.Sprintf("Child%d", i+1)
		e.graph.AddEdge(from, to, float32(i))
	}

	// Link them one after two :P
	for i := 1; i < count-1; i++ {
		from := fmt.Sprintf("Child%d", i)
		to := fmt.Sprintf("Child%d", i+2)
		e.graph.AddEdge(from, to, float32(i))
	}

	// Link them one after three :P
	for i := 1; i < count-2; i++ {
		from := fmt.Sprintf("Child%d", i)
		to := fmt.Sprintf("Child%d", i+3)
		e.graph.AddEdge(from, to, float32(i))
	}
}

func (e *env) print() {
	for _, n := range e.graph.GetNodes() {
		fmt.Println("Node : ", n.ID())
		edges, err := e.graph.GetEdges(n.ID())
		if err != nil {
			e.test.Error("Get edge failed")
		}
		for _, e := range edges {
			fmt.Println(e)
		}
	}
}

func TestGraph(t *testing.T) {
	g := NewGraph()
	tester := &env{test: t, graph: g}
	tester.fill(20)
	tester.print()
}
