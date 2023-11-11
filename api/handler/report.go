package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create report godoc
// @ID create_report
// @Router /report [POST]
// @Summary Create Report
// @Description Create Report
// @Tags Report
// @Accept json
// @Procedure json
// @Param Report body models.CreateReport true "CreateReportRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateReport(c *gin.Context) {

	var createReport models.CreateReport
	err := c.ShouldBindJSON(&createReport)
	if err != nil {
		h.handlerResponse(c, "error report should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Report().Create(c.Request.Context(), &createReport)
	if err != nil {
		h.handlerResponse(c, "storage.report.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Report().GetByID(c.Request.Context(), &models.ReportPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.report.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create report resposne", http.StatusCreated, resp)
}

// GetByID report godoc
// @ID get_by_id_report
// @Router /report/{id} [GET]
// @Summary Get By ID Report
// @Description Get By ID Report
// @Tags Report
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdReport(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// ReportData := val.(helper.TokenInfo)
	// if len(ReportData) > 0 {
	// 	id = ReportData.ReportID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Report().GetByID(c.Request.Context(), &models.ReportPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.report.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id report resposne", http.StatusOK, resp)
}

// GetList report godoc
// @ID get_list_report
// @Router /report [GET]
// @Summary Get List Report
// @Description Get List Report
// @Tags Report
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListReport(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list report offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list report limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Report().GetList(c.Request.Context(), &models.ReportGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.report.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list report resposne", http.StatusOK, resp)
}

// Update report godoc
// @ID update_report
// @Router /report/{id} [PUT]
// @Summary Update Report
// @Description Update Report
// @Tags Report
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Report body models.UpdateReport true "UpdateReportRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateReport(c *gin.Context) {

	var (
		id           string = c.Param("id")
		updateReport models.UpdateReport
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateReport)
	if err != nil {
		h.handlerResponse(c, "error report should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateReport.Id = id
	rowsAffected, err := h.strg.Report().Update(c.Request.Context(), &updateReport)
	if err != nil {
		h.handlerResponse(c, "storage.report.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.report.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Report().GetByID(c.Request.Context(), &models.ReportPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.report.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create report resposne", http.StatusAccepted, resp)
}

// Delete report godoc
// @ID delete_report
// @Router /report/{id} [DELETE]
// @Summary Delete Report
// @Description Delete Report
// @Tags Report
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteReport(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Report().Delete(c.Request.Context(), &models.ReportPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.report.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create report resposne", http.StatusNoContent, nil)
}
