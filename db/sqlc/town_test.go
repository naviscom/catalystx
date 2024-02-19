package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomTown(t *testing.T, district District) Town {
	arg := CreateTownParams{
		TownName:   util.RandomName(8),
		TownDesc:   util.RandomName(8),
		DistrictID: district.ID,
	}
	town, err := testStore.CreateTown(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, town)
	require.Equal(t, arg.TownName, town.TownName)
	require.Equal(t, arg.TownDesc, town.TownDesc)
	require.Equal(t, arg.DistrictID, town.DistrictID)
	return town
}

func TestCreateTown(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	createRandomTown(t, district)
}

func TestGetTown0(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town1 := createRandomTown(t, district)
	town2, err := testStore.GetTown0(context.Background(), town1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, town2)

	require.Equal(t, town1.ID, town2.ID)
	require.Equal(t, town1.TownName, town2.TownName)
	require.Equal(t, town1.TownDesc, town2.TownDesc)
	require.Equal(t, town1.DistrictID, town2.DistrictID)
}

func TestGetTown1(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town1 := createRandomTown(t, district)
	town2, err := testStore.GetTown1(context.Background(), town1.TownName)
	require.NoError(t, err)
	require.NotEmpty(t, town2)

	require.Equal(t, town1.ID, town2.ID)
	require.Equal(t, town1.TownName, town2.TownName)
	require.Equal(t, town1.TownDesc, town2.TownDesc)
	require.Equal(t, town1.DistrictID, town2.DistrictID)
}

func TestListTowns(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	for i := 0; i < 10; i++ {
		createRandomTown(t, district)

	}
	arg := ListTownsParams{
		DistrictID: district.ID,
		Limit:      5,
		Offset:     5,
	}
	towns, err := testStore.ListTowns(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, towns, 5)

	for _, town := range towns {
		require.NotEmpty(t, town)
		require.True(t, arg.DistrictID == town.DistrictID)
	}
}

func TestUpdateTown(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town1 := createRandomTown(t, district)
	arg := UpdateTownParams{
		ID:         town1.ID,
		TownName:   util.RandomName(8),
		TownDesc:   util.RandomName(8),
		DistrictID: district.ID,
	}
	town2, err := testStore.UpdateTown(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, town2)

	require.Equal(t, town1.ID, town2.ID)
	require.Equal(t, arg.TownName, town2.TownName)
	require.Equal(t, arg.TownDesc, town2.TownDesc)
	require.Equal(t, town1.DistrictID, town2.DistrictID)

}

func TestDeleteTown(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town1 := createRandomTown(t, district)
	err := testStore.DeleteTown(context.Background(), town1.ID)
	require.NoError(t, err)
	town2, err := testStore.GetTown0(context.Background(), town1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, town2)

}
