package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// @Security ApiKeyAuth
// Create oqituvchi godoc
// @ID create_oqituvchi
// @Router /oqituvchi [POST]
// @Summary Create Oqituvchi
// @Description Create Oqituvchi
// @Tags Oqituvchi
// @Accept json
// @Procedure json
// @Param oqituvchi body models.CreateOqituvchi true "CreateOqituvchiRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateOqituvchi(c *gin.Context) {

	var createOqituvchi models.CreateOqituvchi
	err := c.ShouldBindJSON(&createOqituvchi)
	if err != nil {
		h.handlerResponse(c, "error oqituvchi should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Oqituvchi().Create(c.Request.Context(), &createOqituvchi)
	if err != nil {
		h.handlerResponse(c, "storage.oqituvchi.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Oqituvchi().GetByID(c.Request.Context(), &models.OqituvchiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.oqituvchi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create oqituvchi resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID oqituvchi godoc
// @ID get_by_id_oqituvchi
// @Router /oqituvchi/{id} [GET]
// @Summary Get By ID Oqituvchi
// @Description Get By ID Oqituvchi
// @Tags Oqituvchi
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdOqituvchi(c *gin.Context) {
	var id string

	// Here We Check id from Token
	val, exist := c.Get("Auth")
	if !exist {
		h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
		return
	}

	oqituvchiData := val.(helper.OqituvchiTokenInfo)
	if len(oqituvchiData.OqituvchiID) > 0 {
		id = oqituvchiData.OqituvchiID
	} else {
		id = c.Param("id")
	}

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Oqituvchi().GetByID(c.Request.Context(), &models.OqituvchiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.oqituvchi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id oqituvchi resposne", http.StatusOK, resp)
}

// GetList oqituvchi godoc
// @ID get_list_oqituvchi
// @Router /oqituvchi [GET]
// @Summary Get List Oqituvchi
// @Description Get List Oqituvchi
// @Tags Oqituvchi
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search_first_name query string false "search_first_name"
// @Param search_last_name query string false "search_last_name"
// @Param search_middle_name query string false "search_middle_name"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListOqituvchi(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list oqituvchi offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list oqituvchi limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Oqituvchi().GetList(c.Request.Context(), &models.OqituvchiGetListRequest{
		Offset:           offset,
		Limit:            limit,
		SearchFirstName:  c.Query("search_first_name"),
		SearchLastName:   c.Query("search_last_name"),
		SearchMiddleName: c.Query("search_middle_name"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.oqituvchi.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list oqituvchi resposne", http.StatusOK, resp)
}

// Update oqituvchi godoc
// @ID update_oqituvchi
// @Router /oqituvchi/{id} [PUT]
// @Summary Update Oqituvchi
// @Description Update Oqituvchi
// @Tags Oqituvchi
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param oqituvchi body models.UpdateOqituvchi true "UpdateOqituvchiRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateOqituvchi(c *gin.Context) {

	var (
		id              string = c.Param("id")
		updateOqituvchi models.UpdateOqituvchi
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateOqituvchi)
	if err != nil {
		h.handlerResponse(c, "error oqituvchi should bind json", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.strg.Oqituvchi().Update(c.Request.Context(), &updateOqituvchi)
	if err != nil {
		h.handlerResponse(c, "storage.oqituvchi.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.oqituvchi.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Oqituvchi().GetByID(c.Request.Context(), &models.OqituvchiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.oqituvchi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create oqituvchi resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete oqituvchi godoc
// @ID delete_oqituvchi
// @Router /oqituvchi/{id} [DELETE]
// @Summary Delete Oqituvchi
// @Description Delete Oqituvchi
// @Tags Oqituvchi
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteOqituvchi(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Oqituvchi().Delete(c.Request.Context(), &models.OqituvchiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.oqituvchi.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create oqituvchi resposne", http.StatusNoContent, nil)
}
