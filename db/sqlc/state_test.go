package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomState(t *testing.T, country Country, area Area) State {
	arg := CreateStateParams{
		StateName: util.RandomName(8),
		StateDesc: util.RandomName(8),
		CountryID: country.ID,
		AreaID:    area.ID,
	}
	state, err := testStore.CreateState(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, state)
	require.Equal(t, arg.StateName, state.StateName)
	require.Equal(t, arg.StateDesc, state.StateDesc)
	require.Equal(t, arg.CountryID, state.CountryID)
	require.Equal(t, arg.AreaID, state.AreaID)
	return state
}

func TestCreateState(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	createRandomState(t, country, area)
}

func TestGetState0(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state1 := createRandomState(t, country, area)
	state2, err := testStore.GetState0(context.Background(), state1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, state2)

	require.Equal(t, state1.ID, state2.ID)
	require.Equal(t, state1.StateName, state2.StateName)
	require.Equal(t, state1.StateDesc, state2.StateDesc)
	require.Equal(t, state1.CountryID, state2.CountryID)
	require.Equal(t, state1.AreaID, state2.AreaID)
}

func TestGetState1(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state1 := createRandomState(t, country, area)
	state2, err := testStore.GetState1(context.Background(), state1.StateName)
	require.NoError(t, err)
	require.NotEmpty(t, state2)

	require.Equal(t, state1.ID, state2.ID)
	require.Equal(t, state1.StateName, state2.StateName)
	require.Equal(t, state1.StateDesc, state2.StateDesc)
	require.Equal(t, state1.CountryID, state2.CountryID)
	require.Equal(t, state1.AreaID, state2.AreaID)
}

func TestListStates(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	for i := 0; i < 10; i++ {
		createRandomState(t, country, area)

	}
	arg := ListStatesParams{
		CountryID: country.ID,
		AreaID:    area.ID,
		Limit:     5,
		Offset:    5,
	}
	states, err := testStore.ListStates(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, states, 5)

	for _, state := range states {
		require.NotEmpty(t, state)
		require.True(t, state.CountryID == country.ID || state.AreaID == area.ID)
	}
}

func TestUpdateState(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state1 := createRandomState(t, country, area)
	arg := UpdateStateParams{
		ID:        state1.ID,
		StateName: util.RandomName(8),
		StateDesc: util.RandomName(8),
		CountryID: country.ID,
		AreaID:    area.ID,
	}
	state2, err := testStore.UpdateState(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, state2)

	require.Equal(t, state1.ID, state2.ID)
	require.Equal(t, arg.StateName, state2.StateName)
	require.Equal(t, arg.StateDesc, state2.StateDesc)
	require.Equal(t, state1.CountryID, state2.CountryID)
	require.Equal(t, state1.AreaID, state2.AreaID)

}

func TestDeleteState(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state1 := createRandomState(t, country, area)
	err := testStore.DeleteState(context.Background(), state1.ID)
	require.NoError(t, err)
	state2, err := testStore.GetState0(context.Background(), state1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, state2)

}
