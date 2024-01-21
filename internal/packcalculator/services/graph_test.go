package services_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"rpg/internal/packcalculator/services"
)

// Helper function to compare slices regardless of order.
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// TestGeneratePermutations checks the behavior of the GeneratePermutations function.
func TestGeneratePermutations(t *testing.T) {
	// Subtest: GeneratePermutations with Valid Input.
	t.Run("GeneratePermutations with Valid Input", func(t *testing.T) {
		// Create an instance of QuantityGraph.
		graph := services.NewQuantityGraph(3)

		// Create a node and call the GeneratePermutations method.
		node := services.NewQuantityNode(10)
		sizes := []int{1, 2, 3}
		graph.GeneratePermutations(node, sizes)

		// Get all nodes in the graph.
		allNodes := graph.Nodes()
		var result []int
		for allNodes.Next() {
			result = append(result, int(allNodes.Node().ID()))
		}

		// Output the actual and expected results.
		fmt.Println("Actual result:", result)

		// Expect that all possible combinations of nodes are present.
		assert.True(t, slicesEqual(result, []int{-2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), "Generated permutations are incorrect")
	})

	// Subtest: GeneratePermutations with Empty Sizes.
	t.Run("GeneratePermutations with Empty Sizes", func(t *testing.T) {
		// Create an instance of QuantityGraph.
		graph := services.NewQuantityGraph(3)

		// Create a node and call the GeneratePermutations method with empty sizes.
		node := services.NewQuantityNode(10)
		var sizes []int
		graph.GeneratePermutations(node, sizes)

		// Expect that the slice of nodes is empty since sizes are empty.
		assert.Empty(t, graph.Nodes(), "No nodes should be generated with empty sizes")
	})

	// Subtest: GeneratePermutations with Negative Quantity.
	t.Run("GeneratePermutations with Negative Quantity", func(t *testing.T) {
		// Create an instance of QuantityGraph.
		graph := services.NewQuantityGraph(3)

		// Create a node with negative quantity and call the GeneratePermutations method.
		node := services.NewQuantityNode(-5)
		sizes := []int{1, 2, 3}
		graph.GeneratePermutations(node, sizes)

		// Expect that the nodes are added to Candidates.
		assert.NotEmpty(t, graph.Candidates, "Nodes should be added to Candidates for negative quantity")
	})

	// Subtest: GeneratePermutations with Zero Quantity.
	t.Run("GeneratePermutations with Zero Quantity", func(t *testing.T) {
		// Create an instance of QuantityGraph.
		graph := services.NewQuantityGraph(3)

		// Create a node with zero quantity and call the GeneratePermutations method.
		node := services.NewQuantityNode(0)
		sizes := []int{1, 2, 3}
		graph.GeneratePermutations(node, sizes)

		// Expect that the nodes are added to Candidates.
		assert.NotEmpty(t, graph.Candidates, "Nodes should be added to Candidates for zero quantity")
	})
}

// TestClosestCandidate checks the behavior of the ClosestCandidate function.
func TestClosestCandidate(t *testing.T) {
	// Subtest: NoCandidates.
	t.Run("NoCandidates", func(t *testing.T) {
		// Create an instance of QuantityGraph with no candidates.
		graph := services.NewQuantityGraph(3)

		// Add some nodes directly to the graph.
		graph.AddNode(services.NewQuantityNode(-2))
		graph.AddNode(services.NewQuantityNode(-5))
		graph.AddNode(services.NewQuantityNode(-1))

		// Manually add a candidate to the map to avoid an empty map.
		candidateNode := services.NewQuantityNode(-1)
		graph.Candidates[-1] = candidateNode

		// Call the ClosestCandidate method.
		candidate := graph.ClosestCandidate()

		// Expect that the candidate is the one with the closest quantity to zero.
		assert.Equal(t, candidateNode, candidate, "Unexpected candidate node")
	})

	// Subtest: SingleCandidate.
	t.Run("SingleCandidate", func(t *testing.T) {
		// Create an instance of QuantityGraph with a single candidate.
		graph := services.NewQuantityGraph(3)

		// Add a single node directly to the graph.
		node1 := services.NewQuantityNode(-2)
		node2 := services.NewQuantityNode(-5)
		graph.AddNode(node1)
		graph.AddNode(node2)

		// Create a new node with a unique ID.
		candidateNode := services.NewQuantityNode(-1)

		// Add the candidate node to the graph.
		graph.AddNode(candidateNode)

		// Call the ClosestCandidate method.
		candidate := graph.ClosestCandidate()

		// Expect that the candidate is the one with the closest quantity to zero.
		assert.Equal(t, services.NewQuantityNode(0), candidate, "Unexpected candidate node")

		// Validate that the candidate node is present in the graph nodes.
		var candidatePresent bool
		it := graph.Nodes()
		for it.Next() {
			if it.Node().ID() == candidateNode.ID() {
				candidatePresent = true
				break
			}
		}
		assert.True(t, candidatePresent, "Candidate node not present in graph nodes")
	})

	// Subtest: MultipleCandidates.
	t.Run("MultipleCandidates", func(t *testing.T) {
		// Create an instance of QuantityGraph with multiple candidates.
		graph := services.NewQuantityGraph(3)

		// Add some nodes directly to the graph.
		graph.AddNode(services.NewQuantityNode(-2))
		graph.AddNode(services.NewQuantityNode(-5))
		graph.AddNode(services.NewQuantityNode(-1))

		// Call the ClosestCandidate method.
		candidate := graph.ClosestCandidate()

		// Expect that the candidate is the one with the closest quantity to zero.
		assert.Equal(t, services.NewQuantityNode(0), candidate, "Unexpected candidate node")
	})
}

// TestPruneNodes checks the behavior of the PruneNodes function.
func TestPruneNodes(t *testing.T) {
	// Subtest: PruneNodes with Single Candidate.
	t.Run("PruneNodes with Single Candidate", func(t *testing.T) {
		// Create an instance of QuantityGraph.
		graph := services.NewQuantityGraph(3)

		// Add nodes to the graph.
		node1 := services.NewQuantityNode(1)
		node2 := services.NewQuantityNode(2)
		node3 := services.NewQuantityNode(3)
		graph.AddNode(node1)
		graph.AddNode(node2)
		graph.AddNode(node3)

		// Call PruneNodes with a candidate that is present in the graph.
		candidate := node2
		graph.PruneNodes(candidate)

		// Check if the candidate and its connected nodes are present, and others are removed.
		nodes := graph.Nodes()
		var nodeIDs []int
		for nodes.Next() {
			nodeIDs = append(nodeIDs, int(nodes.Node().ID()))
		}

		// Expect that the candidate and its connected nodes are present, and others are removed.
		expectedNodeIDs := []int{2}
		assert.ElementsMatch(t, expectedNodeIDs, nodeIDs, "Nodes should be pruned correctly")
	})

	// Subtest: PruneNodes with Multiple Candidates.
	t.Run("PruneNodes with Multiple Candidates", func(t *testing.T) {
		// Create an instance of QuantityGraph.
		graph := services.NewQuantityGraph(3)

		// Add nodes to the graph.
		node1 := services.NewQuantityNode(1)
		node2 := services.NewQuantityNode(2)
		node3 := services.NewQuantityNode(3)
		node4 := services.NewQuantityNode(4)
		graph.AddNode(node1)
		graph.AddNode(node2)
		graph.AddNode(node3)
		graph.AddNode(node4)

		// Call PruneNodes with multiple candidates.
		candidate1 := node2
		graph.PruneNodes(candidate1)

		// Check if the candidate and its connected nodes are present, and others are removed.
		nodes := graph.Nodes()
		var nodeIDs []int
		for nodes.Next() {
			nodeIDs = append(nodeIDs, int(nodes.Node().ID()))
		}

		// Expect that the candidate and its connected nodes are present, and others are removed.
		expectedNodeIDs := []int{2}
		assert.ElementsMatch(t, expectedNodeIDs, nodeIDs, "Nodes should be pruned correctly")
	})
}

// TestHasWeightedLine checks the behavior of the HasWeightedLine function.
func TestHasWeightedLine(t *testing.T) {
	// Subtest: Empty graph should not have a weighted line.
	t.Run("EmptyGraph", func(t *testing.T) {
		graph := services.NewQuantityGraph(3)
		assert.False(t, graph.HasWeightedLine(services.NewQuantityNode(1), services.NewQuantityNode(2), 3.0), "Empty graph should not have a weighted line")
	})

	// Subtest: Graph with a single node should not have a weighted line.
	t.Run("SingleNodeGraph", func(t *testing.T) {
		graph := services.NewQuantityGraph(3)
		node := services.NewQuantityNode(5)
		graph.AddNode(node)
		assert.False(t, graph.HasWeightedLine(services.NewQuantityNode(1), services.NewQuantityNode(2), 3.0), "Graph with a single node should not have a weighted line")
	})

	// Subtest: Graph with two connected nodes with positive weights should have a weighted line.
	t.Run("WeightedLinePresent", func(t *testing.T) {
		graph := services.NewQuantityGraph(3)
		node1 := services.NewQuantityNode(5)
		node2 := services.NewQuantityNode(3)
		graph.AddNode(node1)
		graph.AddNode(node2)
		graph.AddWeightedLine(node1, node2, 2)
		assert.True(t, graph.HasWeightedLine(node1, node2, 2.0), "Graph with two connected nodes with positive weights should have a weighted line")
	})

	// Subtest: Graph with two connected nodes with zero weight should not have a weighted line.
	t.Run("ZeroWeightedLine", func(t *testing.T) {
		graph := services.NewQuantityGraph(3)
		node1 := services.NewQuantityNode(5)
		node2 := services.NewQuantityNode(3)
		graph.AddNode(node1)
		graph.AddNode(node2)
		graph.AddWeightedLine(node1, node2, 0)
		assert.False(t, graph.HasWeightedLine(node1, node2, 1), "Graph with two connected nodes with zero weight should not have a weighted line")
	})

	// Subtest: Graph with two connected nodes with negative weight should not have a weighted line.
	t.Run("NegativeWeightedLine", func(t *testing.T) {
		graph := services.NewQuantityGraph(3)
		node1 := services.NewQuantityNode(5)
		node2 := services.NewQuantityNode(3)
		graph.AddNode(node1)
		graph.AddNode(node2)
		graph.AddWeightedLine(node1, node2, -2)
		assert.False(t, graph.HasWeightedLine(node1, node2, 1), "Graph with two connected nodes with negative weight should not have a weighted line")
	})
}

// TestAddWeightedLine checks the behavior of the AddWeightedLine function.
func TestAddWeightedLine(t *testing.T) {
	// Subtest: Adding a weighted line to an empty graph.
	t.Run("AddToEmptyGraph", func(t *testing.T) {
		graph := services.NewQuantityGraph(3)
		node1 := services.NewQuantityNode(1)
		node2 := services.NewQuantityNode(2)
		weight := 3.0

		graph.AddWeightedLine(node1, node2, weight)

		// Check the weighted edges.
		edges := graph.WeightedDirectedGraph.WeightedEdges()
		var found bool
		for edges.Next() {
			edge := edges.WeightedEdge()
			if edge.From().ID() == node1.ID() && edge.To().ID() == node2.ID() && edge.Weight() == weight {
				found = true
				break
			}
		}
		assert.True(t, found, "The weighted edge should be found in the graph")
	})

	// Subtest: Adding a duplicate weighted line.
	t.Run("AddDuplicateLine", func(t *testing.T) {
		graph := services.NewQuantityGraph(3)
		node1 := services.NewQuantityNode(1)
		node2 := services.NewQuantityNode(2)
		weight := 3.0

		graph.AddWeightedLine(node1, node2, weight)
		graph.AddWeightedLine(node1, node2, weight)

		// Check the weighted edges.
		edges := graph.WeightedDirectedGraph.WeightedEdges()
		var count int
		for edges.Next() {
			count++
		}
		assert.Equal(t, 1, count, "There should be one weighted edge")
	})

	// Subtest: Adding a weighted line with zero weight.
	t.Run("AddZeroWeightLine", func(t *testing.T) {
		graph := services.NewQuantityGraph(3)
		node1 := services.NewQuantityNode(1)
		node2 := services.NewQuantityNode(2)
		weight := 0.0

		graph.AddWeightedLine(node1, node2, weight)

		// Check the weighted edges.
		edges := graph.WeightedDirectedGraph.WeightedEdges()
		var found bool
		for edges.Next() {
			edge := edges.WeightedEdge()
			if edge.From().ID() == node1.ID() && edge.To().ID() == node2.ID() && edge.Weight() == weight {
				found = true
				break
			}
		}
		assert.True(t, found, "The weighted edge should be found in the graph")
	})
}
