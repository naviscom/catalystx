package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createTrafficRequest struct {
	TrafficDate              time.Time `json:"traffic_date" binding:"required"`
	Avgdailydldatamb         float32   `json:"avgdailydldatamb" binding:"required"`
	Avgdailyuldatamb         float32   `json:"avgdailyuldatamb" binding:"required"`
	Avgdailytotdatamb        float32   `json:"avgdailytotdatamb" binding:"required"`
	Avgdailytotvoicemin      float32   `json:"avgdailytotvoicemin" binding:"required"`
	Avgdailytotvideomin      float32   `json:"avgdailytotvideomin" binding:"required"`
	Qci1Data                 float32   `json:"qci1_data" binding:"required"`
	Qci6Data                 float32   `json:"qci6_data" binding:"required"`
	Qci8Data                 float32   `json:"qci8_data" binding:"required"`
	QciOtherData             float32   `json:"qci_other_data" binding:"required"`
	Avgdailytotvoicemin4g    float32   `json:"avgdailytotvoicemin4g" binding:"required"`
	Avgdailytotvoicemintotal float32   `json:"avgdailytotvoicemintotal" binding:"required"`
	Userdlthroughput         float32   `json:"userdlthroughput" binding:"required"`
	Dlpacketlossrate         float32   `json:"dlpacketlossrate" binding:"required"`
	Overallpsdropcallrate    float32   `json:"overallpsdropcallrate" binding:"required"`
	Bhdldatamb               float32   `json:"bhdldatamb" binding:"required"`
	Bhupdatamb               float32   `json:"bhupdatamb" binding:"required"`
	Bhtotdatamb              float32   `json:"bhtotdatamb" binding:"required"`
	Bhtotvoicemin            float32   `json:"bhtotvoicemin" binding:"required"`
	Bhtotvideomin            float32   `json:"bhtotvideomin" binding:"required"`
	Bhcsusers                float32   `json:"bhcsusers" binding:"required"`
	Bhhsupausers             float32   `json:"bhhsupausers" binding:"required"`
	Bhhsdpausers             float32   `json:"bhhsdpausers" binding:"required"`
	Bhr99uldl                float32   `json:"bhr99uldl" binding:"required"`
	Powercapacity            float32   `json:"powercapacity" binding:"required"`
	Powerutilization         float32   `json:"powerutilization" binding:"required"`
	Codecapacity             float32   `json:"codecapacity" binding:"required"`
	Codeutilization          float32   `json:"codeutilization" binding:"required"`
	Ceulcapacity             float32   `json:"ceulcapacity" binding:"required"`
	Ceulutilization          float32   `json:"ceulutilization" binding:"required"`
	Cedlcapacity             float32   `json:"cedlcapacity" binding:"required"`
	Cedlutilization          float32   `json:"cedlutilization" binding:"required"`
	Iubcapacity              float32   `json:"iubcapacity" binding:"required"`
	Iubutlization            float32   `json:"iubutlization" binding:"required"`
	Bhrrcusers               float32   `json:"bhrrcusers" binding:"required"`
	CellID                   int64     `json:"cell_id" binding:"required"`
}

func (server *Server) createTraffic(ctx *gin.Context) {
	var req createTrafficRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTrafficParams{
		TrafficDate:              req.TrafficDate,
		Avgdailydldatamb:         req.Avgdailydldatamb,
		Avgdailyuldatamb:         req.Avgdailyuldatamb,
		Avgdailytotdatamb:        req.Avgdailytotdatamb,
		Avgdailytotvoicemin:      req.Avgdailytotvoicemin,
		Avgdailytotvideomin:      req.Avgdailytotvideomin,
		Qci1Data:                 req.Qci1Data,
		Qci6Data:                 req.Qci6Data,
		Qci8Data:                 req.Qci8Data,
		QciOtherData:             req.QciOtherData,
		Avgdailytotvoicemin4g:    req.Avgdailytotvoicemin4g,
		Avgdailytotvoicemintotal: req.Avgdailytotvoicemintotal,
		Userdlthroughput:         req.Userdlthroughput,
		Dlpacketlossrate:         req.Dlpacketlossrate,
		Overallpsdropcallrate:    req.Overallpsdropcallrate,
		Bhdldatamb:               req.Bhdldatamb,
		Bhupdatamb:               req.Bhupdatamb,
		Bhtotdatamb:              req.Bhtotdatamb,
		Bhtotvoicemin:            req.Bhtotvoicemin,
		Bhtotvideomin:            req.Bhtotvideomin,
		Bhcsusers:                req.Bhcsusers,
		Bhhsupausers:             req.Bhhsupausers,
		Bhhsdpausers:             req.Bhhsdpausers,
		Bhr99uldl:                req.Bhr99uldl,
		Powercapacity:            req.Powercapacity,
		Powerutilization:         req.Powerutilization,
		Codecapacity:             req.Codecapacity,
		Codeutilization:          req.Codeutilization,
		Ceulcapacity:             req.Ceulcapacity,
		Ceulutilization:          req.Ceulutilization,
		Cedlcapacity:             req.Cedlcapacity,
		Cedlutilization:          req.Cedlutilization,
		Iubcapacity:              req.Iubcapacity,
		Iubutlization:            req.Iubutlization,
		Bhrrcusers:               req.Bhrrcusers,
		CellID:                   req.CellID,
	}
	traffic, err := server.store.CreateTraffic(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, traffic)
}

type getTrafficRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTraffic0(ctx *gin.Context) {
	var req getTrafficRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	traffic, err := server.store.GetTraffic0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, traffic)
}

type listTrafficRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTraffic(ctx *gin.Context) {
	var req listTrafficRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTrafficParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	traffic, err := server.store.ListTraffic(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, traffic)
}

