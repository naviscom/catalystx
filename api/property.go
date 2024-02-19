package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createPropertyRequest struct {
	PropertyName string  `json:"property_name" binding:"required"`
	Lat          float32 `json:"lat" binding:"required"`
	Long         float32 `json:"long" binding:"required"`
	BlockID      int64   `json:"block_id" binding:"required"`
}

func (server *Server) createProperty(ctx *gin.Context) {
	var req createPropertyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePropertyParams{
		PropertyName: req.PropertyName,
		Lat:          req.Lat,
		Long:         req.Long,
		BlockID:      req.BlockID,
	}
	property, err := server.store.CreateProperty(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, property)
}

type getPropertyRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProperty0(ctx *gin.Context) {
	var req getPropertyRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	property, err := server.store.GetProperty0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, property)
}

type getPropertyRequest1 struct {
	PropertyName string `uri:"property_name" binding:"required,min=1"`
}

func (server *Server) getProperty1(ctx *gin.Context) {
	var req getPropertyRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	property, err := server.store.GetProperty1(ctx, req.PropertyName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, property)
}

type listPropertyRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listProperties(ctx *gin.Context) {
	var req listPropertyRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPropertiesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	properties, err := server.store.ListProperties(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, properties)
}

type updatePropertyRequest struct {
	ID           int64   `json:"id" binding:"required"`
	PropertyName string  `json:"property_name" binding:"required"`
	Lat          float32 `json:"lat" binding:"required"`
	Long         float32 `json:"long" binding:"required"`
	BlockID      int64   `json:"block_id" binding:"required"`
}

func (server *Server) updateProperty(ctx *gin.Context) {
	var req updatePropertyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdatePropertyParams{
		ID:           req.ID,
		PropertyName: req.PropertyName,
		Lat:          req.Lat,
		Long:         req.Long,
		BlockID:      req.BlockID,
	}
	property, err := server.store.UpdateProperty(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, property)
}

type deletePropertyRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteProperty(ctx *gin.Context) {
	var req deletePropertyRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteProperty(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
