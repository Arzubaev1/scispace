package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// @Security ApiKeyAuth
// Create other godoc
// @ID create_other
// @Router /other [POST]
// @Summary Create Other
// @Description Create Other
// @Tags Other
// @Accept json
// @Procedure json
// @Param other body models.CreateOther true "CreateOtherRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateOther(c *gin.Context) {

	var createOther models.CreateOther
	err := c.ShouldBindJSON(&createOther)
	if err != nil {
		h.handlerResponse(c, "error other should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Other().Create(c.Request.Context(), &createOther)
	if err != nil {
		h.handlerResponse(c, "storage.other.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Other().GetByID(c.Request.Context(), &models.OtherPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.other.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create other resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID other godoc
// @ID get_by_id_other
// @Router /other/{id} [GET]
// @Summary Get By ID Other
// @Description Get By ID Other
// @Tags Other
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdOther(c *gin.Context) {
	var id string

	// Here We Check id from Token
	val, exist := c.Get("Auth")
	if !exist {
		h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
		return
	}

	otherData := val.(helper.OtherTokenInfo)
	if len(otherData.OtherID) > 0 {
		id = otherData.OtherID
	} else {
		id = c.Param("id")
	}

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Other().GetByID(c.Request.Context(), &models.OtherPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.other.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id other resposne", http.StatusOK, resp)
}

// GetList other godoc
// @ID get_list_other
// @Router /other [GET]
// @Summary Get List Other
// @Description Get List Other
// @Tags Other
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
func (h *handler) GetListOther(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list other offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list other limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Other().GetList(c.Request.Context(), &models.OtherGetListRequest{
		Offset:           offset,
		Limit:            limit,
		SearchFirstName:  c.Query("search_first_name"),
		SearchLastName:   c.Query("search_last_name"),
		SearchMiddleName: c.Query("search_middle_name"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.other.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list other resposne", http.StatusOK, resp)
}

// Update other godoc
// @ID update_other
// @Router /other/{id} [PUT]
// @Summary Update Other
// @Description Update Other
// @Tags Other
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param other body models.UpdateOther true "UpdateOtherRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateOther(c *gin.Context) {

	var (
		id          string = c.Param("id")
		updateOther models.UpdateOther
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateOther)
	if err != nil {
		h.handlerResponse(c, "error other should bind json", http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.strg.Other().Update(c.Request.Context(), &updateOther)
	if err != nil {
		h.handlerResponse(c, "storage.other.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.other.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Other().GetByID(c.Request.Context(), &models.OtherPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.other.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create other resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete other godoc
// @ID delete_other
// @Router /other/{id} [DELETE]
// @Summary Delete Other
// @Description Delete Other
// @Tags Other
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteOther(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Other().Delete(c.Request.Context(), &models.OtherPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.other.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create other resposne", http.StatusNoContent, nil)
}
