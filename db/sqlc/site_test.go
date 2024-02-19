package db

import (
	"context"
	"testing"
	"time"

	"github.com/naviscom/catalystx/util"
	"github.com/stretchr/testify/require"
)

func createRandomSite(t *testing.T, property Property, sitetype Sitetype, vendor Vendor) Site {
	arg := CreateSiteParams{
		SiteName:       util.RandomName(8),
		SiteNameOld:    util.RandomName(8),
		SiteIDGivin:    util.RandomName(8),
		SiteIDGivinOld: util.RandomName(8),
		Lac:            util.RandomName(8),
		Rac:            util.RandomName(8),
		Rnc:            util.RandomName(8),
		SiteOnAirDate:  time.Now().UTC(),
		PropertyID:     property.ID,
		SitetypeID:     sitetype.ID,
		VendorID:       vendor.ID,
	}
	site, err := testStore.CreateSite(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, site)
	require.Equal(t, arg.SiteName, site.SiteName)
	require.Equal(t, arg.SiteNameOld, site.SiteNameOld)
	require.Equal(t, arg.SiteIDGivin, site.SiteIDGivin)
	require.Equal(t, arg.SiteIDGivinOld, site.SiteIDGivinOld)
	require.Equal(t, arg.Lac, site.Lac)
	require.Equal(t, arg.Rac, site.Rac)
	require.Equal(t, arg.Rnc, site.Rnc)
	require.WithinDuration(t, arg.SiteOnAirDate, site.SiteOnAirDate, time.Second)
	require.Equal(t, arg.PropertyID, site.PropertyID)
	require.Equal(t, arg.SitetypeID, site.SitetypeID)
	require.Equal(t, arg.VendorID, site.VendorID)
	return site
}

func TestCreateSite(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	createRandomSite(t, property, sitetype, vendor)
}

func TestGetSite0(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	site1 := createRandomSite(t, property, sitetype, vendor)
	site2, err := testStore.GetSite0(context.Background(), site1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, site2)

	require.Equal(t, site1.ID, site2.ID)
	require.Equal(t, site1.SiteName, site2.SiteName)
	require.Equal(t, site1.SiteNameOld, site2.SiteNameOld)
	require.Equal(t, site1.SiteIDGivin, site2.SiteIDGivin)
	require.Equal(t, site1.SiteIDGivinOld, site2.SiteIDGivinOld)
	require.Equal(t, site1.Lac, site2.Lac)
	require.Equal(t, site1.Rac, site2.Rac)
	require.Equal(t, site1.Rnc, site2.Rnc)
	require.WithinDuration(t, site1.SiteOnAirDate, site2.SiteOnAirDate, time.Second)
	require.Equal(t, site1.PropertyID, site2.PropertyID)
	require.Equal(t, site1.SitetypeID, site2.SitetypeID)
	require.Equal(t, site1.VendorID, site2.VendorID)
}

func TestGetSite1(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	site1 := createRandomSite(t, property, sitetype, vendor)
	site2, err := testStore.GetSite1(context.Background(), site1.SiteName)
	require.NoError(t, err)
	require.NotEmpty(t, site2)

	require.Equal(t, site1.ID, site2.ID)
	require.Equal(t, site1.SiteName, site2.SiteName)
	require.Equal(t, site1.SiteNameOld, site2.SiteNameOld)
	require.Equal(t, site1.SiteIDGivin, site2.SiteIDGivin)
	require.Equal(t, site1.SiteIDGivinOld, site2.SiteIDGivinOld)
	require.Equal(t, site1.Lac, site2.Lac)
	require.Equal(t, site1.Rac, site2.Rac)
	require.Equal(t, site1.Rnc, site2.Rnc)
	require.WithinDuration(t, site1.SiteOnAirDate, site2.SiteOnAirDate, time.Second)
	require.Equal(t, site1.PropertyID, site2.PropertyID)
	require.Equal(t, site1.SitetypeID, site2.SitetypeID)
	require.Equal(t, site1.VendorID, site2.VendorID)
}

func TestListSites(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	for i := 0; i < 10; i++ {
		createRandomSite(t, property, sitetype, vendor)

	}
	arg := ListSitesParams{
		PropertyID: property.ID,
		SitetypeID: sitetype.ID,
		VendorID:   vendor.ID,
		Limit:      5,
		Offset:     5,
	}
	sites, err := testStore.ListSites(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, sites, 5)

	for _, site := range sites {
		require.NotEmpty(t, site)
		require.True(t, site.PropertyID == property.ID || site.SitetypeID == sitetype.ID || site.VendorID == vendor.ID)
	}
}

func TestUpdateSite(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	site1 := createRandomSite(t, property, sitetype, vendor)
	arg := UpdateSiteParams{
		ID:             site1.ID,
		SiteName:       util.RandomName(8),
		SiteNameOld:    util.RandomName(8),
		SiteIDGivin:    util.RandomName(8),
		SiteIDGivinOld: util.RandomName(8),
		Lac:            util.RandomName(8),
		Rac:            util.RandomName(8),
		Rnc:            util.RandomName(8),
		SiteOnAirDate:  time.Now().UTC(),
		PropertyID:     property.ID,
		SitetypeID:     sitetype.ID,
		VendorID:       vendor.ID,
	}
	site2, err := testStore.UpdateSite(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, site2)

	require.Equal(t, site1.ID, site2.ID)
	require.Equal(t, arg.SiteName, site2.SiteName)
	require.Equal(t, arg.SiteNameOld, site2.SiteNameOld)
	require.Equal(t, arg.SiteIDGivin, site2.SiteIDGivin)
	require.Equal(t, arg.SiteIDGivinOld, site2.SiteIDGivinOld)
	require.Equal(t, arg.Lac, site2.Lac)
	require.Equal(t, arg.Rac, site2.Rac)
	require.Equal(t, arg.Rnc, site2.Rnc)
	require.WithinDuration(t, arg.SiteOnAirDate, site2.SiteOnAirDate, time.Second)
	require.Equal(t, site1.PropertyID, site2.PropertyID)
	require.Equal(t, site1.SitetypeID, site2.SitetypeID)
	require.Equal(t, site1.VendorID, site2.VendorID)

}

func TestDeleteSite(t *testing.T) {
	continent := createRandomContinent(t)
	country := createRandomCountry(t, continent)
	area := createRandomArea(t)
	state := createRandomState(t, country, area)
	city := createRandomCity(t, state)
	district := createRandomDistrict(t, city)
	town := createRandomTown(t, district)
	clutter := createRandomClutter(t)
	block := createRandomBlock(t, town, clutter)
	property := createRandomProperty(t, block)
	sitetype := createRandomSitetype(t)
	vendor := createRandomVendor(t)
	site1 := createRandomSite(t, property, sitetype, vendor)
	err := testStore.DeleteSite(context.Background(), site1.ID)
	require.NoError(t, err)
	site2, err := testStore.GetSite0(context.Background(), site1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, site2)

}
