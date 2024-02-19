package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomCity(t *testing.T, state State) City {
	arg := CreateCityParams{
		CityName: util.RandomName(8),
		CityDesc: util.RandomName(8),
		StateID:  state.ID,
	}
	city, err := testStore.CreateCity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, city)
	require.Equal(t, arg.CityName, city.CityName)
	require.Equal(t, arg.CityDesc, city.CityDesc)
	require.Equal(t, arg.StateID, city.StateID)
	return city
}

func TestCreateCity(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	createRandomCity(t, state)
}

func TestGetCity0(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city1 := createRandomCity(t, state)
	city2, err := testStore.GetCity0(context.Background(), city1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, city2)

	require.Equal(t, city1.ID, city2.ID)
	require.Equal(t, city1.CityName, city2.CityName)
	require.Equal(t, city1.CityDesc, city2.CityDesc)
	require.Equal(t, city1.StateID, city2.StateID)
}

func TestGetCity1(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city1 := createRandomCity(t, state)
	city2, err := testStore.GetCity1(context.Background(), city1.CityName)
	require.NoError(t, err)
	require.NotEmpty(t, city2)

	require.Equal(t, city1.ID, city2.ID)
	require.Equal(t, city1.CityName, city2.CityName)
	require.Equal(t, city1.CityDesc, city2.CityDesc)
	require.Equal(t, city1.StateID, city2.StateID)
}

func TestListCities(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	for i := 0; i < 10; i++ {
		createRandomCity(t, state)

	}
	arg := ListCitiesParams{
		StateID: state.ID,
		Limit:   5,
		Offset:  5,
	}
	cities, err := testStore.ListCities(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, cities, 5)

	for _, city := range cities {
		require.NotEmpty(t, city)
		require.True(t, arg.StateID == city.StateID)
	}
}

func TestUpdateCity(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city1 := createRandomCity(t, state)
	arg := UpdateCityParams{
		ID:       city1.ID,
		CityName: util.RandomName(8),
		CityDesc: util.RandomName(8),
		StateID:  state.ID,
	}
	city2, err := testStore.UpdateCity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, city2)

	require.Equal(t, city1.ID, city2.ID)
	require.Equal(t, arg.CityName, city2.CityName)
	require.Equal(t, arg.CityDesc, city2.CityDesc)
	require.Equal(t, city1.StateID, city2.StateID)

}

func TestDeleteCity(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city1 := createRandomCity(t, state)
	err := testStore.DeleteCity(context.Background(), city1.ID)
	require.NoError(t, err)
	city2, err := testStore.GetCity0(context.Background(), city1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, city2)

}
