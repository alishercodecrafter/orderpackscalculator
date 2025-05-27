// internal/service/service.go
package service

import (
	"fmt"
	"slices"

	"github.com/alishercodecrafter/orderpackscalculator/internal/model"
)

// PacksRepository defines the interface for pack data storage
type PacksRepository interface {
	// GetPacks returns all available packs
	GetPacks() model.Packs
	// AddPack adds a new pack
	AddPack(pack model.Pack) error
	// RemovePack removes a pack by its size
	RemovePack(packSize int) error
}

// PacksServiceImpl handles the business logic for pack calculations
type PacksServiceImpl struct {
	repo PacksRepository
}

// NewPacksService creates a new PacksServiceImpl
func NewPacksService(repo PacksRepository) *PacksServiceImpl {
	return &PacksServiceImpl{
		repo: repo,
	}
}

// GetPackSizes returns all available pack sizes
func (s *PacksServiceImpl) GetPacks() model.Packs {
	return s.repo.GetPacks()
}

// AddPack adds a new pack
func (s *PacksServiceImpl) AddPack(pack model.Pack) error {
	if pack.PackSize <= 0 {
		return fmt.Errorf("pack size must be greater than zero")
	}

	return s.repo.AddPack(pack)
}

// RemovePackSize removes a pack size
func (s *PacksServiceImpl) RemovePack(packSize int) error {
	return s.repo.RemovePack(packSize)
}

// CalculatePacks calculates the optimal number of packs needed for an order
func (s *PacksServiceImpl) CalculatePacks(orderSize int) (model.CalculationResponse, error) {
	packList := s.repo.GetPacks()
	packsRule2 := make(map[int]int)

	// If no packList or invalid order size, return empty packsRule2
	if len(packList) == 0 {
		return model.CalculationResponse{}, fmt.Errorf("available packsRule2 list is empty")
	}

	if orderSize <= 0 {
		return model.CalculationResponse{}, fmt.Errorf("order size must be greater than zero")
	}

	// Extract pack sizes for calculation
	packSizes := make([]int, len(packList))
	for i, pack := range packList {
		packSizes[i] = pack.PackSize
	}
	// Sort pack sizes in ascending order
	slices.Sort(packSizes)

	packsRule3 := make(map[int]int)
	originalOrderSize := orderSize

	// If order size is larger than the largest pack,
	// we can use the largest pack first
	maxPackSize := packSizes[len(packSizes)-1]
	if orderSize > maxPackSize {
		count := orderSize / maxPackSize
		orderSize -= count * maxPackSize
		packsRule2[maxPackSize] = count
		packsRule3[maxPackSize] = count
	}

	s.calculatePacks(orderSize, packsRule2, &packSizes, false)
	s.calculatePacks(orderSize, packsRule3, &packSizes, true)

	amount2, count2 := getAmountOfItemsInPacks(packsRule2)
	amount3, count3 := getAmountOfItemsInPacks(packsRule3)
	switch {
	case amount2 < amount3:
		return model.CalculationResponse{
			OrderSize: originalOrderSize,
			Packs:     packsRule2,
		}, nil
	case amount3 < amount2:
		return model.CalculationResponse{
			OrderSize: originalOrderSize,
			Packs:     packsRule3,
		}, nil
	default:
		if count2 < count3 {
			return model.CalculationResponse{
				OrderSize: originalOrderSize,
				Packs:     packsRule2,
			}, nil
		}

		return model.CalculationResponse{
			OrderSize: originalOrderSize,
			Packs:     packsRule3,
		}, nil
	}
}

func getAmountOfItemsInPacks(packs map[int]int) (int, int) {
	amount := 0
	totalCount := 0
	for packSize, count := range packs {
		amount += packSize * count
		totalCount += count
	}

	return amount, totalCount
}

// calculatePacks is a helper function to calculate the optimal number of packs needed for an order
func (s *PacksServiceImpl) calculatePacks(orderSize int, packs map[int]int, sortedPackSizes *[]int, leastFewPacks bool) {
	if orderSize == 0 {
		return
	}

	// Find the largest pack size that can be used
	packSize := (*sortedPackSizes)[0]
	selectedPackSizeIndex := -1
	for i, size := range *sortedPackSizes {
		if orderSize < size {
			break
		}
		// set packSize to the largest size that fits in orderSize
		packSize = size
		selectedPackSizeIndex = i
	}

	// If leastFewPacks is true then we use nearest larger pack size to pack orderSize
	if leastFewPacks {
		orderSize = 0
		packs[(*sortedPackSizes)[selectedPackSizeIndex+1]]++

		return
	}

	// Devide the order size by the pack size to get the count of packs needed
	count := orderSize / packSize
	if count == 0 {
		orderSize = 0
		packs[packSize]++
	} else {
		orderSize -= count * packSize
		packs[packSize] += count
	}

	// Recursively call calculatePacks with the remaining order size
	s.calculatePacks(orderSize, packs, sortedPackSizes, leastFewPacks)
}
