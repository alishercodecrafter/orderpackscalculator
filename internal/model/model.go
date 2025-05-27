package model

// Pack represents a pack entity with its properties
type Pack struct {
	PackSize int `json:"packSize" binding:"required"`
	// More fields can be added in the future
}

// Packs represents a collection of Pack entities
type Packs []Pack

// AddPackRequest represents a request to add a new pack
type AddPackRequest struct {
	Pack Pack `json:"pack" binding:"required"`
}

// CalculationRequest represents a request to calculate packs
type CalculationRequest struct {
	OrderSize int `json:"orderSize"`
}

// CalculationResponse represents the result of a pack calculation
type CalculationResponse struct {
	OrderSize int         `json:"orderSize"`
	Packs     map[int]int `json:"packs"` // map of pack size to count
}