type updateTrafficRequest struct {
	ID                       int64     `json:"id" binding:"required"`
	TrafficDate              time.Time `json:"traffic_date" binding:"required"`
	Avgdailydldatamb         float32   `json:"avgdailydldatamb" binding:"required"`
	Avgdailyuldatamb         float32   `json:"avgdailyuldatamb" binding:"required"`
	Avgdailytotdatamb        float32   `json:"avgdailytotdatamb" binding:"required"`
	Avgdailytotvoicemin      float32   `json:"avgdailytotvoicemin" binding:"required"`
	Avgdailytotvideomin      float32   `json:"avgdailytotvideomin" binding:"required"`
	Qci1Data                 float32   `json:"qci1_data" binding:"required"`
	Qci6Data                 float32   `json:"qci6_data" binding:"required"`
	Qci8Data                 float32   `json:"qci8_data" binding:"required"`
	QciOtherData             float32   `json:"qci_other_data" binding:"required"`
	Avgdailytotvoicemin4g    float32   `json:"avgdailytotvoicemin4g" binding:"required"`
	Avgdailytotvoicemintotal float32   `json:"avgdailytotvoicemintotal" binding:"required"`
	Userdlthroughput         float32   `json:"userdlthroughput" binding:"required"`
	Dlpacketlossrate         float32   `json:"dlpacketlossrate" binding:"required"`
	Overallpsdropcallrate    float32   `json:"overallpsdropcallrate" binding:"required"`
	Bhdldatamb               float32   `json:"bhdldatamb" binding:"required"`
	Bhupdatamb               float32   `json:"bhupdatamb" binding:"required"`
	Bhtotdatamb              float32   `json:"bhtotdatamb" binding:"required"`
	Bhtotvoicemin            float32   `json:"bhtotvoicemin" binding:"required"`
	Bhtotvideomin            float32   `json:"bhtotvideomin" binding:"required"`
	Bhcsusers                float32   `json:"bhcsusers" binding:"required"`
	Bhhsupausers             float32   `json:"bhhsupausers" binding:"required"`
	Bhhsdpausers             float32   `json:"bhhsdpausers" binding:"required"`
	Bhr99uldl                float32   `json:"bhr99uldl" binding:"required"`
	Powercapacity            float32   `json:"powercapacity" binding:"required"`
	Powerutilization         float32   `json:"powerutilization" binding:"required"`
	Codecapacity             float32   `json:"codecapacity" binding:"required"`
	Codeutilization          float32   `json:"codeutilization" binding:"required"`
	Ceulcapacity             float32   `json:"ceulcapacity" binding:"required"`
	Ceulutilization          float32   `json:"ceulutilization" binding:"required"`
	Cedlcapacity             float32   `json:"cedlcapacity" binding:"required"`
	Cedlutilization          float32   `json:"cedlutilization" binding:"required"`
	Iubcapacity              float32   `json:"iubcapacity" binding:"required"`
	Iubutlization            float32   `json:"iubutlization" binding:"required"`
	Bhrrcusers               float32   `json:"bhrrcusers" binding:"required"`
	CellID                   int64     `json:"cell_id" binding:"required"`
}

func (server *Server) updateTraffic(ctx *gin.Context) {
	var req updateTrafficRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTrafficParams{
		ID:                       req.ID,
		TrafficDate:              req.TrafficDate,
		Avgdailydldatamb:         req.Avgdailydldatamb,
		Avgdailyuldatamb:         req.Avgdailyuldatamb,
		Avgdailytotdatamb:        req.Avgdailytotdatamb,
		Avgdailytotvoicemin:      req.Avgdailytotvoicemin,
		Avgdailytotvideomin:      req.Avgdailytotvideomin,
		Qci1Data:                 req.Qci1Data,
		Qci6Data:                 req.Qci6Data,
		Qci8Data:                 req.Qci8Data,
		QciOtherData:             req.QciOtherData,
		Avgdailytotvoicemin4g:    req.Avgdailytotvoicemin4g,
		Avgdailytotvoicemintotal: req.Avgdailytotvoicemintotal,
		Userdlthroughput:         req.Userdlthroughput,
		Dlpacketlossrate:         req.Dlpacketlossrate,
		Overallpsdropcallrate:    req.Overallpsdropcallrate,
		Bhdldatamb:               req.Bhdldatamb,
		Bhupdatamb:               req.Bhupdatamb,
		Bhtotdatamb:              req.Bhtotdatamb,
		Bhtotvoicemin:            req.Bhtotvoicemin,
		Bhtotvideomin:            req.Bhtotvideomin,
		Bhcsusers:                req.Bhcsusers,
		Bhhsupausers:             req.Bhhsupausers,
		Bhhsdpausers:             req.Bhhsdpausers,
		Bhr99uldl:                req.Bhr99uldl,
		Powercapacity:            req.Powercapacity,
		Powerutilization:         req.Powerutilization,
		Codecapacity:             req.Codecapacity,
		Codeutilization:          req.Codeutilization,
		Ceulcapacity:             req.Ceulcapacity,
		Ceulutilization:          req.Ceulutilization,
		Cedlcapacity:             req.Cedlcapacity,
		Cedlutilization:          req.Cedlutilization,
		Iubcapacity:              req.Iubcapacity,
		Iubutlization:            req.Iubutlization,
		Bhrrcusers:               req.Bhrrcusers,
		CellID:                   req.CellID,
	}
	traffic, err := server.store.UpdateTraffic(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, traffic)
}

type deleteTrafficRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteTraffic(ctx *gin.Context) {
	var req deleteTrafficRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteTraffic(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
