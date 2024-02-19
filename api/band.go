package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createBandRequest struct {
	BandName  string `json:"band_name" binding:"required"`
	BandDesc  string `json:"band_desc" binding:"required"`
	Size      int64  `json:"size" binding:"required"`
	StartFreq int64  `json:"start_freq" binding:"required"`
	EndFreq   int64  `json:"end_freq" binding:"required"`
	TechID    int64  `json:"tech_id" binding:"required"`
}

func (server *Server) createBand(ctx *gin.Context) {
	var req createBandRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBandParams{
		BandName:  req.BandName,
		BandDesc:  req.BandDesc,
		Size:      req.Size,
		StartFreq: req.StartFreq,
		EndFreq:   req.EndFreq,
		TechID:    req.TechID,
	}
	band, err := server.store.CreateBand(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, band)
}

type getBandRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBand0(ctx *gin.Context) {
	var req getBandRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	band, err := server.store.GetBand0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, band)
}

type getBandRequest1 struct {
	BandName string `uri:"band_name" binding:"required,min=1"`
}

func (server *Server) getBand1(ctx *gin.Context) {
	var req getBandRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	band, err := server.store.GetBand1(ctx, req.BandName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, band)
}

type listBandRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listBands(ctx *gin.Context) {
	var req listBandRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListBandsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	bands, err := server.store.ListBands(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, bands)
}

type updateBandRequest struct {
	ID        int64  `json:"id" binding:"required"`
	BandName  string `json:"band_name" binding:"required"`
	BandDesc  string `json:"band_desc" binding:"required"`
	Size      int64  `json:"size" binding:"required"`
	StartFreq int64  `json:"start_freq" binding:"required"`
	EndFreq   int64  `json:"end_freq" binding:"required"`
	TechID    int64  `json:"tech_id" binding:"required"`
}

func (server *Server) updateBand(ctx *gin.Context) {
	var req updateBandRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateBandParams{
		ID:        req.ID,
		BandName:  req.BandName,
		BandDesc:  req.BandDesc,
		Size:      req.Size,
		StartFreq: req.StartFreq,
		EndFreq:   req.EndFreq,
		TechID:    req.TechID,
	}
	band, err := server.store.UpdateBand(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, band)
}

type deleteBandRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteBand(ctx *gin.Context) {
	var req deleteBandRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteBand(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
