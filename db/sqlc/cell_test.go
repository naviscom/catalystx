package db

import (
	"context"
	"testing"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomCell(t *testing.T, site Site, carrier Carrier, serviceareatype Serviceareatype) Cell {
	arg := CreateCellParams{
		CellName:          util.RandomName(8),
		CellNameOld:       util.RandomName(8),
		CellIDGivin:       util.RandomName(8),
		CellIDGivinOld:    util.RandomName(8),
		SectorName:        util.RandomName(8),
		Uplinkuarfcn:      util.RandomName(8),
		Downlinkuarfcn:    util.RandomName(8),
		Dlprscramblecode:  util.RandomName(8),
		Azimuth:           util.RandomName(8),
		Height:            util.RandomName(8),
		Etilt:             util.RandomName(8),
		Mtilt:             util.RandomName(8),
		Antennatype:       util.RandomName(8),
		Antennamodel:      util.RandomName(8),
		Ecgi:              util.RandomName(8),
		SiteID:            site.ID,
		CarrierID:         carrier.ID,
		ServiceareatypeID: serviceareatype.ID,
	}
	cell, err := testStore.CreateCell(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cell)
	require.Equal(t, arg.CellName, cell.CellName)
	require.Equal(t, arg.CellNameOld, cell.CellNameOld)
	require.Equal(t, arg.CellIDGivin, cell.CellIDGivin)
	require.Equal(t, arg.CellIDGivinOld, cell.CellIDGivinOld)
	require.Equal(t, arg.SectorName, cell.SectorName)
	require.Equal(t, arg.Uplinkuarfcn, cell.Uplinkuarfcn)
	require.Equal(t, arg.Downlinkuarfcn, cell.Downlinkuarfcn)
	require.Equal(t, arg.Dlprscramblecode, cell.Dlprscramblecode)
	require.Equal(t, arg.Azimuth, cell.Azimuth)
	require.Equal(t, arg.Height, cell.Height)
	require.Equal(t, arg.Etilt, cell.Etilt)
	require.Equal(t, arg.Mtilt, cell.Mtilt)
	require.Equal(t, arg.Antennatype, cell.Antennatype)
	require.Equal(t, arg.Antennamodel, cell.Antennamodel)
	require.Equal(t, arg.Ecgi, cell.Ecgi)
	require.Equal(t, arg.SiteID, cell.SiteID)
	require.Equal(t, arg.CarrierID, cell.CarrierID)
	require.Equal(t, arg.ServiceareatypeID, cell.ServiceareatypeID)
	return cell
}

func TestCreateCell(t *testing.T) {
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
	createRandomCell(t, site, carrier, serviceareatype)
}

func TestGetCell0(t *testing.T) {
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
	cell1 := createRandomCell(t, site, carrier, serviceareatype)
	cell2, err := testStore.GetCell0(context.Background(), cell1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, cell2)

	require.Equal(t, cell1.ID, cell2.ID)
	require.Equal(t, cell1.CellName, cell2.CellName)
	require.Equal(t, cell1.CellNameOld, cell2.CellNameOld)
	require.Equal(t, cell1.CellIDGivin, cell2.CellIDGivin)
	require.Equal(t, cell1.CellIDGivinOld, cell2.CellIDGivinOld)
	require.Equal(t, cell1.SectorName, cell2.SectorName)
	require.Equal(t, cell1.Uplinkuarfcn, cell2.Uplinkuarfcn)
	require.Equal(t, cell1.Downlinkuarfcn, cell2.Downlinkuarfcn)
	require.Equal(t, cell1.Dlprscramblecode, cell2.Dlprscramblecode)
	require.Equal(t, cell1.Azimuth, cell2.Azimuth)
	require.Equal(t, cell1.Height, cell2.Height)
	require.Equal(t, cell1.Etilt, cell2.Etilt)
	require.Equal(t, cell1.Mtilt, cell2.Mtilt)
	require.Equal(t, cell1.Antennatype, cell2.Antennatype)
	require.Equal(t, cell1.Antennamodel, cell2.Antennamodel)
	require.Equal(t, cell1.Ecgi, cell2.Ecgi)
	require.Equal(t, cell1.SiteID, cell2.SiteID)
	require.Equal(t, cell1.CarrierID, cell2.CarrierID)
	require.Equal(t, cell1.ServiceareatypeID, cell2.ServiceareatypeID)
}

func TestGetCell1(t *testing.T) {
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
	cell1 := createRandomCell(t, site, carrier, serviceareatype)
	cell2, err := testStore.GetCell1(context.Background(), cell1.CellName)
	require.NoError(t, err)
	require.NotEmpty(t, cell2)

	require.Equal(t, cell1.ID, cell2.ID)
	require.Equal(t, cell1.CellName, cell2.CellName)
	require.Equal(t, cell1.CellNameOld, cell2.CellNameOld)
	require.Equal(t, cell1.CellIDGivin, cell2.CellIDGivin)
	require.Equal(t, cell1.CellIDGivinOld, cell2.CellIDGivinOld)
	require.Equal(t, cell1.SectorName, cell2.SectorName)
	require.Equal(t, cell1.Uplinkuarfcn, cell2.Uplinkuarfcn)
	require.Equal(t, cell1.Downlinkuarfcn, cell2.Downlinkuarfcn)
	require.Equal(t, cell1.Dlprscramblecode, cell2.Dlprscramblecode)
	require.Equal(t, cell1.Azimuth, cell2.Azimuth)
	require.Equal(t, cell1.Height, cell2.Height)
	require.Equal(t, cell1.Etilt, cell2.Etilt)
	require.Equal(t, cell1.Mtilt, cell2.Mtilt)
	require.Equal(t, cell1.Antennatype, cell2.Antennatype)
	require.Equal(t, cell1.Antennamodel, cell2.Antennamodel)
	require.Equal(t, cell1.Ecgi, cell2.Ecgi)
	require.Equal(t, cell1.SiteID, cell2.SiteID)
	require.Equal(t, cell1.CarrierID, cell2.CarrierID)
	require.Equal(t, cell1.ServiceareatypeID, cell2.ServiceareatypeID)
}

func TestListCells(t *testing.T) {
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
	for i := 0; i < 10; i++ {
		createRandomCell(t, site, carrier, serviceareatype)

	}
	arg := ListCellsParams{
		SiteID:            site.ID,
		CarrierID:         carrier.ID,
		ServiceareatypeID: serviceareatype.ID,
		Limit:             5,
		Offset:            5,
	}
	cells, err := testStore.ListCells(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, cells, 5)

	for _, cell := range cells {
		require.NotEmpty(t, cell)
		require.True(t, cell.SiteID == site.ID || cell.CarrierID == carrier.ID || cell.ServiceareatypeID == serviceareatype.ID)
	}
}

func TestUpdateCell(t *testing.T) {
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
	cell1 := createRandomCell(t, site, carrier, serviceareatype)
	arg := UpdateCellParams{
		ID:                cell1.ID,
		CellName:          util.RandomName(8),
		CellNameOld:       util.RandomName(8),
		CellIDGivin:       util.RandomName(8),
		CellIDGivinOld:    util.RandomName(8),
		SectorName:        util.RandomName(8),
		Uplinkuarfcn:      util.RandomName(8),
		Downlinkuarfcn:    util.RandomName(8),
		Dlprscramblecode:  util.RandomName(8),
		Azimuth:           util.RandomName(8),
		Height:            util.RandomName(8),
		Etilt:             util.RandomName(8),
		Mtilt:             util.RandomName(8),
		Antennatype:       util.RandomName(8),
		Antennamodel:      util.RandomName(8),
		Ecgi:              util.RandomName(8),
		SiteID:            site.ID,
		CarrierID:         carrier.ID,
		ServiceareatypeID: serviceareatype.ID,
	}
	cell2, err := testStore.UpdateCell(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cell2)

	require.Equal(t, cell1.ID, cell2.ID)
	require.Equal(t, arg.CellName, cell2.CellName)
	require.Equal(t, arg.CellNameOld, cell2.CellNameOld)
	require.Equal(t, arg.CellIDGivin, cell2.CellIDGivin)
	require.Equal(t, arg.CellIDGivinOld, cell2.CellIDGivinOld)
	require.Equal(t, arg.SectorName, cell2.SectorName)
	require.Equal(t, arg.Uplinkuarfcn, cell2.Uplinkuarfcn)
	require.Equal(t, arg.Downlinkuarfcn, cell2.Downlinkuarfcn)
	require.Equal(t, arg.Dlprscramblecode, cell2.Dlprscramblecode)
	require.Equal(t, arg.Azimuth, cell2.Azimuth)
	require.Equal(t, arg.Height, cell2.Height)
	require.Equal(t, arg.Etilt, cell2.Etilt)
	require.Equal(t, arg.Mtilt, cell2.Mtilt)
	require.Equal(t, arg.Antennatype, cell2.Antennatype)
	require.Equal(t, arg.Antennamodel, cell2.Antennamodel)
	require.Equal(t, arg.Ecgi, cell2.Ecgi)
	require.Equal(t, cell1.SiteID, cell2.SiteID)
	require.Equal(t, cell1.CarrierID, cell2.CarrierID)
	require.Equal(t, cell1.ServiceareatypeID, cell2.ServiceareatypeID)

}

func TestDeleteCell(t *testing.T) {
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
	cell1 := createRandomCell(t, site, carrier, serviceareatype)
	err := testStore.DeleteCell(context.Background(), cell1.ID)
	require.NoError(t, err)
	cell2, err := testStore.GetCell0(context.Background(), cell1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, cell2)

}
