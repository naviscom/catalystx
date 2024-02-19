package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createAreaRequest struct {
	AreaName string `json:"area_name" binding:"required"`
	AreaDesc string `json:"area_desc" binding:"required"`
}

func (server *Server) createArea(ctx *gin.Context) {
	var req createAreaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAreaParams{
		AreaName: req.AreaName,
		AreaDesc: req.AreaDesc,
	}
	area, err := server.store.CreateArea(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, area)
}

type getAreaRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getArea0(ctx *gin.Context) {
	var req getAreaRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	area, err := server.store.GetArea0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, area)
}

type getAreaRequest1 struct {
	AreaName string `uri:"area_name" binding:"required,min=1"`
}

func (server *Server) getArea1(ctx *gin.Context) {
	var req getAreaRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	area, err := server.store.GetArea1(ctx, req.AreaName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, area)
}

type listAreaRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAreas(ctx *gin.Context) {
	var req listAreaRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAreasParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	areas, err := server.store.ListAreas(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, areas)
}

type updateAreaRequest struct {
	ID       int64  `json:"id" binding:"required"`
	AreaName string `json:"area_name" binding:"required"`
	AreaDesc string `json:"area_desc" binding:"required"`
}

func (server *Server) updateArea(ctx *gin.Context) {
	var req updateAreaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAreaParams{
		ID:       req.ID,
		AreaName: req.AreaName,
		AreaDesc: req.AreaDesc,
	}
	area, err := server.store.UpdateArea(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, area)
}

type deleteAreaRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteArea(ctx *gin.Context) {
	var req deleteAreaRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteArea(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
