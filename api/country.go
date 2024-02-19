package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createCountryRequest struct {
	CountryName string `json:"country_name" binding:"required"`
	CountryDesc string `json:"country_desc" binding:"required"`
	ContinentID int64  `json:"continent_id" binding:"required"`
}

func (server *Server) createCountry(ctx *gin.Context) {
	var req createCountryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCountryParams{
		CountryName: req.CountryName,
		CountryDesc: req.CountryDesc,
		ContinentID: req.ContinentID,
	}
	country, err := server.store.CreateCountry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, country)
}

type getCountryRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCountry0(ctx *gin.Context) {
	var req getCountryRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	country, err := server.store.GetCountry0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, country)
}

type getCountryRequest1 struct {
	CountryName string `uri:"country_name" binding:"required,min=1"`
}

func (server *Server) getCountry1(ctx *gin.Context) {
	var req getCountryRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	country, err := server.store.GetCountry1(ctx, req.CountryName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, country)
}

type listCountryRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCountries(ctx *gin.Context) {
	var req listCountryRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCountriesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	countries, err := server.store.ListCountries(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, countries)
}

type updateCountryRequest struct {
	ID          int64  `json:"id" binding:"required"`
	CountryName string `json:"country_name" binding:"required"`
	CountryDesc string `json:"country_desc" binding:"required"`
	ContinentID int64  `json:"continent_id" binding:"required"`
}

func (server *Server) updateCountry(ctx *gin.Context) {
	var req updateCountryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCountryParams{
		ID:          req.ID,
		CountryName: req.CountryName,
		CountryDesc: req.CountryDesc,
		ContinentID: req.ContinentID,
	}
	country, err := server.store.UpdateCountry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, country)
}

type deleteCountryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteCountry(ctx *gin.Context) {
	var req deleteCountryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCountry(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
