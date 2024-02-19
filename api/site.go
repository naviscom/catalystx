package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx/db/sqlc"
)

type createSiteRequest struct {
	SiteName       string    `json:"site_name" binding:"required"`
	SiteNameOld    string    `json:"site_name_old" binding:"required"`
	SiteIDGivin    string    `json:"site_id_givin" binding:"required"`
	SiteIDGivinOld string    `json:"site_id_givin_old" binding:"required"`
	Lac            string    `json:"lac" binding:"required"`
	Rac            string    `json:"rac" binding:"required"`
	Rnc            string    `json:"rnc" binding:"required"`
	SiteOnAirDate  time.Time `json:"site_on_air_date" binding:"required"`
	PropertyID     int64     `json:"property_id" binding:"required"`
	SitetypeID     int64     `json:"sitetype_id" binding:"required"`
	VendorID       int64     `json:"vendor_id" binding:"required"`
}

func (server *Server) createSite(ctx *gin.Context) {
	var req createSiteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSiteParams{
		SiteName:       req.SiteName,
		SiteNameOld:    req.SiteNameOld,
		SiteIDGivin:    req.SiteIDGivin,
		SiteIDGivinOld: req.SiteIDGivinOld,
		Lac:            req.Lac,
		Rac:            req.Rac,
		Rnc:            req.Rnc,
		SiteOnAirDate:  req.SiteOnAirDate,
		PropertyID:     req.PropertyID,
		SitetypeID:     req.SitetypeID,
		VendorID:       req.VendorID,
	}
	site, err := server.store.CreateSite(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, site)
}

type getSiteRequest0 struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getSite0(ctx *gin.Context) {
	var req getSiteRequest0
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	site, err := server.store.GetSite0(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, site)
}

type getSiteRequest1 struct {
	SiteName string `uri:"site_name" binding:"required,min=1"`
}

func (server *Server) getSite1(ctx *gin.Context) {
	var req getSiteRequest1
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	site, err := server.store.GetSite1(ctx, req.SiteName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, site)
}

type listSiteRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listSites(ctx *gin.Context) {
	var req listSiteRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListSitesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	sites, err := server.store.ListSites(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, sites)
}

type updateSiteRequest struct {
	ID             int64     `json:"id" binding:"required"`
	SiteName       string    `json:"site_name" binding:"required"`
	SiteNameOld    string    `json:"site_name_old" binding:"required"`
	SiteIDGivin    string    `json:"site_id_givin" binding:"required"`
	SiteIDGivinOld string    `json:"site_id_givin_old" binding:"required"`
	Lac            string    `json:"lac" binding:"required"`
	Rac            string    `json:"rac" binding:"required"`
	Rnc            string    `json:"rnc" binding:"required"`
	SiteOnAirDate  time.Time `json:"site_on_air_date" binding:"required"`
	PropertyID     int64     `json:"property_id" binding:"required"`
	SitetypeID     int64     `json:"sitetype_id" binding:"required"`
	VendorID       int64     `json:"vendor_id" binding:"required"`
}

func (server *Server) updateSite(ctx *gin.Context) {
	var req updateSiteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateSiteParams{
		ID:             req.ID,
		SiteName:       req.SiteName,
		SiteNameOld:    req.SiteNameOld,
		SiteIDGivin:    req.SiteIDGivin,
		SiteIDGivinOld: req.SiteIDGivinOld,
		Lac:            req.Lac,
		Rac:            req.Rac,
		Rnc:            req.Rnc,
		SiteOnAirDate:  req.SiteOnAirDate,
		PropertyID:     req.PropertyID,
		SitetypeID:     req.SitetypeID,
		VendorID:       req.VendorID,
	}
	site, err := server.store.UpdateSite(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, site)
}

type deleteSiteRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteSite(ctx *gin.Context) {
	var req deleteSiteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteSite(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "record deleted successfully")
}
