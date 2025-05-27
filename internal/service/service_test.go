package service

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alishercodecrafter/orderpackscalculator/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

//go:generate mockgen -destination=mock_repository.go -package=service github.com/alishercodecrafter/orderpackscalculator/internal/service PacksRepository

func TestPacksServiceImpl_GetPacks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockPacksRepository(ctrl)
	mockPacks := model.Packs{{PackSize: 250}, {PackSize: 500}}

	mockRepo.EXPECT().GetPacks().Return(mockPacks)

	service := NewPacksService(mockRepo)
	result := service.GetPacks()

	require.Equal(t, mockPacks, result)
}

func TestPacksServiceImpl_AddPack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockPacksRepository(ctrl)

	// Test with valid pack
	validPack := model.Pack{PackSize: 100}
	mockRepo.EXPECT().AddPack(validPack).Return(nil)

	service := NewPacksService(mockRepo)
	err := service.AddPack(validPack)

	require.NoError(t, err)

	// Test with invalid pack size
	invalidPack := model.Pack{PackSize: 0}

	// No expectations needed for this test as it shouldn't reach the repository
	service = NewPacksService(mockRepo)
	err = service.AddPack(invalidPack)

	require.Error(t, err)
	require.Contains(t, err.Error(), "greater than zero")

	// Test when repository returns error
	errorPack := model.Pack{PackSize: 200}
	mockRepo.EXPECT().AddPack(errorPack).Return(errors.New("repo error"))

	err = service.AddPack(errorPack)
	require.Error(t, err)
	require.Contains(t, err.Error(), "repo error")
}

func TestPacksServiceImpl_RemovePack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockPacksRepository(ctrl)

	// Test successful removal
	mockRepo.EXPECT().RemovePack(250).Return(nil)

	service := NewPacksService(mockRepo)
	err := service.RemovePack(250)

	require.NoError(t, err)

	// Test when repository returns error
	mockRepo.EXPECT().RemovePack(999).Return(errors.New("not found"))

	err = service.RemovePack(999)
	require.Error(t, err)
	require.Contains(t, err.Error(), "not found")
}

func TestPacksServiceImpl_CalculatePacks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockPacksRepository(ctrl)
	packs := model.Packs{
		{PackSize: 250},
		{PackSize: 500},
		{PackSize: 1000},
		{PackSize: 2000},
		{PackSize: 5000},
	}

	// Since this will be called for each test case, we use AnyTimes()
	mockRepo.EXPECT().GetPacks().Return(packs).AnyTimes()

	service := NewPacksService(mockRepo)

	testCases := []struct {
		orderSize     int
		expectedPacks map[int]int
	}{
		{
			orderSize:     1,
			expectedPacks: map[int]int{250: 1},
		},
		{
			orderSize:     250,
			expectedPacks: map[int]int{250: 1},
		},
		{
			orderSize:     251,
			expectedPacks: map[int]int{500: 1},
		},
		{
			orderSize:     501,
			expectedPacks: map[int]int{250: 1, 500: 1},
		},
		{
			orderSize:     12001,
			expectedPacks: map[int]int{250: 1, 2000: 1, 5000: 2},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Order size %d", tc.orderSize), func(t *testing.T) {
			result, err := service.CalculatePacks(tc.orderSize)
			require.NoError(t, err)

			require.Equal(t, tc.orderSize, result.OrderSize)
			require.Equal(t, tc.expectedPacks, result.Packs)
		})
	}
}

func TestPacksServiceImpl_CalculatePacksEdgeCases(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockPacksRepository(ctrl)
	packs := model.Packs{
		{PackSize: 10},
		{PackSize: 15},
		{PackSize: 20},
		{PackSize: 50},
		{PackSize: 100},
	}

	// Since this will be called for each test case, we use AnyTimes()
	mockRepo.EXPECT().GetPacks().Return(packs).AnyTimes()

	service := NewPacksService(mockRepo)

	testCases := []struct {
		orderSize       int
		expectedPacks   map[int]int
		isErrorExpected bool
	}{
		{
			orderSize:       -12,
			isErrorExpected: true,
		},
		{
			orderSize:       0,
			isErrorExpected: true,
		},
		{
			orderSize:     1,
			expectedPacks: map[int]int{10: 1},
		},
		{
			orderSize:     250,
			expectedPacks: map[int]int{100: 2, 50: 1},
		},
		{
			orderSize:     251,
			expectedPacks: map[int]int{100: 2, 50: 1, 10: 1},
		},
		{
			orderSize:     17,
			expectedPacks: map[int]int{20: 1},
		},
		{
			orderSize:     40,
			expectedPacks: map[int]int{20: 2},
		},
		{
			orderSize:     23,
			expectedPacks: map[int]int{20: 1, 10: 1},
		},
		{
			orderSize:     111,
			expectedPacks: map[int]int{100: 1, 15: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Order size %d", tc.orderSize), func(t *testing.T) {
			result, err := service.CalculatePacks(tc.orderSize)
			if tc.isErrorExpected {
				require.Error(t, err)

				return
			}
			require.NoError(t, err)

			require.Equal(t, tc.orderSize, result.OrderSize)
			require.Equal(t, tc.expectedPacks, result.Packs)
		})
	}
}
