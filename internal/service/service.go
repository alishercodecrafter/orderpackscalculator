package service

import (
	"cmp"
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
	RemovePack(packSize model.PackSize) error
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
	if pack.Size <= 0 {
		return fmt.Errorf("pack size must be greater than zero")
	}

	return s.repo.AddPack(pack)
}

// RemovePackSize removes a pack size
func (s *PacksServiceImpl) RemovePack(packSize model.PackSize) error {
	return s.repo.RemovePack(packSize)
}

// CalculatePacks calculates the optimal number of packs needed for an order
func (s *PacksServiceImpl) CalculatePacks(orderSize int) (model.CalculationResponse, error) {
	packList := s.repo.GetPacks()
	// If no packList or invalid order size, return empty packsRule2
	if len(packList) == 0 {
		return model.CalculationResponse{}, fmt.Errorf("available packsRule2 list is empty")
	}

	if orderSize <= 0 {
		return model.CalculationResponse{}, fmt.Errorf("order size must be greater than zero")
	}

	// Sort packs list in descending order based on size
	slices.SortFunc(packList, func(a, b model.Pack) int {
		return cmp.Compare(b.Size, a.Size)
	})

	// calculate packs for the remaining order size
	results := calculatePacks(orderSize, orderSize, &packList, 0)

	return model.CalculationResponse{
		OrderSize: orderSize,
		Packs:     getTheBestCombinationOfPacks(results),
	}, nil
}

// getAmountOfItemsInPacks calculates the total amount of items in packs and the total count of packs
// Returns:
// - amount: the total amount of items in packs
// - totalCount: the total count of packs
func getAmountOfItemsInPacks(packs map[model.PackSize]int) (int, int) {
	amount := 0
	totalCount := 0
	for packSize, count := range packs {
		amount += int(packSize) * count
		totalCount += count
	}

	return amount, totalCount
}

// calculatePacks is a helper function to calculate the optimal number of packs needed for an order
// Parameters:
// - originalOrderSize: the original size of the order
// - orderSize: the current size of the order being processed
// - packsList: a pointer to the list of available packs sorted in descending order by pack size
// - startFrom: the index in packsList from which to start processing
func calculatePacks(
	originalOrderSize,
	orderSize int,
	packsList *model.Packs,
	startFrom int,
) []map[model.PackSize]int {
	// If orderSize is zero or less, return nil
	if orderSize <= 0 {
		return nil
	}

	// If startFrom index is out of bounds, return nil
	if startFrom >= len(*packsList) {
		return nil
	}

	// list of packs to return
	result := make([]map[model.PackSize]int, 0, len(*packsList))

	// on 1st iteration we pack whole order size for each available pack sizes
	if startFrom == 0 {
		for i := 0; i < len(*packsList); i++ {
			packSize := (*packsList)[i].Size

			count := originalOrderSize / int(packSize)
			remainOfDivision := originalOrderSize % int(packSize)
			if remainOfDivision != 0 {
				count++
			}

			m := make(map[model.PackSize]int)
			m[packSize] = count

			result = append(result, m)
		}
	}

	// on 2nd and further iterations we try remaining order in packs bigger than it
	for i := startFrom; i < len(*packsList); i++ {
		packSize := (*packsList)[i].Size

		if orderSize != originalOrderSize && orderSize <= int(packSize) {
			m := make(map[model.PackSize]int)
			m[packSize] = 1
			result = append(result, m)
		}
	}

	m := make(map[model.PackSize]int)
	currentPackSize := (*packsList)[startFrom].Size
	// if orderSize is larger than the current pack size, we can use it
	if orderSize > int(currentPackSize) {
		count := orderSize / int(currentPackSize)
		m[currentPackSize] += count
		remainOfDivision := orderSize % int(currentPackSize)
		// if there is no remainder, we can return the result because we have found the best combination
		if remainOfDivision == 0 {
			result = append(result, m)

			return result
		} else {
			orderSize = remainOfDivision
		}

		// if we have more packs to process,
		// we can try to find the best combination of packs for the remaining order size
		if startFrom < len(*packsList)-1 {
			innerMap := getTheBestCombinationOfPacks(calculatePacks(
				originalOrderSize,
				orderSize,
				packsList,
				startFrom+1,
			))
			for packSize, count := range innerMap {
				m[packSize] += count
			}
		} else {
			m[currentPackSize]++
		}

		result = append(result, m)
	} else {
		// go to next pack size
		result = append(result, getTheBestCombinationOfPacks(calculatePacks(
			originalOrderSize,
			originalOrderSize,
			packsList,
			startFrom+1,
		)))
	}

	return result
}

// getTheBestCombinationOfPacks finds the best combination of packs from a list of pack combinations
func getTheBestCombinationOfPacks(listOfPacks []map[model.PackSize]int) map[model.PackSize]int {
	if len(listOfPacks) == 0 {
		return map[model.PackSize]int{}
	}

	// exclude empty packs from the list
	normalListOfPacks := make([]map[model.PackSize]int, 0, len(listOfPacks))
	for _, packs := range listOfPacks {
		if len(packs) == 0 {
			continue
		}

		normalListOfPacks = append(normalListOfPacks, packs)
	}

	slices.SortFunc(normalListOfPacks, func(a, b map[model.PackSize]int) int {
		amountA, countA := getAmountOfItemsInPacks(a)
		amountB, countB := getAmountOfItemsInPacks(b)
		// Compare by total amount of items in packs first
		if amountA != amountB {
			return cmp.Compare(amountA, amountB)
		}
		// If amounts are equal, compare by the total count of packs
		return cmp.Compare(countA, countB)
	})

	if len(normalListOfPacks) == 0 {
		return map[model.PackSize]int{}
	}

	return normalListOfPacks[0]
}
