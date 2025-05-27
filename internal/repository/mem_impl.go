package repository

import (
	"fmt"
	"sort"

	"github.com/alishercodecrafter/orderpackscalculator/internal/model"
	"github.com/alishercodecrafter/orderpackscalculator/internal/service"
)

// MemoryRepository implements service.PacksRepository using in-memory storage
type MemoryRepository struct {
	packs model.Packs
}

// NewMemoryRepository creates a new MemoryRepository
func NewMemoryRepository() *MemoryRepository {
	// Initialize with default pack sizes
	return &MemoryRepository{
		packs: model.Packs{
			{PackSize: 250},
			{PackSize: 500},
			{PackSize: 1000},
			{PackSize: 2000},
			{PackSize: 5000},
		},
	}
}

// Ensure MemoryRepository implements service.PacksRepository
var _ service.PacksRepository = (*MemoryRepository)(nil)

// GetPacks returns all available packs
func (r *MemoryRepository) GetPacks() model.Packs {
	// Make a copy to prevent external modification
	result := make(model.Packs, len(r.packs))
	copy(result, r.packs)

	// Sort by pack size
	sort.Slice(result, func(i, j int) bool {
		return result[i].PackSize < result[j].PackSize
	})

	return result
}

// AddPack adds a new pack
func (r *MemoryRepository) AddPack(pack model.Pack) error {
	// Check if pack size already exists
	for _, p := range r.packs {
		if p.PackSize == pack.PackSize {
			return fmt.Errorf("pack size %d already exists", pack.PackSize)
		}
	}
	r.packs = append(r.packs, pack)

	return nil
}

// RemovePack removes a pack by its size
func (r *MemoryRepository) RemovePack(packSize int) error {
	for i, pack := range r.packs {
		if pack.PackSize == packSize {
			// Remove the pack by replacing it with the last element and truncating
			r.packs[i] = r.packs[len(r.packs)-1]
			r.packs = r.packs[:len(r.packs)-1]
			return nil
		}
	}
	return fmt.Errorf("pack size %d not found", packSize)
}
