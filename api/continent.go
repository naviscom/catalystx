package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createContinentRequest struct {
	ContinentName string `json:"continent_name" binding:"required"`
	ContinentDesc string `json:"continent_desc" binding:"required"`
}

func (server *Server) createContinent(ctx *gin.Context) {
	var req createContinentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateContinentParams{
		ContinentName: req.ContinentName,
		ContinentDesc: req.ContinentDesc,
	}
	continent, err := server.store.CreateContinent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, continent)
}

type getContinentRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getContinent0(ctx *gin.Context) {
	var req getContinentRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	continent, err := server.store.GetContinent0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, continent)
}

type getContinentRequest1 struct {
	ContinentName string `uri:"continent_name" binding:"required,min=1"`
}

func (server *Server) getContinent1(ctx *gin.Context) {
	var req getContinentRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	continent, err := server.store.GetContinent1(ctx, req.ContinentName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, continent)
}

type listContinentRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listContinents(ctx *gin.Context) {
	var req listContinentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListContinentsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	continents, err := server.store.ListContinents(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, continents)
}

type updateContinentRequest struct {
	ID            int64  `json:"id" binding:"required"`
	ContinentName string `json:"continent_name" binding:"required"`
	ContinentDesc string `json:"continent_desc" binding:"required"`
}

func (server *Server) updateContinent(ctx *gin.Context) {
	var req updateContinentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateContinentParams{
		ID:            req.ID,
		ContinentName: req.ContinentName,
		ContinentDesc: req.ContinentDesc,
	}
	continent, err := server.store.UpdateContinent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, continent)
}

type deleteContinentRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteContinent(ctx *gin.Context) {
	var req deleteContinentRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteContinent(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
