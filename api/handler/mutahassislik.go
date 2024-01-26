package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create mutahassislik godoc
// @ID create_mutahassislik
// @Router /mutahassislik [POST]
// @Summary Create Mutahassislik
// @Description Create Mutahassislik
// @Tags Mutahassislik
// @Accept json
// @Procedure json
// @Param Mutahassislik body models.CreateMutahassislik true "CreateMutahassislikRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateMutahassislik(c *gin.Context) {

	var createMutahassislik models.CreateMutahassislik
	err := c.ShouldBindJSON(&createMutahassislik)
	if err != nil {
		h.handlerResponse(c, "error mutahassislik should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Mutahassislik().Create(c.Request.Context(), &createMutahassislik)
	if err != nil {
		h.handlerResponse(c, "storage.mutahassislik.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Mutahassislik().GetByID(c.Request.Context(), &models.MutahassislikPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.mutahassislik.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create mutahassislik resposne", http.StatusCreated, resp)
}

// GetByID mutahassislik godoc
// @ID get_by_id_mutahassislik
// @Router /mutahassislik/{id} [GET]
// @Summary Get By ID Mutahassislik
// @Description Get By ID Mutahassislik
// @Tags Mutahassislik
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdMutahassislik(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// MutahassislikData := val.(helper.TokenInfo)
	// if len(MutahassislikData) > 0 {
	// 	id = MutahassislikData.MutahassislikID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Mutahassislik().GetByID(c.Request.Context(), &models.MutahassislikPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.mutahassislik.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id mutahassislik resposne", http.StatusOK, resp)
}

// GetList mutahassislik godoc
// @ID get_list_mutahassislik
// @Router /mutahassislik [GET]
// @Summary Get List Mutahassislik
// @Description Get List Mutahassislik
// @Tags Mutahassislik
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListMutahassislik(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list mutahassislik offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list mutahassislik limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Mutahassislik().GetList(c.Request.Context(), &models.MutahassislikGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.mutahassislik.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list mutahassislik resposne", http.StatusOK, resp)
}

// Update mutahassislik godoc
// @ID update_mutahassislik
// @Router /mutahassislik/{id} [PUT]
// @Summary Update Mutahassislik
// @Description Update Mutahassislik
// @Tags Mutahassislik
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Mutahassislik body models.UpdateMutahassislik true "UpdateMutahassislikRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateMutahassislik(c *gin.Context) {

	var (
		id                  string = c.Param("id")
		updateMutahassislik models.UpdateMutahassislik
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateMutahassislik)
	if err != nil {
		h.handlerResponse(c, "error mutahassislik should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateMutahassislik.Id = id
	rowsAffected, err := h.strg.Mutahassislik().Update(c.Request.Context(), &updateMutahassislik)
	if err != nil {
		h.handlerResponse(c, "storage.mutahassislik.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.mutahassislik.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Mutahassislik().GetByID(c.Request.Context(), &models.MutahassislikPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.mutahassislik.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create mutahassislik resposne", http.StatusAccepted, resp)
}

// Delete mutahassislik godoc
// @ID delete_mutahassislik
// @Router /mutahassislik/{id} [DELETE]
// @Summary Delete Mutahassislik
// @Description Delete Mutahassislik
// @Tags Mutahassislik
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteMutahassislik(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Mutahassislik().Delete(c.Request.Context(), &models.MutahassislikPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.mutahassislik.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create mutahassislik resposne", http.StatusNoContent, nil)
}
