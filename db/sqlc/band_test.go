package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomBand(t *testing.T, tech Tech) Band {
	arg := CreateBandParams{
		BandName:  util.RandomName(8),
		BandDesc:  util.RandomName(8),
		Size:      util.RandomInteger(1, 100),
		StartFreq: util.RandomInteger(1, 100),
		EndFreq:   util.RandomInteger(1, 100),
		TechID:    tech.ID,
	}
	band, err := testStore.CreateBand(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, band)
	require.Equal(t, arg.BandName, band.BandName)
	require.Equal(t, arg.BandDesc, band.BandDesc)
	require.Equal(t, arg.Size, band.Size)
	require.Equal(t, arg.StartFreq, band.StartFreq)
	require.Equal(t, arg.EndFreq, band.EndFreq)
	require.Equal(t, arg.TechID, band.TechID)
	return band
}

func TestCreateBand(t *testing.T) {
	tech := createRandomTech(t)
	createRandomBand(t, tech)
}

func TestGetBand0(t *testing.T) {
	tech := createRandomTech(t)
	band1 := createRandomBand(t, tech)
	band2, err := testStore.GetBand0(context.Background(), band1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, band2)

	require.Equal(t, band1.ID, band2.ID)
	require.Equal(t, band1.BandName, band2.BandName)
	require.Equal(t, band1.BandDesc, band2.BandDesc)
	require.Equal(t, band1.Size, band2.Size)
	require.Equal(t, band1.StartFreq, band2.StartFreq)
	require.Equal(t, band1.EndFreq, band2.EndFreq)
	require.Equal(t, band1.TechID, band2.TechID)
}

func TestGetBand1(t *testing.T) {
	tech := createRandomTech(t)
	band1 := createRandomBand(t, tech)
	band2, err := testStore.GetBand1(context.Background(), band1.BandName)
	require.NoError(t, err)
	require.NotEmpty(t, band2)

	require.Equal(t, band1.ID, band2.ID)
	require.Equal(t, band1.BandName, band2.BandName)
	require.Equal(t, band1.BandDesc, band2.BandDesc)
	require.Equal(t, band1.Size, band2.Size)
	require.Equal(t, band1.StartFreq, band2.StartFreq)
	require.Equal(t, band1.EndFreq, band2.EndFreq)
	require.Equal(t, band1.TechID, band2.TechID)
}

func TestListBands(t *testing.T) {
	tech := createRandomTech(t)
	for i := 0; i < 10; i++ {
		createRandomBand(t, tech)

	}
	arg := ListBandsParams{
		TechID: tech.ID,
		Limit:  5,
		Offset: 5,
	}
	bands, err := testStore.ListBands(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, bands, 5)

	for _, band := range bands {
		require.NotEmpty(t, band)
		require.True(t, arg.TechID == band.TechID)
	}
}

func TestUpdateBand(t *testing.T) {
	tech := createRandomTech(t)
	band1 := createRandomBand(t, tech)
	arg := UpdateBandParams{
		ID:        band1.ID,
		BandName:  util.RandomName(8),
		BandDesc:  util.RandomName(8),
		Size:      util.RandomInteger(1, 100),
		StartFreq: util.RandomInteger(1, 100),
		EndFreq:   util.RandomInteger(1, 100),
		TechID:    tech.ID,
	}
	band2, err := testStore.UpdateBand(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, band2)

	require.Equal(t, band1.ID, band2.ID)
	require.Equal(t, arg.BandName, band2.BandName)
	require.Equal(t, arg.BandDesc, band2.BandDesc)
	require.Equal(t, arg.Size, band2.Size)
	require.Equal(t, arg.StartFreq, band2.StartFreq)
	require.Equal(t, arg.EndFreq, band2.EndFreq)
	require.Equal(t, band1.TechID, band2.TechID)

}

func TestDeleteBand(t *testing.T) {
	tech := createRandomTech(t)
	band1 := createRandomBand(t, tech)
	err := testStore.DeleteBand(context.Background(), band1.ID)
	require.NoError(t, err)
	band2, err := testStore.GetBand0(context.Background(), band1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, band2)

}
