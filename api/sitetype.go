package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createSitetypeRequest struct {
	TypeName string `json:"type_name" binding:"required"`
	TypeDesc string `json:"type_desc" binding:"required"`
}

func (server *Server) createSitetype(ctx *gin.Context) {
	var req createSitetypeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSitetypeParams{
		TypeName: req.TypeName,
		TypeDesc: req.TypeDesc,
	}
	sitetype, err := server.store.CreateSitetype(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sitetype)
}

type getSitetypeRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getSitetype0(ctx *gin.Context) {
	var req getSitetypeRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	sitetype, err := server.store.GetSitetype0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sitetype)
}

type getSitetypeRequest1 struct {
	TypeName string `uri:"type_name" binding:"required,min=1"`
}

func (server *Server) getSitetype1(ctx *gin.Context) {
	var req getSitetypeRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	sitetype, err := server.store.GetSitetype1(ctx, req.TypeName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sitetype)
}

type listSitetypeRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listSitetypes(ctx *gin.Context) {
	var req listSitetypeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListSitetypesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	sitetypes, err := server.store.ListSitetypes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sitetypes)
}

type updateSitetypeRequest struct {
	ID       int64  `json:"id" binding:"required"`
	TypeName string `json:"type_name" binding:"required"`
	TypeDesc string `json:"type_desc" binding:"required"`
}

func (server *Server) updateSitetype(ctx *gin.Context) {
	var req updateSitetypeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateSitetypeParams{
		ID:       req.ID,
		TypeName: req.TypeName,
		TypeDesc: req.TypeDesc,
	}
	sitetype, err := server.store.UpdateSitetype(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sitetype)
}

type deleteSitetypeRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteSitetype(ctx *gin.Context) {
	var req deleteSitetypeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteSitetype(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
