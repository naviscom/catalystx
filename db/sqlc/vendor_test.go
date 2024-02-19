package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomVendor(t *testing.T) Vendor {
	arg := CreateVendorParams{
		VendorName: util.RandomName(8),
		VendorDesc: util.RandomName(8),
	}
	vendor, err := testStore.CreateVendor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, vendor)
	require.Equal(t, arg.VendorName, vendor.VendorName)
	require.Equal(t, arg.VendorDesc, vendor.VendorDesc)
	return vendor
}

func TestCreateVendor(t *testing.T) {
	createRandomVendor(t)
}

func TestGetVendor0(t *testing.T) {
	vendor1 := createRandomVendor(t)
	vendor2, err := testStore.GetVendor0(context.Background(), vendor1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, vendor2)

	require.Equal(t, vendor1.ID, vendor2.ID)
	require.Equal(t, vendor1.VendorName, vendor2.VendorName)
	require.Equal(t, vendor1.VendorDesc, vendor2.VendorDesc)
}

func TestGetVendor1(t *testing.T) {
	vendor1 := createRandomVendor(t)
	vendor2, err := testStore.GetVendor1(context.Background(), vendor1.VendorName)
	require.NoError(t, err)
	require.NotEmpty(t, vendor2)

	require.Equal(t, vendor1.ID, vendor2.ID)
	require.Equal(t, vendor1.VendorName, vendor2.VendorName)
	require.Equal(t, vendor1.VendorDesc, vendor2.VendorDesc)
}

func TestListVendors(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomVendor(t)

	}
	arg := ListVendorsParams{
		Limit:  5,
		Offset: 5,
	}
	vendors, err := testStore.ListVendors(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, vendors, 5)

	for _, vendor := range vendors {
		require.NotEmpty(t, vendor)
	}
}

func TestUpdateVendor(t *testing.T) {
	vendor1 := createRandomVendor(t)
	arg := UpdateVendorParams{
		ID:         vendor1.ID,
		VendorName: util.RandomName(8),
		VendorDesc: util.RandomName(8),
	}
	vendor2, err := testStore.UpdateVendor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, vendor2)

	require.Equal(t, vendor1.ID, vendor2.ID)
	require.Equal(t, arg.VendorName, vendor2.VendorName)
	require.Equal(t, arg.VendorDesc, vendor2.VendorDesc)

}

func TestDeleteVendor(t *testing.T) {
	vendor1 := createRandomVendor(t)
	err := testStore.DeleteVendor(context.Background(), vendor1.ID)
	require.NoError(t, err)
	vendor2, err := testStore.GetVendor0(context.Background(), vendor1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, vendor2)

}
