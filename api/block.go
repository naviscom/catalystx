package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createBlockRequest struct {
	BlockName       string `json:"block_name" binding:"required"`
	BlockDesc       string `json:"block_desc" binding:"required"`
	TotalPopulation int64  `json:"total_population" binding:"required"`
	TownID          int64  `json:"town_id" binding:"required"`
	ClutterID       int64  `json:"clutter_id" binding:"required"`
}

func (server *Server) createBlock(ctx *gin.Context) {
	var req createBlockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBlockParams{
		BlockName:       req.BlockName,
		BlockDesc:       req.BlockDesc,
		TotalPopulation: req.TotalPopulation,
		TownID:          req.TownID,
		ClutterID:       req.ClutterID,
	}
	block, err := server.store.CreateBlock(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, block)
}

type getBlockRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBlock0(ctx *gin.Context) {
	var req getBlockRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	block, err := server.store.GetBlock0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, block)
}

type getBlockRequest1 struct {
	BlockName string `uri:"block_name" binding:"required,min=1"`
}

func (server *Server) getBlock1(ctx *gin.Context) {
	var req getBlockRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	block, err := server.store.GetBlock1(ctx, req.BlockName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, block)
}

type listBlockRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listBlocks(ctx *gin.Context) {
	var req listBlockRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListBlocksParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	blocks, err := server.store.ListBlocks(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, blocks)
}

type updateBlockRequest struct {
	ID              int64  `json:"id" binding:"required"`
	BlockName       string `json:"block_name" binding:"required"`
	BlockDesc       string `json:"block_desc" binding:"required"`
	TotalPopulation int64  `json:"total_population" binding:"required"`
	TownID          int64  `json:"town_id" binding:"required"`
	ClutterID       int64  `json:"clutter_id" binding:"required"`
}

func (server *Server) updateBlock(ctx *gin.Context) {
	var req updateBlockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateBlockParams{
		ID:              req.ID,
		BlockName:       req.BlockName,
		BlockDesc:       req.BlockDesc,
		TotalPopulation: req.TotalPopulation,
		TownID:          req.TownID,
		ClutterID:       req.ClutterID,
	}
	block, err := server.store.UpdateBlock(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, block)
}

type deleteBlockRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteBlock(ctx *gin.Context) {
	var req deleteBlockRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteBlock(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
