package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomDistrict(t *testing.T, city City) District {
	arg := CreateDistrictParams{
		DistrictName: util.RandomName(8),
		DistrictDesc: util.RandomName(8),
		CityID:       city.ID,
	}
	district, err := testStore.CreateDistrict(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, district)
	require.Equal(t, arg.DistrictName, district.DistrictName)
	require.Equal(t, arg.DistrictDesc, district.DistrictDesc)
	require.Equal(t, arg.CityID, district.CityID)
	return district
}

func TestCreateDistrict(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	createRandomDistrict(t, city)
}

func TestGetDistrict0(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district1 := createRandomDistrict(t, city)
	district2, err := testStore.GetDistrict0(context.Background(), district1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, district2)

	require.Equal(t, district1.ID, district2.ID)
	require.Equal(t, district1.DistrictName, district2.DistrictName)
	require.Equal(t, district1.DistrictDesc, district2.DistrictDesc)
	require.Equal(t, district1.CityID, district2.CityID)
}

func TestGetDistrict1(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district1 := createRandomDistrict(t, city)
	district2, err := testStore.GetDistrict1(context.Background(), district1.DistrictName)
	require.NoError(t, err)
	require.NotEmpty(t, district2)

	require.Equal(t, district1.ID, district2.ID)
	require.Equal(t, district1.DistrictName, district2.DistrictName)
	require.Equal(t, district1.DistrictDesc, district2.DistrictDesc)
	require.Equal(t, district1.CityID, district2.CityID)
}

func TestListDistricts(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	for i := 0; i < 10; i++ {
		createRandomDistrict(t, city)

	}
	arg := ListDistrictsParams{
		CityID: city.ID,
		Limit:  5,
		Offset: 5,
	}
	districts, err := testStore.ListDistricts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, districts, 5)

	for _, district := range districts {
		require.NotEmpty(t, district)
		require.True(t, arg.CityID == district.CityID)
	}
}

func TestUpdateDistrict(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district1 := createRandomDistrict(t, city)
	arg := UpdateDistrictParams{
		ID:           district1.ID,
		DistrictName: util.RandomName(8),
		DistrictDesc: util.RandomName(8),
		CityID:       city.ID,
	}
	district2, err := testStore.UpdateDistrict(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, district2)

	require.Equal(t, district1.ID, district2.ID)
	require.Equal(t, arg.DistrictName, district2.DistrictName)
	require.Equal(t, arg.DistrictDesc, district2.DistrictDesc)
	require.Equal(t, district1.CityID, district2.CityID)

}

func TestDeleteDistrict(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district1 := createRandomDistrict(t, city)
	err := testStore.DeleteDistrict(context.Background(), district1.ID)
	require.NoError(t, err)
	district2, err := testStore.GetDistrict0(context.Background(), district1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, district2)

}
