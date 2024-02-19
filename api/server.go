package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
	"github.com/naviscom/catalystx/util"
)

// // Server serves HTTP requests.
type Server struct {
	config util.Config
	store  *db.Store
	//tokenMaker token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP Gin server and set up routing & CORS.
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
		//tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	//router.POST("/users", server.createUser)
	//router.POST("/users/login", server.loginUser)
	//router.POST("/tokens/renew_access", server.renewAccessToken)
	//authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	router.POST("/serviceareatypes", server.createServiceareatype)
	router.GET("/serviceareatypes0/:id", server.getServiceareatype0)
	router.GET("/serviceareatypes1/:serviceareatype_name", server.getServiceareatype1)
	router.GET("/serviceareatypes", server.listServiceareatypes)
	router.POST("/updateserviceareatypes", server.updateServiceareatype)
	router.GET("/deleteserviceareatypes/:id", server.deleteServiceareatype)
	router.POST("/areas", server.createArea)
	router.GET("/areas0/:id", server.getArea0)
	router.GET("/areas1/:area_name", server.getArea1)
	router.GET("/areas", server.listAreas)
	router.POST("/updateareas", server.updateArea)
	router.GET("/deleteareas/:id", server.deleteArea)
	router.POST("/clutters", server.createClutter)
	router.GET("/clutters0/:id", server.getClutter0)
	router.GET("/clutters1/:clutter_name", server.getClutter1)
	router.GET("/clutters", server.listClutters)
	router.POST("/updateclutters", server.updateClutter)
	router.GET("/deleteclutters/:id", server.deleteClutter)
	router.POST("/sitetypes", server.createSitetype)
	router.GET("/sitetypes0/:id", server.getSitetype0)
	router.GET("/sitetypes1/:type_name", server.getSitetype1)
	router.GET("/sitetypes", server.listSitetypes)
	router.POST("/updatesitetypes", server.updateSitetype)
	router.GET("/deletesitetypes/:id", server.deleteSitetype)
	router.POST("/vendors", server.createVendor)
	router.GET("/vendors0/:id", server.getVendor0)
	router.GET("/vendors1/:vendor_name", server.getVendor1)
	router.GET("/vendors", server.listVendors)
	router.POST("/updatevendors", server.updateVendor)
	router.GET("/deletevendors/:id", server.deleteVendor)
	router.POST("/techs", server.createTech)
	router.GET("/techs0/:id", server.getTech0)
	router.GET("/techs1/:tech_name", server.getTech1)
	router.GET("/techs", server.listTechs)
	router.POST("/updatetechs", server.updateTech)
	router.GET("/deletetechs/:id", server.deleteTech)
	router.POST("/bands", server.createBand)
	router.GET("/bands0/:id", server.getBand0)
	router.GET("/bands1/:band_name", server.getBand1)
	router.GET("/bands", server.listBands)
	router.POST("/updatebands", server.updateBand)
	router.GET("/deletebands/:id", server.deleteBand)
	router.POST("/carriers", server.createCarrier)
	router.GET("/carriers0/:id", server.getCarrier0)
	router.GET("/carriers1/:carrier_name", server.getCarrier1)
	router.GET("/carriers", server.listCarriers)
	router.POST("/updatecarriers", server.updateCarrier)
	router.GET("/deletecarriers/:id", server.deleteCarrier)
	router.POST("/continents", server.createContinent)
	router.GET("/continents0/:id", server.getContinent0)
	router.GET("/continents1/:continent_name", server.getContinent1)
	router.GET("/continents", server.listContinents)
	router.POST("/updatecontinents", server.updateContinent)
	router.GET("/deletecontinents/:id", server.deleteContinent)
	router.POST("/countries", server.createCountry)
	router.GET("/countries0/:id", server.getCountry0)
	router.GET("/countries1/:country_name", server.getCountry1)
	router.GET("/countries", server.listCountries)
	router.POST("/updatecountries", server.updateCountry)
	router.GET("/deletecountries/:id", server.deleteCountry)
	router.POST("/states", server.createState)
	router.GET("/states0/:id", server.getState0)
	router.GET("/states1/:state_name", server.getState1)
	router.GET("/states", server.listStates)
	router.POST("/updatestates", server.updateState)
	router.GET("/deletestates/:id", server.deleteState)
	router.POST("/cities", server.createCity)
	router.GET("/cities0/:id", server.getCity0)
	router.GET("/cities1/:city_name", server.getCity1)
	router.GET("/cities", server.listCities)
	router.POST("/updatecities", server.updateCity)
	router.GET("/deletecities/:id", server.deleteCity)
	router.POST("/districts", server.createDistrict)
	router.GET("/districts0/:id", server.getDistrict0)
	router.GET("/districts1/:district_name", server.getDistrict1)
	router.GET("/districts", server.listDistricts)
	router.POST("/updatedistricts", server.updateDistrict)
	router.GET("/deletedistricts/:id", server.deleteDistrict)
	router.POST("/towns", server.createTown)
	router.GET("/towns0/:id", server.getTown0)
	router.GET("/towns1/:town_name", server.getTown1)
	router.GET("/towns", server.listTowns)
	router.POST("/updatetowns", server.updateTown)
	router.GET("/deletetowns/:id", server.deleteTown)
	router.POST("/blocks", server.createBlock)
	router.GET("/blocks0/:id", server.getBlock0)
	router.GET("/blocks1/:block_name", server.getBlock1)
	router.GET("/blocks", server.listBlocks)
	router.POST("/updateblocks", server.updateBlock)
	router.GET("/deleteblocks/:id", server.deleteBlock)
	router.POST("/properties", server.createProperty)
	router.GET("/properties0/:id", server.getProperty0)
	router.GET("/properties1/:property_name", server.getProperty1)
	router.GET("/properties", server.listProperties)
	router.POST("/updateproperties", server.updateProperty)
	router.GET("/deleteproperties/:id", server.deleteProperty)
	router.POST("/sites", server.createSite)
	router.GET("/sites0/:id", server.getSite0)
	router.GET("/sites1/:site_name", server.getSite1)
	router.GET("/sites", server.listSites)
	router.POST("/updatesites", server.updateSite)
	router.GET("/deletesites/:id", server.deleteSite)
	router.POST("/cells", server.createCell)
	router.GET("/cells0/:id", server.getCell0)
	router.GET("/cells1/:cell_name", server.getCell1)
	router.GET("/cells", server.listCells)
	router.POST("/updatecells", server.updateCell)
	router.GET("/deletecells/:id", server.deleteCell)
	router.POST("/traffic", server.createTraffic)
	router.GET("/traffic0/:id", server.getTraffic0)
	router.GET("/traffic", server.listTraffic)
	router.POST("/updatetraffic", server.updateTraffic)
	router.GET("/deletetraffic/:id", server.deleteTraffic)
	router.POST("/updatenetworkdb", server.updatenetworkdb)

	// setting the Gin router to Default() allows all origins
	router.Use(cors.Default())

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
