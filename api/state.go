package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createStateRequest struct {
	StateName string `json:"state_name" binding:"required"`
	StateDesc string `json:"state_desc" binding:"required"`
	CountryID int64  `json:"country_id" binding:"required"`
	AreaID    int64  `json:"area_id" binding:"required"`
}

func (server *Server) createState(ctx *gin.Context) {
	var req createStateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateStateParams{
		StateName: req.StateName,
		StateDesc: req.StateDesc,
		CountryID: req.CountryID,
		AreaID:    req.AreaID,
	}
	state, err := server.store.CreateState(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, state)
}

type getStateRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getState0(ctx *gin.Context) {
	var req getStateRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	state, err := server.store.GetState0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, state)
}

type getStateRequest1 struct {
	StateName string `uri:"state_name" binding:"required,min=1"`
}

func (server *Server) getState1(ctx *gin.Context) {
	var req getStateRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	state, err := server.store.GetState1(ctx, req.StateName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, state)
}

type listStateRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listStates(ctx *gin.Context) {
	var req listStateRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListStatesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	states, err := server.store.ListStates(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, states)
}

type updateStateRequest struct {
	ID        int64  `json:"id" binding:"required"`
	StateName string `json:"state_name" binding:"required"`
	StateDesc string `json:"state_desc" binding:"required"`
	CountryID int64  `json:"country_id" binding:"required"`
	AreaID    int64  `json:"area_id" binding:"required"`
}

func (server *Server) updateState(ctx *gin.Context) {
	var req updateStateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateStateParams{
		ID:        req.ID,
		StateName: req.StateName,
		StateDesc: req.StateDesc,
		CountryID: req.CountryID,
		AreaID:    req.AreaID,
	}
	state, err := server.store.UpdateState(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, state)
}

type deleteStateRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteState(ctx *gin.Context) {
	var req deleteStateRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteState(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
