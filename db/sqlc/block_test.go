package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomBlock(t *testing.T, town Town, clutter Clutter) Block {
	arg := CreateBlockParams{
		BlockName:       util.RandomName(8),
		BlockDesc:       util.RandomName(8),
		TotalPopulation: util.RandomInteger(1, 100),
		TownID:          town.ID,
		ClutterID:       clutter.ID,
	}
	block, err := testStore.CreateBlock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, block)
	require.Equal(t, arg.BlockName, block.BlockName)
	require.Equal(t, arg.BlockDesc, block.BlockDesc)
	require.Equal(t, arg.TotalPopulation, block.TotalPopulation)
	require.Equal(t, arg.TownID, block.TownID)
	require.Equal(t, arg.ClutterID, block.ClutterID)
	return block
}

func TestCreateBlock(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	createRandomBlock(t, town, clutter)
}

func TestGetBlock0(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block1 := createRandomBlock(t, town, clutter)
	block2, err := testStore.GetBlock0(context.Background(), block1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, block2)

	require.Equal(t, block1.ID, block2.ID)
	require.Equal(t, block1.BlockName, block2.BlockName)
	require.Equal(t, block1.BlockDesc, block2.BlockDesc)
	require.Equal(t, block1.TotalPopulation, block2.TotalPopulation)
	require.Equal(t, block1.TownID, block2.TownID)
	require.Equal(t, block1.ClutterID, block2.ClutterID)
}

func TestGetBlock1(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block1 := createRandomBlock(t, town, clutter)
	block2, err := testStore.GetBlock1(context.Background(), block1.BlockName)
	require.NoError(t, err)
	require.NotEmpty(t, block2)

	require.Equal(t, block1.ID, block2.ID)
	require.Equal(t, block1.BlockName, block2.BlockName)
	require.Equal(t, block1.BlockDesc, block2.BlockDesc)
	require.Equal(t, block1.TotalPopulation, block2.TotalPopulation)
	require.Equal(t, block1.TownID, block2.TownID)
	require.Equal(t, block1.ClutterID, block2.ClutterID)
}

func TestListBlocks(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	for i := 0; i < 10; i++ {
		createRandomBlock(t, town, clutter)

	}
	arg := ListBlocksParams{
		TownID:    town.ID,
		ClutterID: clutter.ID,
		Limit:     5,
		Offset:    5,
	}
	blocks, err := testStore.ListBlocks(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, blocks, 5)

	for _, block := range blocks {
		require.NotEmpty(t, block)
		require.True(t, block.TownID == town.ID || block.ClutterID == clutter.ID)
	}
}

func TestUpdateBlock(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block1 := createRandomBlock(t, town, clutter)
	arg := UpdateBlockParams{
		ID:              block1.ID,
		BlockName:       util.RandomName(8),
		BlockDesc:       util.RandomName(8),
		TotalPopulation: util.RandomInteger(1, 100),
		TownID:          town.ID,
		ClutterID:       clutter.ID,
	}
	block2, err := testStore.UpdateBlock(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, block2)

	require.Equal(t, block1.ID, block2.ID)
	require.Equal(t, arg.BlockName, block2.BlockName)
	require.Equal(t, arg.BlockDesc, block2.BlockDesc)
	require.Equal(t, arg.TotalPopulation, block2.TotalPopulation)
	require.Equal(t, block1.TownID, block2.TownID)
	require.Equal(t, block1.ClutterID, block2.ClutterID)

}

func TestDeleteBlock(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block1 := createRandomBlock(t, town, clutter)
	err := testStore.DeleteBlock(context.Background(), block1.ID)
	require.NoError(t, err)
	block2, err := testStore.GetBlock0(context.Background(), block1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, block2)

}
