package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createTownRequest struct {
	TownName   string `json:"town_name" binding:"required"`
	TownDesc   string `json:"town_desc" binding:"required"`
	DistrictID int64  `json:"district_id" binding:"required"`
}

func (server *Server) createTown(ctx *gin.Context) {
	var req createTownRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTownParams{
		TownName:   req.TownName,
		TownDesc:   req.TownDesc,
		DistrictID: req.DistrictID,
	}
	town, err := server.store.CreateTown(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, town)
}

type getTownRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTown0(ctx *gin.Context) {
	var req getTownRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	town, err := server.store.GetTown0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, town)
}

type getTownRequest1 struct {
	TownName string `uri:"town_name" binding:"required,min=1"`
}

func (server *Server) getTown1(ctx *gin.Context) {
	var req getTownRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	town, err := server.store.GetTown1(ctx, req.TownName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, town)
}

type listTownRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTowns(ctx *gin.Context) {
	var req listTownRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTownsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	towns, err := server.store.ListTowns(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, towns)
}

type updateTownRequest struct {
	ID         int64  `json:"id" binding:"required"`
	TownName   string `json:"town_name" binding:"required"`
	TownDesc   string `json:"town_desc" binding:"required"`
	DistrictID int64  `json:"district_id" binding:"required"`
}

func (server *Server) updateTown(ctx *gin.Context) {
	var req updateTownRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTownParams{
		ID:         req.ID,
		TownName:   req.TownName,
		TownDesc:   req.TownDesc,
		DistrictID: req.DistrictID,
	}
	town, err := server.store.UpdateTown(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, town)
}

type deleteTownRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteTown(ctx *gin.Context) {
	var req deleteTownRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteTown(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
