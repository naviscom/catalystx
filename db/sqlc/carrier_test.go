package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomCarrier(t *testing.T, band Band) Carrier {
	arg := CreateCarrierParams{
		CarrierName: util.RandomName(8),
		CarrierDesc: util.RandomName(8),
		Size:        util.RandomInteger(1, 100),
		StartFreq:   util.RandomInteger(1, 100),
		EndFreq:     util.RandomInteger(1, 100),
		BandID:      band.ID,
	}
	carrier, err := testStore.CreateCarrier(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, carrier)
	require.Equal(t, arg.CarrierName, carrier.CarrierName)
	require.Equal(t, arg.CarrierDesc, carrier.CarrierDesc)
	require.Equal(t, arg.Size, carrier.Size)
	require.Equal(t, arg.StartFreq, carrier.StartFreq)
	require.Equal(t, arg.EndFreq, carrier.EndFreq)
	require.Equal(t, arg.BandID, carrier.BandID)
	return carrier
}

func TestCreateCarrier(t *testing.T) {
	tech := createRandomTech(t)
	band := createRandomBand(t, tech)
	createRandomCarrier(t, band)
}

func TestGetCarrier0(t *testing.T) {
	tech := createRandomTech(t)
	band := createRandomBand(t, tech)
	carrier1 := createRandomCarrier(t, band)
	carrier2, err := testStore.GetCarrier0(context.Background(), carrier1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, carrier2)

	require.Equal(t, carrier1.ID, carrier2.ID)
	require.Equal(t, carrier1.CarrierName, carrier2.CarrierName)
	require.Equal(t, carrier1.CarrierDesc, carrier2.CarrierDesc)
	require.Equal(t, carrier1.Size, carrier2.Size)
	require.Equal(t, carrier1.StartFreq, carrier2.StartFreq)
	require.Equal(t, carrier1.EndFreq, carrier2.EndFreq)
	require.Equal(t, carrier1.BandID, carrier2.BandID)
}

func TestGetCarrier1(t *testing.T) {
	tech := createRandomTech(t)
	band := createRandomBand(t, tech)
	carrier1 := createRandomCarrier(t, band)
	carrier2, err := testStore.GetCarrier1(context.Background(), carrier1.CarrierName)
	require.NoError(t, err)
	require.NotEmpty(t, carrier2)

	require.Equal(t, carrier1.ID, carrier2.ID)
	require.Equal(t, carrier1.CarrierName, carrier2.CarrierName)
	require.Equal(t, carrier1.CarrierDesc, carrier2.CarrierDesc)
	require.Equal(t, carrier1.Size, carrier2.Size)
	require.Equal(t, carrier1.StartFreq, carrier2.StartFreq)
	require.Equal(t, carrier1.EndFreq, carrier2.EndFreq)
	require.Equal(t, carrier1.BandID, carrier2.BandID)
}

func TestListCarriers(t *testing.T) {
	tech := createRandomTech(t)
	band := createRandomBand(t, tech)
	for i := 0; i < 10; i++ {
		createRandomCarrier(t, band)

	}
	arg := ListCarriersParams{
		BandID: band.ID,
		Limit:  5,
		Offset: 5,
	}
	carriers, err := testStore.ListCarriers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, carriers, 5)

	for _, carrier := range carriers {
		require.NotEmpty(t, carrier)
		require.True(t, arg.BandID == carrier.BandID)
	}
}

func TestUpdateCarrier(t *testing.T) {
	tech := createRandomTech(t)
	band := createRandomBand(t, tech)
	carrier1 := createRandomCarrier(t, band)
	arg := UpdateCarrierParams{
		ID:          carrier1.ID,
		CarrierName: util.RandomName(8),
		CarrierDesc: util.RandomName(8),
		Size:        util.RandomInteger(1, 100),
		StartFreq:   util.RandomInteger(1, 100),
		EndFreq:     util.RandomInteger(1, 100),
		BandID:      band.ID,
	}
	carrier2, err := testStore.UpdateCarrier(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, carrier2)

	require.Equal(t, carrier1.ID, carrier2.ID)
	require.Equal(t, arg.CarrierName, carrier2.CarrierName)
	require.Equal(t, arg.CarrierDesc, carrier2.CarrierDesc)
	require.Equal(t, arg.Size, carrier2.Size)
	require.Equal(t, arg.StartFreq, carrier2.StartFreq)
	require.Equal(t, arg.EndFreq, carrier2.EndFreq)
	require.Equal(t, carrier1.BandID, carrier2.BandID)

}

func TestDeleteCarrier(t *testing.T) {
	tech := createRandomTech(t)
	band := createRandomBand(t, tech)
	carrier1 := createRandomCarrier(t, band)
	err := testStore.DeleteCarrier(context.Background(), carrier1.ID)
	require.NoError(t, err)
	carrier2, err := testStore.GetCarrier0(context.Background(), carrier1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, carrier2)

}
