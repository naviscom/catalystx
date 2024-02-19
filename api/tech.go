package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createTechRequest struct {
	TechName string `json:"tech_name" binding:"required"`
	TechDesc string `json:"tech_desc" binding:"required"`
}

func (server *Server) createTech(ctx *gin.Context) {
	var req createTechRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTechParams{
		TechName: req.TechName,
		TechDesc: req.TechDesc,
	}
	tech, err := server.store.CreateTech(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tech)
}

type getTechRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTech0(ctx *gin.Context) {
	var req getTechRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tech, err := server.store.GetTech0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tech)
}

type getTechRequest1 struct {
	TechName string `uri:"tech_name" binding:"required,min=1"`
}

func (server *Server) getTech1(ctx *gin.Context) {
	var req getTechRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tech, err := server.store.GetTech1(ctx, req.TechName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tech)
}

type listTechRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTechs(ctx *gin.Context) {
	var req listTechRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTechsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	techs, err := server.store.ListTechs(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, techs)
}

type updateTechRequest struct {
	ID       int64  `json:"id" binding:"required"`
	TechName string `json:"tech_name" binding:"required"`
	TechDesc string `json:"tech_desc" binding:"required"`
}

func (server *Server) updateTech(ctx *gin.Context) {
	var req updateTechRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTechParams{
		ID:       req.ID,
		TechName: req.TechName,
		TechDesc: req.TechDesc,
	}
	tech, err := server.store.UpdateTech(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, tech)
}

type deleteTechRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteTech(ctx *gin.Context) {
	var req deleteTechRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteTech(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
