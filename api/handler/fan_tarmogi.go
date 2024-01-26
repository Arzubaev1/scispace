package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create fan_tarmogi godoc
// @ID create_fan_tarmogi
// @Router /fan_tarmogi [POST]
// @Summary Create FanTarmogi
// @Description Create FanTarmogi
// @Tags FanTarmogi
// @Accept json
// @Procedure json
// @Param FanTarmogi body models.CreateFanTarmogi true "CreateFanTarmogiRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateFanTarmogi(c *gin.Context) {

	var createFanTarmogi models.CreateFanTarmogi
	err := c.ShouldBindJSON(&createFanTarmogi)
	if err != nil {
		h.handlerResponse(c, "error fan_tarmogi should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.FanTarmogi().Create(c.Request.Context(), &createFanTarmogi)
	if err != nil {
		h.handlerResponse(c, "storage.fan_tarmogi.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.FanTarmogi().GetByID(c.Request.Context(), &models.FanTarmogiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.fan_tarmogi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create fan_tarmogi resposne", http.StatusCreated, resp)
}

// GetByID fan_tarmogi godoc
// @ID get_by_id_fan_tarmogi
// @Router /fan_tarmogi/{id} [GET]
// @Summary Get By ID FanTarmogi
// @Description Get By ID FanTarmogi
// @Tags FanTarmogi
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdFanTarmogi(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// FanTarmogiData := val.(helper.TokenInfo)
	// if len(FanTarmogiData) > 0 {
	// 	id = FanTarmogiData.FanTarmogiID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.FanTarmogi().GetByID(c.Request.Context(), &models.FanTarmogiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.fan_tarmogi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id fan_tarmogi resposne", http.StatusOK, resp)
}

// GetList fan_tarmogi godoc
// @ID get_list_fan_tarmogi
// @Router /fan_tarmogi [GET]
// @Summary Get List FanTarmogi
// @Description Get List FanTarmogi
// @Tags FanTarmogi
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListFanTarmogi(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list fan_tarmogi offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list fan_tarmogi limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.FanTarmogi().GetList(c.Request.Context(), &models.FanTarmogiGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.fan_tarmogi.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list fan_tarmogi resposne", http.StatusOK, resp)
}

// Update fan_tarmogi godoc
// @ID update_fan_tarmogi
// @Router /fan_tarmogi/{id} [PUT]
// @Summary Update FanTarmogi
// @Description Update FanTarmogi
// @Tags FanTarmogi
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param FanTarmogi body models.UpdateFanTarmogi true "UpdateFanTarmogiRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateFanTarmogi(c *gin.Context) {

	var (
		id             string = c.Param("id")
		updateFanTarmogi models.UpdateFanTarmogi
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateFanTarmogi)
	if err != nil {
		h.handlerResponse(c, "error fan_tarmogi should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateFanTarmogi.Id = id
	rowsAffected, err := h.strg.FanTarmogi().Update(c.Request.Context(), &updateFanTarmogi)
	if err != nil {
		h.handlerResponse(c, "storage.fan_tarmogi.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.fan_tarmogi.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.FanTarmogi().GetByID(c.Request.Context(), &models.FanTarmogiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.fan_tarmogi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create fan_tarmogi resposne", http.StatusAccepted, resp)
}

// Delete fan_tarmogi godoc
// @ID delete_fan_tarmogi
// @Router /fan_tarmogi/{id} [DELETE]
// @Summary Delete FanTarmogi
// @Description Delete FanTarmogi
// @Tags FanTarmogi
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteFanTarmogi(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.FanTarmogi().Delete(c.Request.Context(), &models.FanTarmogiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.fan_tarmogi.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create fan_tarmogi resposne", http.StatusNoContent, nil)
}
