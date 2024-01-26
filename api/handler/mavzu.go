package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create mavzu godoc
// @ID create_mavzu
// @Router /mavzu [POST]
// @Summary Create Mavzu
// @Description Create Mavzu
// @Tags Mavzu
// @Accept json
// @Procedure json
// @Param Mavzu body models.CreateMavzu true "CreateMavzuRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateMavzu(c *gin.Context) {

	var createMavzu models.CreateMavzu
	err := c.ShouldBindJSON(&createMavzu)
	if err != nil {
		h.handlerResponse(c, "error mavzu should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Mavzu().Create(c.Request.Context(), &createMavzu)
	if err != nil {
		h.handlerResponse(c, "storage.mavzu.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Mavzu().GetByID(c.Request.Context(), &models.MavzuPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.mavzu.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create mavzu resposne", http.StatusCreated, resp)
}

// GetByID mavzu godoc
// @ID get_by_id_mavzu
// @Router /mavzu/{id} [GET]
// @Summary Get By ID Mavzu
// @Description Get By ID Mavzu
// @Tags Mavzu
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdMavzu(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// MavzuData := val.(helper.TokenInfo)
	// if len(MavzuData) > 0 {
	// 	id = MavzuData.MavzuID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Mavzu().GetByID(c.Request.Context(), &models.MavzuPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.mavzu.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id mavzu resposne", http.StatusOK, resp)
}

// GetList mavzu godoc
// @ID get_list_mavzu
// @Router /mavzu [GET]
// @Summary Get List Mavzu
// @Description Get List Mavzu
// @Tags Mavzu
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListMavzu(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list mavzu offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list mavzu limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Mavzu().GetList(c.Request.Context(), &models.MavzuGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.mavzu.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list mavzu resposne", http.StatusOK, resp)
}

// Update mavzu godoc
// @ID update_mavzu
// @Router /mavzu/{id} [PUT]
// @Summary Update Mavzu
// @Description Update Mavzu
// @Tags Mavzu
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Mavzu body models.UpdateMavzu true "UpdateMavzuRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateMavzu(c *gin.Context) {

	var (
		id             string = c.Param("id")
		updateMavzu models.UpdateMavzu
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateMavzu)
	if err != nil {
		h.handlerResponse(c, "error mavzu should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateMavzu.Id = id
	rowsAffected, err := h.strg.Mavzu().Update(c.Request.Context(), &updateMavzu)
	if err != nil {
		h.handlerResponse(c, "storage.mavzu.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.mavzu.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Mavzu().GetByID(c.Request.Context(), &models.MavzuPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.mavzu.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create mavzu resposne", http.StatusAccepted, resp)
}

// Delete mavzu godoc
// @ID delete_mavzu
// @Router /mavzu/{id} [DELETE]
// @Summary Delete Mavzu
// @Description Delete Mavzu
// @Tags Mavzu
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteMavzu(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Mavzu().Delete(c.Request.Context(), &models.MavzuPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.mavzu.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create mavzu resposne", http.StatusNoContent, nil)
}
