package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomArea(t *testing.T) Area {
	arg := CreateAreaParams{
		AreaName: util.RandomName(8),
		AreaDesc: util.RandomName(8),
	}
	area, err := testStore.CreateArea(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, area)
	require.Equal(t, arg.AreaName, area.AreaName)
	require.Equal(t, arg.AreaDesc, area.AreaDesc)
	return area
}

func TestCreateArea(t *testing.T) {
	createRandomArea(t)
}

func TestGetArea0(t *testing.T) {
	area1 := createRandomArea(t)
	area2, err := testStore.GetArea0(context.Background(), area1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, area2)

	require.Equal(t, area1.ID, area2.ID)
	require.Equal(t, area1.AreaName, area2.AreaName)
	require.Equal(t, area1.AreaDesc, area2.AreaDesc)
}

func TestGetArea1(t *testing.T) {
	area1 := createRandomArea(t)
	area2, err := testStore.GetArea1(context.Background(), area1.AreaName)
	require.NoError(t, err)
	require.NotEmpty(t, area2)

	require.Equal(t, area1.ID, area2.ID)
	require.Equal(t, area1.AreaName, area2.AreaName)
	require.Equal(t, area1.AreaDesc, area2.AreaDesc)
}

func TestListAreas(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomArea(t)

	}
	arg := ListAreasParams{
		Limit:  5,
		Offset: 5,
	}
	areas, err := testStore.ListAreas(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, areas, 5)

	for _, area := range areas {
		require.NotEmpty(t, area)
	}
}

func TestUpdateArea(t *testing.T) {
	area1 := createRandomArea(t)
	arg := UpdateAreaParams{
		ID:       area1.ID,
		AreaName: util.RandomName(8),
		AreaDesc: util.RandomName(8),
	}
	area2, err := testStore.UpdateArea(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, area2)

	require.Equal(t, area1.ID, area2.ID)
	require.Equal(t, arg.AreaName, area2.AreaName)
	require.Equal(t, arg.AreaDesc, area2.AreaDesc)

}

func TestDeleteArea(t *testing.T) {
	area1 := createRandomArea(t)
	err := testStore.DeleteArea(context.Background(), area1.ID)
	require.NoError(t, err)
	area2, err := testStore.GetArea0(context.Background(), area1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, area2)

}
