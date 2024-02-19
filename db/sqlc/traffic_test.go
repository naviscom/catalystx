package db

import (
	"context"
	"testing"
	"time"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomTraffic(t *testing.T, cell Cell) Traffic {
	arg := CreateTrafficParams{
		TrafficDate:              time.Now().UTC(),
		Avgdailydldatamb:         util.RandomReal(1, 100),
		Avgdailyuldatamb:         util.RandomReal(1, 100),
		Avgdailytotdatamb:        util.RandomReal(1, 100),
		Avgdailytotvoicemin:      util.RandomReal(1, 100),
		Avgdailytotvideomin:      util.RandomReal(1, 100),
		Qci1Data:                 util.RandomReal(1, 100),
		Qci6Data:                 util.RandomReal(1, 100),
		Qci8Data:                 util.RandomReal(1, 100),
		QciOtherData:             util.RandomReal(1, 100),
		Avgdailytotvoicemin4g:    util.RandomReal(1, 100),
		Avgdailytotvoicemintotal: util.RandomReal(1, 100),
		Userdlthroughput:         util.RandomReal(1, 100),
		Dlpacketlossrate:         util.RandomReal(1, 100),
		Overallpsdropcallrate:    util.RandomReal(1, 100),
		Bhdldatamb:               util.RandomReal(1, 100),
		Bhupdatamb:               util.RandomReal(1, 100),
		Bhtotdatamb:              util.RandomReal(1, 100),
		Bhtotvoicemin:            util.RandomReal(1, 100),
		Bhtotvideomin:            util.RandomReal(1, 100),
		Bhcsusers:                util.RandomReal(1, 100),
		Bhhsupausers:             util.RandomReal(1, 100),
		Bhhsdpausers:             util.RandomReal(1, 100),
		Bhr99uldl:                util.RandomReal(1, 100),
		Powercapacity:            util.RandomReal(1, 100),
		Powerutilization:         util.RandomReal(1, 100),
		Codecapacity:             util.RandomReal(1, 100),
		Codeutilization:          util.RandomReal(1, 100),
		Ceulcapacity:             util.RandomReal(1, 100),
		Ceulutilization:          util.RandomReal(1, 100),
		Cedlcapacity:             util.RandomReal(1, 100),
		Cedlutilization:          util.RandomReal(1, 100),
		Iubcapacity:              util.RandomReal(1, 100),
		Iubutlization:            util.RandomReal(1, 100),
		Bhrrcusers:               util.RandomReal(1, 100),
		CellID:                   cell.ID,
	}
	traffic, err := testStore.CreateTraffic(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, traffic)
	require.WithinDuration(t, arg.TrafficDate, traffic.TrafficDate, time.Second)
	require.Equal(t, arg.Avgdailydldatamb, traffic.Avgdailydldatamb)
	require.Equal(t, arg.Avgdailyuldatamb, traffic.Avgdailyuldatamb)
	require.Equal(t, arg.Avgdailytotdatamb, traffic.Avgdailytotdatamb)
	require.Equal(t, arg.Avgdailytotvoicemin, traffic.Avgdailytotvoicemin)
	require.Equal(t, arg.Avgdailytotvideomin, traffic.Avgdailytotvideomin)
	require.Equal(t, arg.Qci1Data, traffic.Qci1Data)
	require.Equal(t, arg.Qci6Data, traffic.Qci6Data)
	require.Equal(t, arg.Qci8Data, traffic.Qci8Data)
	require.Equal(t, arg.QciOtherData, traffic.QciOtherData)
	require.Equal(t, arg.Avgdailytotvoicemin4g, traffic.Avgdailytotvoicemin4g)
	require.Equal(t, arg.Avgdailytotvoicemintotal, traffic.Avgdailytotvoicemintotal)
	require.Equal(t, arg.Userdlthroughput, traffic.Userdlthroughput)
	require.Equal(t, arg.Dlpacketlossrate, traffic.Dlpacketlossrate)
	require.Equal(t, arg.Overallpsdropcallrate, traffic.Overallpsdropcallrate)
	require.Equal(t, arg.Bhdldatamb, traffic.Bhdldatamb)
	require.Equal(t, arg.Bhupdatamb, traffic.Bhupdatamb)
	require.Equal(t, arg.Bhtotdatamb, traffic.Bhtotdatamb)
	require.Equal(t, arg.Bhtotvoicemin, traffic.Bhtotvoicemin)
	require.Equal(t, arg.Bhtotvideomin, traffic.Bhtotvideomin)
	require.Equal(t, arg.Bhcsusers, traffic.Bhcsusers)
	require.Equal(t, arg.Bhhsupausers, traffic.Bhhsupausers)
	require.Equal(t, arg.Bhhsdpausers, traffic.Bhhsdpausers)
	require.Equal(t, arg.Bhr99uldl, traffic.Bhr99uldl)
	require.Equal(t, arg.Powercapacity, traffic.Powercapacity)
	require.Equal(t, arg.Powerutilization, traffic.Powerutilization)
	require.Equal(t, arg.Codecapacity, traffic.Codecapacity)
	require.Equal(t, arg.Codeutilization, traffic.Codeutilization)
	require.Equal(t, arg.Ceulcapacity, traffic.Ceulcapacity)
	require.Equal(t, arg.Ceulutilization, traffic.Ceulutilization)
	require.Equal(t, arg.Cedlcapacity, traffic.Cedlcapacity)
	require.Equal(t, arg.Cedlutilization, traffic.Cedlutilization)
	require.Equal(t, arg.Iubcapacity, traffic.Iubcapacity)
	require.Equal(t, arg.Iubutlization, traffic.Iubutlization)
	require.Equal(t, arg.Bhrrcusers, traffic.Bhrrcusers)
	require.Equal(t, arg.CellID, traffic.CellID)
	return traffic
}

func TestCreateTraffic(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	tech := createRandomTech(t)
	block := createRandomBlock(t, town, clutter)
	band := createRandomBand(t, tech)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	site := createRandomSite(t, property, sitetype, vendor)
	carrier := createRandomCarrier(t, band)
	serviceareatype := createRandomServiceareatype(t)
	cell := createRandomCell(t, site, carrier, serviceareatype)
	createRandomTraffic(t, cell)
}

func TestGetTraffic0(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	tech := createRandomTech(t)
	block := createRandomBlock(t, town, clutter)
	band := createRandomBand(t, tech)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	site := createRandomSite(t, property, sitetype, vendor)
	carrier := createRandomCarrier(t, band)
	serviceareatype := createRandomServiceareatype(t)
	cell := createRandomCell(t, site, carrier, serviceareatype)
	traffic1 := createRandomTraffic(t, cell)
	traffic2, err := testStore.GetTraffic0(context.Background(), traffic1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, traffic2)

	require.Equal(t, traffic1.ID, traffic2.ID)
	require.WithinDuration(t, traffic1.TrafficDate, traffic2.TrafficDate, time.Second)
	require.Equal(t, traffic1.Avgdailydldatamb, traffic2.Avgdailydldatamb)
	require.Equal(t, traffic1.Avgdailyuldatamb, traffic2.Avgdailyuldatamb)
	require.Equal(t, traffic1.Avgdailytotdatamb, traffic2.Avgdailytotdatamb)
	require.Equal(t, traffic1.Avgdailytotvoicemin, traffic2.Avgdailytotvoicemin)
	require.Equal(t, traffic1.Avgdailytotvideomin, traffic2.Avgdailytotvideomin)
	require.Equal(t, traffic1.Qci1Data, traffic2.Qci1Data)
	require.Equal(t, traffic1.Qci6Data, traffic2.Qci6Data)
	require.Equal(t, traffic1.Qci8Data, traffic2.Qci8Data)
	require.Equal(t, traffic1.QciOtherData, traffic2.QciOtherData)
	require.Equal(t, traffic1.Avgdailytotvoicemin4g, traffic2.Avgdailytotvoicemin4g)
	require.Equal(t, traffic1.Avgdailytotvoicemintotal, traffic2.Avgdailytotvoicemintotal)
	require.Equal(t, traffic1.Userdlthroughput, traffic2.Userdlthroughput)
	require.Equal(t, traffic1.Dlpacketlossrate, traffic2.Dlpacketlossrate)
	require.Equal(t, traffic1.Overallpsdropcallrate, traffic2.Overallpsdropcallrate)
	require.Equal(t, traffic1.Bhdldatamb, traffic2.Bhdldatamb)
	require.Equal(t, traffic1.Bhupdatamb, traffic2.Bhupdatamb)
	require.Equal(t, traffic1.Bhtotdatamb, traffic2.Bhtotdatamb)
	require.Equal(t, traffic1.Bhtotvoicemin, traffic2.Bhtotvoicemin)
	require.Equal(t, traffic1.Bhtotvideomin, traffic2.Bhtotvideomin)
	require.Equal(t, traffic1.Bhcsusers, traffic2.Bhcsusers)
	require.Equal(t, traffic1.Bhhsupausers, traffic2.Bhhsupausers)
	require.Equal(t, traffic1.Bhhsdpausers, traffic2.Bhhsdpausers)
	require.Equal(t, traffic1.Bhr99uldl, traffic2.Bhr99uldl)
	require.Equal(t, traffic1.Powercapacity, traffic2.Powercapacity)
	require.Equal(t, traffic1.Powerutilization, traffic2.Powerutilization)
	require.Equal(t, traffic1.Codecapacity, traffic2.Codecapacity)
	require.Equal(t, traffic1.Codeutilization, traffic2.Codeutilization)
	require.Equal(t, traffic1.Ceulcapacity, traffic2.Ceulcapacity)
	require.Equal(t, traffic1.Ceulutilization, traffic2.Ceulutilization)
	require.Equal(t, traffic1.Cedlcapacity, traffic2.Cedlcapacity)
	require.Equal(t, traffic1.Cedlutilization, traffic2.Cedlutilization)
	require.Equal(t, traffic1.Iubcapacity, traffic2.Iubcapacity)
	require.Equal(t, traffic1.Iubutlization, traffic2.Iubutlization)
	require.Equal(t, traffic1.Bhrrcusers, traffic2.Bhrrcusers)
	require.Equal(t, traffic1.CellID, traffic2.CellID)
}

func TestListTraffic(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	tech := createRandomTech(t)
	block := createRandomBlock(t, town, clutter)
	band := createRandomBand(t, tech)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	site := createRandomSite(t, property, sitetype, vendor)
	carrier := createRandomCarrier(t, band)
	serviceareatype := createRandomServiceareatype(t)
	cell := createRandomCell(t, site, carrier, serviceareatype)
	for i := 0; i < 10; i++ {
		createRandomTraffic(t, cell)

	}
	arg := ListTrafficParams{
		CellID: cell.ID,
		Limit:  5,
		Offset: 5,
	}
	traffic, err := testStore.ListTraffic(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, traffic, 5)

	for _, traffic := range traffic {
		require.NotEmpty(t, traffic)
		require.True(t, arg.CellID == traffic.CellID)
	}
}

func TestUpdateTraffic(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	tech := createRandomTech(t)
	block := createRandomBlock(t, town, clutter)
	band := createRandomBand(t, tech)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	site := createRandomSite(t, property, sitetype, vendor)
	carrier := createRandomCarrier(t, band)
	serviceareatype := createRandomServiceareatype(t)
	cell := createRandomCell(t, site, carrier, serviceareatype)
	traffic1 := createRandomTraffic(t, cell)
	arg := UpdateTrafficParams{
		ID:                       traffic1.ID,
		TrafficDate:              time.Now().UTC(),
		Avgdailydldatamb:         util.RandomReal(1, 100),
		Avgdailyuldatamb:         util.RandomReal(1, 100),
		Avgdailytotdatamb:        util.RandomReal(1, 100),
		Avgdailytotvoicemin:      util.RandomReal(1, 100),
		Avgdailytotvideomin:      util.RandomReal(1, 100),
		Qci1Data:                 util.RandomReal(1, 100),
		Qci6Data:                 util.RandomReal(1, 100),
		Qci8Data:                 util.RandomReal(1, 100),
		QciOtherData:             util.RandomReal(1, 100),
		Avgdailytotvoicemin4g:    util.RandomReal(1, 100),
		Avgdailytotvoicemintotal: util.RandomReal(1, 100),
		Userdlthroughput:         util.RandomReal(1, 100),
		Dlpacketlossrate:         util.RandomReal(1, 100),
		Overallpsdropcallrate:    util.RandomReal(1, 100),
		Bhdldatamb:               util.RandomReal(1, 100),
		Bhupdatamb:               util.RandomReal(1, 100),
		Bhtotdatamb:              util.RandomReal(1, 100),
		Bhtotvoicemin:            util.RandomReal(1, 100),
		Bhtotvideomin:            util.RandomReal(1, 100),
		Bhcsusers:                util.RandomReal(1, 100),
		Bhhsupausers:             util.RandomReal(1, 100),
		Bhhsdpausers:             util.RandomReal(1, 100),
		Bhr99uldl:                util.RandomReal(1, 100),
		Powercapacity:            util.RandomReal(1, 100),
		Powerutilization:         util.RandomReal(1, 100),
		Codecapacity:             util.RandomReal(1, 100),
		Codeutilization:          util.RandomReal(1, 100),
		Ceulcapacity:             util.RandomReal(1, 100),
		Ceulutilization:          util.RandomReal(1, 100),
		Cedlcapacity:             util.RandomReal(1, 100),
		Cedlutilization:          util.RandomReal(1, 100),
		Iubcapacity:              util.RandomReal(1, 100),
		Iubutlization:            util.RandomReal(1, 100),
		Bhrrcusers:               util.RandomReal(1, 100),
		CellID:                   cell.ID,
	}
	traffic2, err := testStore.UpdateTraffic(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, traffic2)

	require.Equal(t, traffic1.ID, traffic2.ID)
	require.WithinDuration(t, arg.TrafficDate, traffic2.TrafficDate, time.Second)
	require.Equal(t, arg.Avgdailydldatamb, traffic2.Avgdailydldatamb)
	require.Equal(t, arg.Avgdailyuldatamb, traffic2.Avgdailyuldatamb)
	require.Equal(t, arg.Avgdailytotdatamb, traffic2.Avgdailytotdatamb)
	require.Equal(t, arg.Avgdailytotvoicemin, traffic2.Avgdailytotvoicemin)
	require.Equal(t, arg.Avgdailytotvideomin, traffic2.Avgdailytotvideomin)
	require.Equal(t, arg.Qci1Data, traffic2.Qci1Data)
	require.Equal(t, arg.Qci6Data, traffic2.Qci6Data)
	require.Equal(t, arg.Qci8Data, traffic2.Qci8Data)
	require.Equal(t, arg.QciOtherData, traffic2.QciOtherData)
	require.Equal(t, arg.Avgdailytotvoicemin4g, traffic2.Avgdailytotvoicemin4g)
	require.Equal(t, arg.Avgdailytotvoicemintotal, traffic2.Avgdailytotvoicemintotal)
	require.Equal(t, arg.Userdlthroughput, traffic2.Userdlthroughput)
	require.Equal(t, arg.Dlpacketlossrate, traffic2.Dlpacketlossrate)
	require.Equal(t, arg.Overallpsdropcallrate, traffic2.Overallpsdropcallrate)
	require.Equal(t, arg.Bhdldatamb, traffic2.Bhdldatamb)
	require.Equal(t, arg.Bhupdatamb, traffic2.Bhupdatamb)
	require.Equal(t, arg.Bhtotdatamb, traffic2.Bhtotdatamb)
	require.Equal(t, arg.Bhtotvoicemin, traffic2.Bhtotvoicemin)
	require.Equal(t, arg.Bhtotvideomin, traffic2.Bhtotvideomin)
	require.Equal(t, arg.Bhcsusers, traffic2.Bhcsusers)
	require.Equal(t, arg.Bhhsupausers, traffic2.Bhhsupausers)
	require.Equal(t, arg.Bhhsdpausers, traffic2.Bhhsdpausers)
	require.Equal(t, arg.Bhr99uldl, traffic2.Bhr99uldl)
	require.Equal(t, arg.Powercapacity, traffic2.Powercapacity)
	require.Equal(t, arg.Powerutilization, traffic2.Powerutilization)
	require.Equal(t, arg.Codecapacity, traffic2.Codecapacity)
	require.Equal(t, arg.Codeutilization, traffic2.Codeutilization)
	require.Equal(t, arg.Ceulcapacity, traffic2.Ceulcapacity)
	require.Equal(t, arg.Ceulutilization, traffic2.Ceulutilization)
	require.Equal(t, arg.Cedlcapacity, traffic2.Cedlcapacity)
	require.Equal(t, arg.Cedlutilization, traffic2.Cedlutilization)
	require.Equal(t, arg.Iubcapacity, traffic2.Iubcapacity)
	require.Equal(t, arg.Iubutlization, traffic2.Iubutlization)
	require.Equal(t, arg.Bhrrcusers, traffic2.Bhrrcusers)
	require.Equal(t, traffic1.CellID, traffic2.CellID)

}

func TestDeleteTraffic(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	tech := createRandomTech(t)
	block := createRandomBlock(t, town, clutter)
	band := createRandomBand(t, tech)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	site := createRandomSite(t, property, sitetype, vendor)
	carrier := createRandomCarrier(t, band)
	serviceareatype := createRandomServiceareatype(t)
	cell := createRandomCell(t, site, carrier, serviceareatype)
	traffic1 := createRandomTraffic(t, cell)
	err := testStore.DeleteTraffic(context.Background(), traffic1.ID)
	require.NoError(t, err)
	traffic2, err := testStore.GetTraffic0(context.Background(), traffic1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, traffic2)

}
