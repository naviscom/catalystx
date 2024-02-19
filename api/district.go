package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createDistrictRequest struct {
	DistrictName string `json:"district_name" binding:"required"`
	DistrictDesc string `json:"district_desc" binding:"required"`
	CityID       int64  `json:"city_id" binding:"required"`
}

func (server *Server) createDistrict(ctx *gin.Context) {
	var req createDistrictRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateDistrictParams{
		DistrictName: req.DistrictName,
		DistrictDesc: req.DistrictDesc,
		CityID:       req.CityID,
	}
	district, err := server.store.CreateDistrict(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, district)
}

type getDistrictRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getDistrict0(ctx *gin.Context) {
	var req getDistrictRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	district, err := server.store.GetDistrict0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, district)
}

type getDistrictRequest1 struct {
	DistrictName string `uri:"district_name" binding:"required,min=1"`
}

func (server *Server) getDistrict1(ctx *gin.Context) {
	var req getDistrictRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	district, err := server.store.GetDistrict1(ctx, req.DistrictName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, district)
}

type listDistrictRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listDistricts(ctx *gin.Context) {
	var req listDistrictRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListDistrictsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	districts, err := server.store.ListDistricts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, districts)
}

type updateDistrictRequest struct {
	ID           int64  `json:"id" binding:"required"`
	DistrictName string `json:"district_name" binding:"required"`
	DistrictDesc string `json:"district_desc" binding:"required"`
	CityID       int64  `json:"city_id" binding:"required"`
}

func (server *Server) updateDistrict(ctx *gin.Context) {
	var req updateDistrictRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateDistrictParams{
		ID:           req.ID,
		DistrictName: req.DistrictName,
		DistrictDesc: req.DistrictDesc,
		CityID:       req.CityID,
	}
	district, err := server.store.UpdateDistrict(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, district)
}

type deleteDistrictRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteDistrict(ctx *gin.Context) {
	var req deleteDistrictRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteDistrict(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
