package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomServiceareatype(t *testing.T) Serviceareatype {
	arg := CreateServiceareatypeParams{
		ServiceareatypeName: util.RandomName(8),
		ServiceareatypeDesc: util.RandomName(8),
	}
	serviceareatype, err := testStore.CreateServiceareatype(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, serviceareatype)
	require.Equal(t, arg.ServiceareatypeName, serviceareatype.ServiceareatypeName)
	require.Equal(t, arg.ServiceareatypeDesc, serviceareatype.ServiceareatypeDesc)
	return serviceareatype
}

func TestCreateServiceareatype(t *testing.T) {
	createRandomServiceareatype(t)
}

func TestGetServiceareatype0(t *testing.T) {
	serviceareatype1 := createRandomServiceareatype(t)
	serviceareatype2, err := testStore.GetServiceareatype0(context.Background(), serviceareatype1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, serviceareatype2)

	require.Equal(t, serviceareatype1.ID, serviceareatype2.ID)
	require.Equal(t, serviceareatype1.ServiceareatypeName, serviceareatype2.ServiceareatypeName)
	require.Equal(t, serviceareatype1.ServiceareatypeDesc, serviceareatype2.ServiceareatypeDesc)
}

func TestGetServiceareatype1(t *testing.T) {
	serviceareatype1 := createRandomServiceareatype(t)
	serviceareatype2, err := testStore.GetServiceareatype1(context.Background(), serviceareatype1.ServiceareatypeName)
	require.NoError(t, err)
	require.NotEmpty(t, serviceareatype2)

	require.Equal(t, serviceareatype1.ID, serviceareatype2.ID)
	require.Equal(t, serviceareatype1.ServiceareatypeName, serviceareatype2.ServiceareatypeName)
	require.Equal(t, serviceareatype1.ServiceareatypeDesc, serviceareatype2.ServiceareatypeDesc)
}

func TestListServiceareatypes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomServiceareatype(t)

	}
	arg := ListServiceareatypesParams{
		Limit:  5,
		Offset: 5,
	}
	serviceareatypes, err := testStore.ListServiceareatypes(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, serviceareatypes, 5)

	for _, serviceareatype := range serviceareatypes {
		require.NotEmpty(t, serviceareatype)
	}
}

func TestUpdateServiceareatype(t *testing.T) {
	serviceareatype1 := createRandomServiceareatype(t)
	arg := UpdateServiceareatypeParams{
		ID:                  serviceareatype1.ID,
		ServiceareatypeName: util.RandomName(8),
		ServiceareatypeDesc: util.RandomName(8),
	}
	serviceareatype2, err := testStore.UpdateServiceareatype(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, serviceareatype2)

	require.Equal(t, serviceareatype1.ID, serviceareatype2.ID)
	require.Equal(t, arg.ServiceareatypeName, serviceareatype2.ServiceareatypeName)
	require.Equal(t, arg.ServiceareatypeDesc, serviceareatype2.ServiceareatypeDesc)

}

func TestDeleteServiceareatype(t *testing.T) {
	serviceareatype1 := createRandomServiceareatype(t)
	err := testStore.DeleteServiceareatype(context.Background(), serviceareatype1.ID)
	require.NoError(t, err)
	serviceareatype2, err := testStore.GetServiceareatype0(context.Background(), serviceareatype1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, serviceareatype2)

}
