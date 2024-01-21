package services_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"rpg/internal/packcalculator/services"
)

// TestCalculatePacks_BaseExample verifies the calculation of packs for the provided example.
func TestCalculatePacks_BaseExample(t *testing.T) {
	orderQuantity := 12001
	packSizes := []int{250, 500, 1000, 2000, 5000}

	result, err := services.CalculatePacks(orderQuantity, packSizes)

	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, 3, len(result.Packs), "Incorrect number of packs")

	expectedPacks := []struct {
		PackSize int `json:"pack_size"`
		Quantity int `json:"quantity"`
	}{
		{5000, 2},
		{2000, 1},
		{250, 1},
	}

	// Convert the result.Packs to the expected type.
	var actualPacks []struct {
		PackSize int `json:"pack_size"`
		Quantity int `json:"quantity"`
	}
	for _, pack := range result.Packs {
		actualPacks = append(actualPacks, struct {
			PackSize int `json:"pack_size"`
			Quantity int `json:"quantity"`
		}{PackSize: pack.PackSize, Quantity: pack.Quantity})
	}

	// Sort both slices.
	sort.Slice(expectedPacks, func(i, j int) bool {
		return expectedPacks[i].PackSize < expectedPacks[j].PackSize
	})
	sort.Slice(actualPacks, func(i, j int) bool {
		return actualPacks[i].PackSize < actualPacks[j].PackSize
	})

	// Use assert.Equal to compare slices directly without sorting.
	assert.Equal(t, expectedPacks, actualPacks, "Incorrect packs")
}

// TestCalculatePacks_CustomSize verifies the calculation of packs for the provided example.
func TestCalculatePacks_CustomSize(t *testing.T) {
	orderQuantity := 263
	packSizes := []int{23, 31, 53}

	result, err := services.CalculatePacks(orderQuantity, packSizes)

	assert.NoError(t, err, "Unexpected error")
	assert.Equal(t, 2, len(result.Packs), "Incorrect number of packs")

	expectedPacks := []struct {
		PackSize int `json:"pack_size"`
		Quantity int `json:"quantity"`
	}{
		{PackSize: 31, Quantity: 7},
		{PackSize: 23, Quantity: 2},
	}

	// Convert the result.Packs to the expected type.
	var actualPacks []struct {
		PackSize int `json:"pack_size"`
		Quantity int `json:"quantity"`
	}
	for _, pack := range result.Packs {
		actualPacks = append(actualPacks, struct {
			PackSize int `json:"pack_size"`
			Quantity int `json:"quantity"`
		}{PackSize: pack.PackSize, Quantity: pack.Quantity})
	}

	// Sort both slices.
	sort.Slice(expectedPacks, func(i, j int) bool {
		return expectedPacks[i].PackSize < expectedPacks[j].PackSize
	})
	sort.Slice(actualPacks, func(i, j int) bool {
		return actualPacks[i].PackSize < actualPacks[j].PackSize
	})

	// Use assert.Equal to compare slices directly without sorting.
	assert.Equal(t, expectedPacks, actualPacks, "Incorrect packs")
}

// TestCalculatePacks_ZeroQuantity verifies that no packs are returned for zero quantity.
func TestCalculatePacks_ZeroQuantity(t *testing.T) {
	orderQuantity := 0
	packSizes := []int{2, 5, 10}

	result, err := services.CalculatePacks(orderQuantity, packSizes)

	assert.NoError(t, err, "Unexpected error")
	assert.Empty(t, result.Packs, "Packs should be empty for zero quantity")
}

// TestCalculatePacks_EmptySizes verifies an error is returned when pack sizes are empty.
func TestCalculatePacks_EmptySizes(t *testing.T) {
	orderQuantity := 10
	var packSizes []int

	result, err := services.CalculatePacks(orderQuantity, packSizes)

	assert.Error(t, err, "Expected error for empty pack sizes")
	assert.Nil(t, result.Packs, "Packs should be nil for empty pack sizes error")
}

// TestCalculatePacks_ZeroPackSizes verifies an error is returned when all pack sizes are zero.
func TestCalculatePacks_ZeroPackSizes(t *testing.T) {
	orderQuantity := 10
	packSizes := []int{0, 0, 0}

	result, err := services.CalculatePacks(orderQuantity, packSizes)

	assert.Error(t, err, "Expected error for zero pack sizes")
	assert.Nil(t, result.Packs, "Packs should be nil for zero pack sizes error")
}

// TestCalculatePacks_NegativePackSizes verifies an error is returned when any pack size is negative.
func TestCalculatePacks_NegativePackSizes(t *testing.T) {
	orderQuantity := 10
	packSizes := []int{2, -5, 10}

	result, err := services.CalculatePacks(orderQuantity, packSizes)

	assert.Error(t, err, "Expected error for negative pack sizes")
	assert.Nil(t, result.Packs, "Packs should be nil for negative pack sizes error")
}
