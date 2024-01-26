package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create maqola godoc
// @ID create_maqola
// @Router /maqola [POST]
// @Summary Create Maqola
// @Description Create Maqola
// @Tags Maqola
// @Accept json
// @Procedure json
// @Param Maqola body models.CreateMaqola true "CreateMaqolaRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateMaqola(c *gin.Context) {

	var createMaqola models.CreateMaqola
	err := c.ShouldBindJSON(&createMaqola)
	if err != nil {
		h.handlerResponse(c, "error maqola should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Maqola().Create(c.Request.Context(), &createMaqola)
	if err != nil {
		h.handlerResponse(c, "storage.maqola.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Maqola().GetByID(c.Request.Context(), &models.MaqolaPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.maqola.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create maqola resposne", http.StatusCreated, resp)
}

// GetByID maqola godoc
// @ID get_by_id_maqola
// @Router /maqola/{id} [GET]
// @Summary Get By ID Maqola
// @Description Get By ID Maqola
// @Tags Maqola
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdMaqola(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// MaqolaData := val.(helper.TokenInfo)
	// if len(MaqolaData) > 0 {
	// 	id = MaqolaData.MaqolaID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Maqola().GetByID(c.Request.Context(), &models.MaqolaPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.maqola.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id maqola resposne", http.StatusOK, resp)
}

// GetList maqola godoc
// @ID get_list_maqola
// @Router /maqola [GET]
// @Summary Get List Maqola
// @Description Get List Maqola
// @Tags Maqola
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListMaqola(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list maqola offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list maqola limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Maqola().GetList(c.Request.Context(), &models.MaqolaGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.maqola.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list maqola resposne", http.StatusOK, resp)
}

// Update maqola godoc
// @ID update_maqola
// @Router /maqola/{id} [PUT]
// @Summary Update Maqola
// @Description Update Maqola
// @Tags Maqola
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Maqola body models.UpdateMaqola true "UpdateMaqolaRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateMaqola(c *gin.Context) {

	var (
		id           string = c.Param("id")
		updateMaqola models.UpdateMaqola
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateMaqola)
	if err != nil {
		h.handlerResponse(c, "error maqola should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateMaqola.Id = id
	rowsAffected, err := h.strg.Maqola().Update(c.Request.Context(), &updateMaqola)
	if err != nil {
		h.handlerResponse(c, "storage.maqola.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.maqola.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Maqola().GetByID(c.Request.Context(), &models.MaqolaPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.maqola.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create maqola resposne", http.StatusAccepted, resp)
}

// Delete maqola godoc
// @ID delete_maqola
// @Router /maqola/{id} [DELETE]
// @Summary Delete Maqola
// @Description Delete Maqola
// @Tags Maqola
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteMaqola(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Maqola().Delete(c.Request.Context(), &models.MaqolaPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.maqola.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create maqola resposne", http.StatusNoContent, nil)
}
