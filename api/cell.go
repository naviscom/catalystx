package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createCellRequest struct {
	CellName          string `json:"cell_name" binding:"required"`
	CellNameOld       string `json:"cell_name_old" binding:"required"`
	CellIDGivin       string `json:"cell_id_givin" binding:"required"`
	CellIDGivinOld    string `json:"cell_id_givin_old" binding:"required"`
	SectorName        string `json:"sector_name" binding:"required"`
	Uplinkuarfcn      string `json:"uplinkuarfcn" binding:"required"`
	Downlinkuarfcn    string `json:"downlinkuarfcn" binding:"required"`
	Dlprscramblecode  string `json:"dlprscramblecode" binding:"required"`
	Azimuth           string `json:"azimuth" binding:"required"`
	Height            string `json:"height" binding:"required"`
	Etilt             string `json:"etilt" binding:"required"`
	Mtilt             string `json:"mtilt" binding:"required"`
	Antennatype       string `json:"antennatype" binding:"required"`
	Antennamodel      string `json:"antennamodel" binding:"required"`
	Ecgi              string `json:"ecgi" binding:"required"`
	SiteID            int64  `json:"site_id" binding:"required"`
	CarrierID         int64  `json:"carrier_id" binding:"required"`
	ServiceareatypeID int64  `json:"serviceareatype_id" binding:"required"`
}

func (server *Server) createCell(ctx *gin.Context) {
	var req createCellRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCellParams{
		CellName:          req.CellName,
		CellNameOld:       req.CellNameOld,
		CellIDGivin:       req.CellIDGivin,
		CellIDGivinOld:    req.CellIDGivinOld,
		SectorName:        req.SectorName,
		Uplinkuarfcn:      req.Uplinkuarfcn,
		Downlinkuarfcn:    req.Downlinkuarfcn,
		Dlprscramblecode:  req.Dlprscramblecode,
		Azimuth:           req.Azimuth,
		Height:            req.Height,
		Etilt:             req.Etilt,
		Mtilt:             req.Mtilt,
		Antennatype:       req.Antennatype,
		Antennamodel:      req.Antennamodel,
		Ecgi:              req.Ecgi,
		SiteID:            req.SiteID,
		CarrierID:         req.CarrierID,
		ServiceareatypeID: req.ServiceareatypeID,
	}
	cell, err := server.store.CreateCell(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cell)
}

type getCellRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCell0(ctx *gin.Context) {
	var req getCellRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	cell, err := server.store.GetCell0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cell)
}

type getCellRequest1 struct {
	CellName string `uri:"cell_name" binding:"required,min=1"`
}

func (server *Server) getCell1(ctx *gin.Context) {
	var req getCellRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	cell, err := server.store.GetCell1(ctx, req.CellName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cell)
}

type listCellRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCells(ctx *gin.Context) {
	var req listCellRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCellsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	cells, err := server.store.ListCells(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cells)
}

type updateCellRequest struct {
	ID                int64  `json:"id" binding:"required"`
	CellName          string `json:"cell_name" binding:"required"`
	CellNameOld       string `json:"cell_name_old" binding:"required"`
	CellIDGivin       string `json:"cell_id_givin" binding:"required"`
	CellIDGivinOld    string `json:"cell_id_givin_old" binding:"required"`
	SectorName        string `json:"sector_name" binding:"required"`
	Uplinkuarfcn      string `json:"uplinkuarfcn" binding:"required"`
	Downlinkuarfcn    string `json:"downlinkuarfcn" binding:"required"`
	Dlprscramblecode  string `json:"dlprscramblecode" binding:"required"`
	Azimuth           string `json:"azimuth" binding:"required"`
	Height            string `json:"height" binding:"required"`
	Etilt             string `json:"etilt" binding:"required"`
	Mtilt             string `json:"mtilt" binding:"required"`
	Antennatype       string `json:"antennatype" binding:"required"`
	Antennamodel      string `json:"antennamodel" binding:"required"`
	Ecgi              string `json:"ecgi" binding:"required"`
	SiteID            int64  `json:"site_id" binding:"required"`
	CarrierID         int64  `json:"carrier_id" binding:"required"`
	ServiceareatypeID int64  `json:"serviceareatype_id" binding:"required"`
}

func (server *Server) updateCell(ctx *gin.Context) {
	var req updateCellRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCellParams{
		ID:                req.ID,
		CellName:          req.CellName,
		CellNameOld:       req.CellNameOld,
		CellIDGivin:       req.CellIDGivin,
		CellIDGivinOld:    req.CellIDGivinOld,
		SectorName:        req.SectorName,
		Uplinkuarfcn:      req.Uplinkuarfcn,
		Downlinkuarfcn:    req.Downlinkuarfcn,
		Dlprscramblecode:  req.Dlprscramblecode,
		Azimuth:           req.Azimuth,
		Height:            req.Height,
		Etilt:             req.Etilt,
		Mtilt:             req.Mtilt,
		Antennatype:       req.Antennatype,
		Antennamodel:      req.Antennamodel,
		Ecgi:              req.Ecgi,
		SiteID:            req.SiteID,
		CarrierID:         req.CarrierID,
		ServiceareatypeID: req.ServiceareatypeID,
	}
	cell, err := server.store.UpdateCell(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cell)
}

type deleteCellRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteCell(ctx *gin.Context) {
	var req deleteCellRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCell(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
