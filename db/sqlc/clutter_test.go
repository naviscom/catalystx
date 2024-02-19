package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomClutter(t *testing.T) Clutter {
	arg := CreateClutterParams{
		ClutterName: util.RandomName(8),
		ClutterDesc: util.RandomName(8),
	}
	clutter, err := testStore.CreateClutter(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, clutter)
	require.Equal(t, arg.ClutterName, clutter.ClutterName)
	require.Equal(t, arg.ClutterDesc, clutter.ClutterDesc)
	return clutter
}

func TestCreateClutter(t *testing.T) {
	createRandomClutter(t)
}

func TestGetClutter0(t *testing.T) {
	clutter1 := createRandomClutter(t)
	clutter2, err := testStore.GetClutter0(context.Background(), clutter1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, clutter2)

	require.Equal(t, clutter1.ID, clutter2.ID)
	require.Equal(t, clutter1.ClutterName, clutter2.ClutterName)
	require.Equal(t, clutter1.ClutterDesc, clutter2.ClutterDesc)
}

func TestGetClutter1(t *testing.T) {
	clutter1 := createRandomClutter(t)
	clutter2, err := testStore.GetClutter1(context.Background(), clutter1.ClutterName)
	require.NoError(t, err)
	require.NotEmpty(t, clutter2)

	require.Equal(t, clutter1.ID, clutter2.ID)
	require.Equal(t, clutter1.ClutterName, clutter2.ClutterName)
	require.Equal(t, clutter1.ClutterDesc, clutter2.ClutterDesc)
}

func TestListClutters(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomClutter(t)

	}
	arg := ListCluttersParams{
		Limit:  5,
		Offset: 5,
	}
	clutters, err := testStore.ListClutters(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, clutters, 5)

	for _, clutter := range clutters {
		require.NotEmpty(t, clutter)
	}
}

func TestUpdateClutter(t *testing.T) {
	clutter1 := createRandomClutter(t)
	arg := UpdateClutterParams{
		ID:          clutter1.ID,
		ClutterName: util.RandomName(8),
		ClutterDesc: util.RandomName(8),
	}
	clutter2, err := testStore.UpdateClutter(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, clutter2)

	require.Equal(t, clutter1.ID, clutter2.ID)
	require.Equal(t, arg.ClutterName, clutter2.ClutterName)
	require.Equal(t, arg.ClutterDesc, clutter2.ClutterDesc)

}

func TestDeleteClutter(t *testing.T) {
	clutter1 := createRandomClutter(t)
	err := testStore.DeleteClutter(context.Background(), clutter1.ID)
	require.NoError(t, err)
	clutter2, err := testStore.GetClutter0(context.Background(), clutter1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, clutter2)

}
