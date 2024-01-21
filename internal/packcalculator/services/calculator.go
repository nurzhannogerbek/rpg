package services

import (
	"sort"

	"github.com/go-playground/validator"
	"gonum.org/v1/gonum/graph/path"

	"rpg/internal/packcalculator/models"
	"rpg/utils"
)

// PackCalculator is an interface defining methods used in the code.
type PackCalculator interface {
	Calculate(quantity int) (models.RequiredPacks, error)
}

// GraphPackCalculator generates a graph of quantity permutations with the available pack sizes.
type GraphPackCalculator struct {
	PackSizes []int `validate:"required,min=1,dive,gt=0"` // PackSizes is a slice representing available pack sizes.
}

// Calculate calculates the required number of packs based on the provided quantity and available pack sizes.
func (c GraphPackCalculator) Calculate(quantity int) (models.RequiredPacks, error) {
	// Validate the input using the validator package.
	err := validator.New().Struct(c)
	if err != nil {
		return nil, err.(validator.ValidationErrors)
	}

	// Initialize the map to store the required packs.
	packs := make(models.RequiredPacks)

	// Check if the quantity is zero or negative, in which case no packs are required.
	if quantity <= 0 {
		return packs, nil
	}

	// Sort the available pack sizes in ascending order.
	sizes := c.PackSizes
	sort.Ints(sizes)

	// Reduce the problem space when the quantity is far greater than the sum of available pack sizes.
	if permutationClamp := utils.Sum(sizes) * HeadroomMultiplier; quantity > permutationClamp {
		largestSize := sizes[len(sizes)-1]
		// Subtract packs to bring the quantity down to the clamp.
		packs[largestSize] = int(float64(quantity-permutationClamp) / float64(largestSize))
		quantity -= packs[largestSize] * largestSize
	}

	// Create a graph with the initial quantity as the root node.
	qGraph := NewQuantityGraph(len(sizes))
	rootNode := NewQuantityNode(quantity)
	qGraph.AddNode(rootNode)

	// Generate permutations using the described algorithm.
	qGraph.GeneratePermutations(rootNode, sizes)

	// Aid traversal by removing unnecessary nodes.
	candidateNode := qGraph.ClosestCandidate()
	qGraph.PruneNodes(candidateNode)

	// Find the shortest path to the quantity closest to zero.
	shortest, _ := path.AStar(rootNode, candidateNode, qGraph, nil)
	shortestPath, _ := shortest.To(candidateNode.ID())
	pathLength := len(shortestPath)

	// Count each weighted line that forms the path as a used pack size.
	for i, currentNode := range shortestPath {
		nextIndex := i + 1
		if nextIndex >= pathLength {
			break
		}

		lines := qGraph.WeightedLines(currentNode.ID(), shortestPath[nextIndex].ID())
		lines.Next()
		packs[int(lines.WeightedLine().Weight())]++
	}

	return packs, nil
}

// CalculatePacks returns optimal pack sizes using GraphPackCalculator.
func CalculatePacks(orderQuantity int, packSizes []int) (models.CalculateResponse, error) {
	// Create an instance of GraphPackCalculator.
	calculator := GraphPackCalculator{PackSizes: packSizes}

	// Call the Calculate method of GraphPackCalculator.
	packs, err := calculator.Calculate(orderQuantity)
	if err != nil {
		return models.CalculateResponse{}, err
	}

	// Convert the result to the CalculateResponse structure from models.
	var response models.CalculateResponse
	for size, quantity := range packs {
		response.Packs = append(response.Packs, struct {
			PackSize int `json:"pack_size"`
			Quantity int `json:"quantity"`
		}{PackSize: size, Quantity: quantity})
	}

	return response, nil
}
