package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomProperty(t *testing.T, block Block) Property {
	arg := CreatePropertyParams{
		PropertyName: util.RandomName(8),
		Lat:          util.RandomReal(1, 100),
		Long:         util.RandomReal(1, 100),
		BlockID:      block.ID,
	}
	property, err := testStore.CreateProperty(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, property)
	require.Equal(t, arg.PropertyName, property.PropertyName)
	require.Equal(t, arg.Lat, property.Lat)
	require.Equal(t, arg.Long, property.Long)
	require.Equal(t, arg.BlockID, property.BlockID)
	return property
}

func TestCreateProperty(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	createRandomProperty(t, block)
}

func TestGetProperty0(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	property1 := createRandomProperty(t, block)
	property2, err := testStore.GetProperty0(context.Background(), property1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, property2)

	require.Equal(t, property1.ID, property2.ID)
	require.Equal(t, property1.PropertyName, property2.PropertyName)
	require.Equal(t, property1.Lat, property2.Lat)
	require.Equal(t, property1.Long, property2.Long)
	require.Equal(t, property1.BlockID, property2.BlockID)
}

func TestGetProperty1(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	property1 := createRandomProperty(t, block)
	property2, err := testStore.GetProperty1(context.Background(), property1.PropertyName)
	require.NoError(t, err)
	require.NotEmpty(t, property2)

	require.Equal(t, property1.ID, property2.ID)
	require.Equal(t, property1.PropertyName, property2.PropertyName)
	require.Equal(t, property1.Lat, property2.Lat)
	require.Equal(t, property1.Long, property2.Long)
	require.Equal(t, property1.BlockID, property2.BlockID)
}

func TestListProperties(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	for i := 0; i < 10; i++ {
		createRandomProperty(t, block)

	}
	arg := ListPropertiesParams{
		BlockID: block.ID,
		Limit:   5,
		Offset:  5,
	}
	properties, err := testStore.ListProperties(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, properties, 5)

	for _, property := range properties {
		require.NotEmpty(t, property)
		require.True(t, arg.BlockID == property.BlockID)
	}
}

func TestUpdateProperty(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	property1 := createRandomProperty(t, block)
	arg := UpdatePropertyParams{
		ID:           property1.ID,
		PropertyName: util.RandomName(8),
		Lat:          util.RandomReal(1, 100),
		Long:         util.RandomReal(1, 100),
		BlockID:      block.ID,
	}
	property2, err := testStore.UpdateProperty(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, property2)

	require.Equal(t, property1.ID, property2.ID)
	require.Equal(t, arg.PropertyName, property2.PropertyName)
	require.Equal(t, arg.Lat, property2.Lat)
	require.Equal(t, arg.Long, property2.Long)
	require.Equal(t, property1.BlockID, property2.BlockID)

}

func TestDeleteProperty(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	property1 := createRandomProperty(t, block)
	err := testStore.DeleteProperty(context.Background(), property1.ID)
	require.NoError(t, err)
	property2, err := testStore.GetProperty0(context.Background(), property1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, property2)

}
