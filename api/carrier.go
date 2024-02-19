package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createCarrierRequest struct {
	CarrierName string `json:"carrier_name" binding:"required"`
	CarrierDesc string `json:"carrier_desc" binding:"required"`
	Size        int64  `json:"size" binding:"required"`
	StartFreq   int64  `json:"start_freq" binding:"required"`
	EndFreq     int64  `json:"end_freq" binding:"required"`
	BandID      int64  `json:"band_id" binding:"required"`
}

func (server *Server) createCarrier(ctx *gin.Context) {
	var req createCarrierRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCarrierParams{
		CarrierName: req.CarrierName,
		CarrierDesc: req.CarrierDesc,
		Size:        req.Size,
		StartFreq:   req.StartFreq,
		EndFreq:     req.EndFreq,
		BandID:      req.BandID,
	}
	carrier, err := server.store.CreateCarrier(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, carrier)
}

type getCarrierRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCarrier0(ctx *gin.Context) {
	var req getCarrierRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	carrier, err := server.store.GetCarrier0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, carrier)
}

type getCarrierRequest1 struct {
	CarrierName string `uri:"carrier_name" binding:"required,min=1"`
}

func (server *Server) getCarrier1(ctx *gin.Context) {
	var req getCarrierRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	carrier, err := server.store.GetCarrier1(ctx, req.CarrierName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, carrier)
}

type listCarrierRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCarriers(ctx *gin.Context) {
	var req listCarrierRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCarriersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	carriers, err := server.store.ListCarriers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, carriers)
}

type updateCarrierRequest struct {
	ID          int64  `json:"id" binding:"required"`
	CarrierName string `json:"carrier_name" binding:"required"`
	CarrierDesc string `json:"carrier_desc" binding:"required"`
	Size        int64  `json:"size" binding:"required"`
	StartFreq   int64  `json:"start_freq" binding:"required"`
	EndFreq     int64  `json:"end_freq" binding:"required"`
	BandID      int64  `json:"band_id" binding:"required"`
}

func (server *Server) updateCarrier(ctx *gin.Context) {
	var req updateCarrierRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCarrierParams{
		ID:          req.ID,
		CarrierName: req.CarrierName,
		CarrierDesc: req.CarrierDesc,
		Size:        req.Size,
		StartFreq:   req.StartFreq,
		EndFreq:     req.EndFreq,
		BandID:      req.BandID,
	}
	carrier, err := server.store.UpdateCarrier(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, carrier)
}

type deleteCarrierRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteCarrier(ctx *gin.Context) {
	var req deleteCarrierRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCarrier(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
