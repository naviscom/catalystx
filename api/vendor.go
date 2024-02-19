package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createVendorRequest struct {
	VendorName string `json:"vendor_name" binding:"required"`
	VendorDesc string `json:"vendor_desc" binding:"required"`
}

func (server *Server) createVendor(ctx *gin.Context) {
	var req createVendorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateVendorParams{
		VendorName: req.VendorName,
		VendorDesc: req.VendorDesc,
	}
	vendor, err := server.store.CreateVendor(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, vendor)
}

type getVendorRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getVendor0(ctx *gin.Context) {
	var req getVendorRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	vendor, err := server.store.GetVendor0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, vendor)
}

type getVendorRequest1 struct {
	VendorName string `uri:"vendor_name" binding:"required,min=1"`
}

func (server *Server) getVendor1(ctx *gin.Context) {
	var req getVendorRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	vendor, err := server.store.GetVendor1(ctx, req.VendorName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, vendor)
}

type listVendorRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listVendors(ctx *gin.Context) {
	var req listVendorRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListVendorsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	vendors, err := server.store.ListVendors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, vendors)
}

type updateVendorRequest struct {
	ID         int64  `json:"id" binding:"required"`
	VendorName string `json:"vendor_name" binding:"required"`
	VendorDesc string `json:"vendor_desc" binding:"required"`
}

func (server *Server) updateVendor(ctx *gin.Context) {
	var req updateVendorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateVendorParams{
		ID:         req.ID,
		VendorName: req.VendorName,
		VendorDesc: req.VendorDesc,
	}
	vendor, err := server.store.UpdateVendor(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, vendor)
}

type deleteVendorRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteVendor(ctx *gin.Context) {
	var req deleteVendorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteVendor(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
