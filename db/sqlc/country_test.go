package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomCountry(t *testing.T, continent Continent) Country {
	arg := CreateCountryParams{
		CountryName: util.RandomName(8),
		CountryDesc: util.RandomName(8),
		ContinentID: continent.ID,
	}
	country, err := testStore.CreateCountry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, country)
	require.Equal(t, arg.CountryName, country.CountryName)
	require.Equal(t, arg.CountryDesc, country.CountryDesc)
	require.Equal(t, arg.ContinentID, country.ContinentID)
	return country
}

func TestCreateCountry(t *testing.T) {
	continent := createRandomContinent(t)
	createRandomCountry(t, continent)
}

func TestGetCountry0(t *testing.T) {
	continent := createRandomContinent(t)
	country1 := createRandomCountry(t, continent)
	country2, err := testStore.GetCountry0(context.Background(), country1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, country2)

	require.Equal(t, country1.ID, country2.ID)
	require.Equal(t, country1.CountryName, country2.CountryName)
	require.Equal(t, country1.CountryDesc, country2.CountryDesc)
	require.Equal(t, country1.ContinentID, country2.ContinentID)
}

func TestGetCountry1(t *testing.T) {
	continent := createRandomContinent(t)
	country1 := createRandomCountry(t, continent)
	country2, err := testStore.GetCountry1(context.Background(), country1.CountryName)
	require.NoError(t, err)
	require.NotEmpty(t, country2)

	require.Equal(t, country1.ID, country2.ID)
	require.Equal(t, country1.CountryName, country2.CountryName)
	require.Equal(t, country1.CountryDesc, country2.CountryDesc)
	require.Equal(t, country1.ContinentID, country2.ContinentID)
}

func TestListCountries(t *testing.T) {
	continent := createRandomContinent(t)
	for i := 0; i < 10; i++ {
		createRandomCountry(t, continent)

	}
	arg := ListCountriesParams{
		ContinentID: continent.ID,
		Limit:       5,
		Offset:      5,
	}
	countries, err := testStore.ListCountries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, countries, 5)

	for _, country := range countries {
		require.NotEmpty(t, country)
		require.True(t, arg.ContinentID == country.ContinentID)
	}
}

func TestUpdateCountry(t *testing.T) {
	continent := createRandomContinent(t)
	country1 := createRandomCountry(t, continent)
	arg := UpdateCountryParams{
		ID:          country1.ID,
		CountryName: util.RandomName(8),
		CountryDesc: util.RandomName(8),
		ContinentID: continent.ID,
	}
	country2, err := testStore.UpdateCountry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, country2)

	require.Equal(t, country1.ID, country2.ID)
	require.Equal(t, arg.CountryName, country2.CountryName)
	require.Equal(t, arg.CountryDesc, country2.CountryDesc)
	require.Equal(t, country1.ContinentID, country2.ContinentID)

}

func TestDeleteCountry(t *testing.T) {
	continent := createRandomContinent(t)
	country1 := createRandomCountry(t, continent)
	err := testStore.DeleteCountry(context.Background(), country1.ID)
	require.NoError(t, err)
	country2, err := testStore.GetCountry0(context.Background(), country1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, country2)

}
