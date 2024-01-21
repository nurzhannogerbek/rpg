package models

// RequiredPacks is a map of pack sizes and the number required.
type RequiredPacks map[int]int

// Pack represents information about a pack size and quantity.
type Pack struct {
	PackSize int `json:"pack_size"`
	Quantity int `json:"quantity"`
}

// CalculateResponse represents the JSON response structure.
type CalculateResponse struct {
	Packs []Pack `json:"packs"`
}
