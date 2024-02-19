package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomTech(t *testing.T) Tech {
	arg := CreateTechParams{
		TechName: util.RandomName(8),
		TechDesc: util.RandomName(8),
	}
	tech, err := testStore.CreateTech(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tech)
	require.Equal(t, arg.TechName, tech.TechName)
	require.Equal(t, arg.TechDesc, tech.TechDesc)
	return tech
}

func TestCreateTech(t *testing.T) {
	createRandomTech(t)
}

func TestGetTech0(t *testing.T) {
	tech1 := createRandomTech(t)
	tech2, err := testStore.GetTech0(context.Background(), tech1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tech2)

	require.Equal(t, tech1.ID, tech2.ID)
	require.Equal(t, tech1.TechName, tech2.TechName)
	require.Equal(t, tech1.TechDesc, tech2.TechDesc)
}

func TestGetTech1(t *testing.T) {
	tech1 := createRandomTech(t)
	tech2, err := testStore.GetTech1(context.Background(), tech1.TechName)
	require.NoError(t, err)
	require.NotEmpty(t, tech2)

	require.Equal(t, tech1.ID, tech2.ID)
	require.Equal(t, tech1.TechName, tech2.TechName)
	require.Equal(t, tech1.TechDesc, tech2.TechDesc)
}

func TestListTechs(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTech(t)

	}
	arg := ListTechsParams{
		Limit:  5,
		Offset: 5,
	}
	techs, err := testStore.ListTechs(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, techs, 5)

	for _, tech := range techs {
		require.NotEmpty(t, tech)
	}
}

func TestUpdateTech(t *testing.T) {
	tech1 := createRandomTech(t)
	arg := UpdateTechParams{
		ID:       tech1.ID,
		TechName: util.RandomName(8),
		TechDesc: util.RandomName(8),
	}
	tech2, err := testStore.UpdateTech(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tech2)

	require.Equal(t, tech1.ID, tech2.ID)
	require.Equal(t, arg.TechName, tech2.TechName)
	require.Equal(t, arg.TechDesc, tech2.TechDesc)

}

func TestDeleteTech(t *testing.T) {
	tech1 := createRandomTech(t)
	err := testStore.DeleteTech(context.Background(), tech1.ID)
	require.NoError(t, err)
	tech2, err := testStore.GetTech0(context.Background(), tech1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, tech2)

}
