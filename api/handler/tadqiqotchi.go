package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// @Security ApiKeyAuth
// Create tadqiqotchi godoc
// @ID create_tadqiqotchi
// @Router /tadqiqotchi [POST]
// @Summary Create Tadqiqotchi
// @Description Create Tadqiqotchi
// @Tags Tadqiqotchi
// @Accept json
// @Procedure json
// @Param tadqiqotchi body models.CreateTadqiqotchi true "CreateTadqiqotchiRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateTadqiqotchi(c *gin.Context) {

	var createTadqiqotchi models.CreateTadqiqotchi
	err := c.ShouldBindJSON(&createTadqiqotchi)
	if err != nil {
		h.handlerResponse(c, "error tadqiqotchi should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Tadqiqotchi().Create(c.Request.Context(), &createTadqiqotchi)
	if err != nil {
		h.handlerResponse(c, "storage.tadqiqotchi.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Tadqiqotchi().GetByID(c.Request.Context(), &models.TadqiqotchiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.tadqiqotchi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create tadqiqotchi resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID tadqiqotchi godoc
// @ID get_by_id_tadqiqotchi
// @Router /tadqiqotchi/{id} [GET]
// @Summary Get By ID Tadqiqotchi
// @Description Get By ID Tadqiqotchi
// @Tags Tadqiqotchi
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdTadqiqotchi(c *gin.Context) {
	var id string

	// Here We Check id from Token
	val, exist := c.Get("Auth")
	if !exist {
		h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
		return
	}

	tadqiqotchiData := val.(helper.TadqiqotchiTokenInfo)
	if len(tadqiqotchiData.TadqiqotchiID) > 0 {
		id = tadqiqotchiData.TadqiqotchiID
	} else {
		id = c.Param("id")
	}

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Tadqiqotchi().GetByID(c.Request.Context(), &models.TadqiqotchiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.tadqiqotchi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id tadqiqotchi resposne", http.StatusOK, resp)
}

// GetList tadqiqotchi godoc
// @ID get_list_tadqiqotchi
// @Router /tadqiqotchi [GET]
// @Summary Get List Tadqiqotchi
// @Description Get List Tadqiqotchi
// @Tags Tadqiqotchi
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
func (h *handler) GetListTadqiqotchi(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list tadqiqotchi offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list tadqiqotchi limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Tadqiqotchi().GetList(c.Request.Context(), &models.TadqiqotchiGetListRequest{
		Offset:           offset,
		Limit:            limit,
		SearchFirstName:  c.Query("search_first_name"),
		SearchLastName:   c.Query("search_last_name"),
		SearchMiddleName: c.Query("search_middle_name"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.tadqiqotchi.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list tadqiqotchi resposne", http.StatusOK, resp)
}

// Update tadqiqotchi godoc
// @ID update_tadqiqotchi
// @Router /tadqiqotchi/{id} [PUT]
// @Summary Update Tadqiqotchi
// @Description Update Tadqiqotchi
// @Tags Tadqiqotchi
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param tadqiqotchi body models.UpdateTadqiqotchi true "UpdateTadqiqotchiRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateTadqiqotchi(c *gin.Context) {

	var (
		id                string = c.Param("id")
		updateTadqiqotchi models.UpdateTadqiqotchi
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateTadqiqotchi)
	if err != nil {
		h.handlerResponse(c, "error tadqiqotchi should bind json", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.strg.Tadqiqotchi().Update(c.Request.Context(), &updateTadqiqotchi)
	if err != nil {
		h.handlerResponse(c, "storage.tadqiqotchi.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.tadqiqotchi.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Tadqiqotchi().GetByID(c.Request.Context(), &models.TadqiqotchiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.tadqiqotchi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create tadqiqotchi resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete tadqiqotchi godoc
// @ID delete_tadqiqotchi
// @Router /tadqiqotchi/{id} [DELETE]
// @Summary Delete Tadqiqotchi
// @Description Delete Tadqiqotchi
// @Tags Tadqiqotchi
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteTadqiqotchi(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Tadqiqotchi().Delete(c.Request.Context(), &models.TadqiqotchiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.tadqiqotchi.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create tadqiqotchi resposne", http.StatusNoContent, nil)
}
