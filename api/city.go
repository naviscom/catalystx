package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createCityRequest struct {
	CityName string `json:"city_name" binding:"required"`
	CityDesc string `json:"city_desc" binding:"required"`
	StateID  int64  `json:"state_id" binding:"required"`
}

func (server *Server) createCity(ctx *gin.Context) {
	var req createCityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCityParams{
		CityName: req.CityName,
		CityDesc: req.CityDesc,
		StateID:  req.StateID,
	}
	city, err := server.store.CreateCity(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, city)
}

type getCityRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCity0(ctx *gin.Context) {
	var req getCityRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	city, err := server.store.GetCity0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, city)
}

type getCityRequest1 struct {
	CityName string `uri:"city_name" binding:"required,min=1"`
}

func (server *Server) getCity1(ctx *gin.Context) {
	var req getCityRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	city, err := server.store.GetCity1(ctx, req.CityName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, city)
}

type listCityRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCities(ctx *gin.Context) {
	var req listCityRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCitiesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	cities, err := server.store.ListCities(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cities)
}

type updateCityRequest struct {
	ID       int64  `json:"id" binding:"required"`
	CityName string `json:"city_name" binding:"required"`
	CityDesc string `json:"city_desc" binding:"required"`
	StateID  int64  `json:"state_id" binding:"required"`
}

func (server *Server) updateCity(ctx *gin.Context) {
	var req updateCityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCityParams{
		ID:       req.ID,
		CityName: req.CityName,
		CityDesc: req.CityDesc,
		StateID:  req.StateID,
	}
	city, err := server.store.UpdateCity(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, city)
}

type deleteCityRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteCity(ctx *gin.Context) {
	var req deleteCityRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCity(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
