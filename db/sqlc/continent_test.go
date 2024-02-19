package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomContinent(t *testing.T) Continent {
	arg := CreateContinentParams{
		ContinentName: util.RandomName(8),
		ContinentDesc: util.RandomName(8),
	}
	continent, err := testStore.CreateContinent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, continent)
	require.Equal(t, arg.ContinentName, continent.ContinentName)
	require.Equal(t, arg.ContinentDesc, continent.ContinentDesc)
	return continent
}

func TestCreateContinent(t *testing.T) {
	createRandomContinent(t)
}

func TestGetContinent0(t *testing.T) {
	continent1 := createRandomContinent(t)
	continent2, err := testStore.GetContinent0(context.Background(), continent1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, continent2)

	require.Equal(t, continent1.ID, continent2.ID)
	require.Equal(t, continent1.ContinentName, continent2.ContinentName)
	require.Equal(t, continent1.ContinentDesc, continent2.ContinentDesc)
}

func TestGetContinent1(t *testing.T) {
	continent1 := createRandomContinent(t)
	continent2, err := testStore.GetContinent1(context.Background(), continent1.ContinentName)
	require.NoError(t, err)
	require.NotEmpty(t, continent2)

	require.Equal(t, continent1.ID, continent2.ID)
	require.Equal(t, continent1.ContinentName, continent2.ContinentName)
	require.Equal(t, continent1.ContinentDesc, continent2.ContinentDesc)
}

func TestListContinents(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomContinent(t)

	}
	arg := ListContinentsParams{
		Limit:  5,
		Offset: 5,
	}
	continents, err := testStore.ListContinents(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, continents, 5)

	for _, continent := range continents {
		require.NotEmpty(t, continent)
	}
}

func TestUpdateContinent(t *testing.T) {
	continent1 := createRandomContinent(t)
	arg := UpdateContinentParams{
		ID:            continent1.ID,
		ContinentName: util.RandomName(8),
		ContinentDesc: util.RandomName(8),
	}
	continent2, err := testStore.UpdateContinent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, continent2)

	require.Equal(t, continent1.ID, continent2.ID)
	require.Equal(t, arg.ContinentName, continent2.ContinentName)
	require.Equal(t, arg.ContinentDesc, continent2.ContinentDesc)

}

func TestDeleteContinent(t *testing.T) {
	continent1 := createRandomContinent(t)
	err := testStore.DeleteContinent(context.Background(), continent1.ID)
	require.NoError(t, err)
	continent2, err := testStore.GetContinent0(context.Background(), continent1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, continent2)

}
