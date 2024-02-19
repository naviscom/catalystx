package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createServiceareatypeRequest struct {
	ServiceareatypeName string `json:"serviceareatype_name" binding:"required"`
	ServiceareatypeDesc string `json:"serviceareatype_desc" binding:"required"`
}

func (server *Server) createServiceareatype(ctx *gin.Context) {
	var req createServiceareatypeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateServiceareatypeParams{
		ServiceareatypeName: req.ServiceareatypeName,
		ServiceareatypeDesc: req.ServiceareatypeDesc,
	}
	serviceareatype, err := server.store.CreateServiceareatype(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, serviceareatype)
}

type getServiceareatypeRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getServiceareatype0(ctx *gin.Context) {
	var req getServiceareatypeRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	serviceareatype, err := server.store.GetServiceareatype0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, serviceareatype)
}

type getServiceareatypeRequest1 struct {
	ServiceareatypeName string `uri:"serviceareatype_name" binding:"required,min=1"`
}

func (server *Server) getServiceareatype1(ctx *gin.Context) {
	var req getServiceareatypeRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	serviceareatype, err := server.store.GetServiceareatype1(ctx, req.ServiceareatypeName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, serviceareatype)
}

type listServiceareatypeRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listServiceareatypes(ctx *gin.Context) {
	var req listServiceareatypeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListServiceareatypesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	serviceareatypes, err := server.store.ListServiceareatypes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, serviceareatypes)
}

type updateServiceareatypeRequest struct {
	ID                  int64  `json:"id" binding:"required"`
	ServiceareatypeName string `json:"serviceareatype_name" binding:"required"`
	ServiceareatypeDesc string `json:"serviceareatype_desc" binding:"required"`
}

func (server *Server) updateServiceareatype(ctx *gin.Context) {
	var req updateServiceareatypeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateServiceareatypeParams{
		ID:                  req.ID,
		ServiceareatypeName: req.ServiceareatypeName,
		ServiceareatypeDesc: req.ServiceareatypeDesc,
	}
	serviceareatype, err := server.store.UpdateServiceareatype(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, serviceareatype)
}

type deleteServiceareatypeRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteServiceareatype(ctx *gin.Context) {
	var req deleteServiceareatypeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteServiceareatype(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
