package services

import (
	"sort"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/multi"
)

// GraphQuantity is an interface defining methods used in the code.
type GraphQuantity interface {
	GeneratePermutations(node QuantityNode, sizes []int)
	ClosestCandidate() QuantityNode
	PruneNodes(candidate graph.Node)
	HasWeightedLine(from, to QuantityNode, weight float64) bool
	AddWeightedLine(from, to QuantityNode, weight float64)
}

// HeadroomMultiplier is a constant multiplier used for reducing the problem space.
const HeadroomMultiplier int = 50

// QuantityGraph is a multi-graph of quantities, allowing for multiple weights (lines) between two nodes (edge).
type QuantityGraph struct {
	NodeCount  int
	Candidates map[int]QuantityNode
	*multi.WeightedDirectedGraph
}

// QuantityNode is a node in the QuantityGraph.
type QuantityNode struct {
	Quantity int
}

// ID returns the ID of the QuantityNode.
func (n QuantityNode) ID() int64 {
	return int64(n.Quantity)
}

// NewQuantityGraph creates a new QuantityGraph with the given node count.
func NewQuantityGraph(nodeCount int) *QuantityGraph {
	return &QuantityGraph{
		NodeCount:             nodeCount,
		Candidates:            make(map[int]QuantityNode),
		WeightedDirectedGraph: multi.NewWeightedDirectedGraph(),
	}
}

// NewQuantityNode creates a new QuantityNode with the given quantity.
func NewQuantityNode(quantity int) QuantityNode {
	return QuantityNode{Quantity: quantity}
}

// GeneratePermutations generates permutations by recursively subtracting quantities.
func (g *QuantityGraph) GeneratePermutations(node QuantityNode, sizes []int) {
	// Stop generating permutations if there are more paths to 0 than available quantities.
	if nodesToZero := g.To(int64(0)); nodesToZero.Len() >= g.NodeCount {
		return
	}

	for _, size := range sizes {
		// Find or create a node by the subtracted quantity.
		nextQuantity := node.Quantity - size
		nextNode := NewQuantityNode(nextQuantity)
		if existingNode := g.Node(nextNode.ID()); existingNode == nil {
			g.AddNode(nextNode)
		}

		// Maintain unique weights for edges between two quantities to avoid unnecessary recalculations.
		weight := float64(size)
		if g.HasWeightedLine(node, nextNode, weight) {
			continue
		}

		// Link the nodes by quantity.
		g.SetWeightedLine(g.NewWeightedLine(node, nextNode, weight))

		// Track nodes that satisfy the required quantity, stopping at this depth.
		if nextQuantity <= 0 {
			g.Candidates[nextQuantity] = nextNode
			continue
		}

		// Subtract from the next quantity, increasing depth.
		g.GeneratePermutations(nextNode, sizes)
	}
}

// ClosestCandidate finds the candidate node with the quantity closest to zero.
func (g *QuantityGraph) ClosestCandidate() QuantityNode {
	// No candidates, return a zero-initialized QuantityNode.
	if len(g.Candidates) == 0 {
		return NewQuantityNode(0)
	}

	// Create a slice of quantities from the map keys.
	quantities := make([]int, len(g.Candidates))
	i := 0
	for k := range g.Candidates {
		quantities[i] = k
		i++
	}

	// Reverse sort so the closest candidate is first.
	sort.Sort(sort.Reverse(sort.IntSlice(quantities)))

	return g.Candidates[quantities[0]]
}

// PruneNodes removes unnecessary nodes from the graph.
func (g *QuantityGraph) PruneNodes(candidate graph.Node) {
	// Remove other candidates from the graph.
	for _, node := range g.Candidates {
		if node != candidate {
			g.RemoveNode(node.ID())
		}
	}

	// Remove nodes which don't have any edges going out.
	var retraverse bool
	for {
		retraverse = false
		it := g.Nodes()
		for it.Next() {
			if node := it.Node(); node != candidate && len(graph.NodesOf(g.From(node.ID()))) == 0 {
				g.RemoveNode(node.ID())
				retraverse = true
			}
		}
		if !retraverse {
			break
		}
	}
}

// HasWeightedLine checks if there is a weighted line between two nodes with a specific weight.
func (g *QuantityGraph) HasWeightedLine(from, to QuantityNode, weight float64) bool {
	for _, line := range graph.WeightedLinesOf(g.WeightedLines(from.ID(), to.ID())) {
		if line.Weight() == weight {
			return true
		}
	}
	return false
}

// AddWeightedLine adds a weighted line from one node to another with the specified weight.
func (g *QuantityGraph) AddWeightedLine(from, to QuantityNode, weight float64) {
	g.WeightedDirectedGraph.SetWeightedLine(g.NewWeightedLine(from, to, weight))
}
