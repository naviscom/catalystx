package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomSitetype(t *testing.T) Sitetype {
	arg := CreateSitetypeParams{
		TypeName: util.RandomName(8),
		TypeDesc: util.RandomName(8),
	}
	sitetype, err := testStore.CreateSitetype(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sitetype)
	require.Equal(t, arg.TypeName, sitetype.TypeName)
	require.Equal(t, arg.TypeDesc, sitetype.TypeDesc)
	return sitetype
}

func TestCreateSitetype(t *testing.T) {
	createRandomSitetype(t)
}

func TestGetSitetype0(t *testing.T) {
	sitetype1 := createRandomSitetype(t)
	sitetype2, err := testStore.GetSitetype0(context.Background(), sitetype1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, sitetype2)

	require.Equal(t, sitetype1.ID, sitetype2.ID)
	require.Equal(t, sitetype1.TypeName, sitetype2.TypeName)
	require.Equal(t, sitetype1.TypeDesc, sitetype2.TypeDesc)
}

func TestGetSitetype1(t *testing.T) {
	sitetype1 := createRandomSitetype(t)
	sitetype2, err := testStore.GetSitetype1(context.Background(), sitetype1.TypeName)
	require.NoError(t, err)
	require.NotEmpty(t, sitetype2)

	require.Equal(t, sitetype1.ID, sitetype2.ID)
	require.Equal(t, sitetype1.TypeName, sitetype2.TypeName)
	require.Equal(t, sitetype1.TypeDesc, sitetype2.TypeDesc)
}

func TestListSitetypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomSitetype(t)

	}
	arg := ListSitetypesParams{
		Limit:  5,
		Offset: 5,
	}
	sitetypes, err := testStore.ListSitetypes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, sitetypes, 5)

	for _, sitetype := range sitetypes {
		require.NotEmpty(t, sitetype)
	}
}

func TestUpdateSitetype(t *testing.T) {
	sitetype1 := createRandomSitetype(t)
	arg := UpdateSitetypeParams{
		ID:       sitetype1.ID,
		TypeName: util.RandomName(8),
		TypeDesc: util.RandomName(8),
	}
	sitetype2, err := testStore.UpdateSitetype(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sitetype2)

	require.Equal(t, sitetype1.ID, sitetype2.ID)
	require.Equal(t, arg.TypeName, sitetype2.TypeName)
	require.Equal(t, arg.TypeDesc, sitetype2.TypeDesc)

}

func TestDeleteSitetype(t *testing.T) {
	sitetype1 := createRandomSitetype(t)
	err := testStore.DeleteSitetype(context.Background(), sitetype1.ID)
	require.NoError(t, err)
	sitetype2, err := testStore.GetSitetype0(context.Background(), sitetype1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, sitetype2)

}
