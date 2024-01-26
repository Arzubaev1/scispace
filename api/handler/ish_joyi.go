package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create ish_joyi godoc
// @ID create_ish_joyi
// @Router /ish_joyi [POST]
// @Summary Create IshJoyi
// @Description Create IshJoyi
// @Tags IshJoyi
// @Accept json
// @Procedure json
// @Param IshJoyi body models.CreateIshJoyi true "CreateIshJoyiRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateIshJoyi(c *gin.Context) {

	var createIshJoyi models.CreateIshJoyi
	err := c.ShouldBindJSON(&createIshJoyi)
	if err != nil {
		h.handlerResponse(c, "error ish_joyi should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.IshJoyi().Create(c.Request.Context(), &createIshJoyi)
	if err != nil {
		h.handlerResponse(c, "storage.ish_joyi.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.IshJoyi().GetByID(c.Request.Context(), &models.IshJoyiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.ish_joyi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create ish_joyi resposne", http.StatusCreated, resp)
}

// GetByID ish_joyi godoc
// @ID get_by_id_ish_joyi
// @Router /ish_joyi/{id} [GET]
// @Summary Get By ID IshJoyi
// @Description Get By ID IshJoyi
// @Tags IshJoyi
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdIshJoyi(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// IshJoyiData := val.(helper.TokenInfo)
	// if len(IshJoyiData) > 0 {
	// 	id = IshJoyiData.IshJoyiID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.IshJoyi().GetByID(c.Request.Context(), &models.IshJoyiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.ish_joyi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id ish_joyi resposne", http.StatusOK, resp)
}

// GetList ish_joyi godoc
// @ID get_list_ish_joyi
// @Router /ish_joyi [GET]
// @Summary Get List IshJoyi
// @Description Get List IshJoyi
// @Tags IshJoyi
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListIshJoyi(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list ish_joyi offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list ish_joyi limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.IshJoyi().GetList(c.Request.Context(), &models.IshJoyiGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.ish_joyi.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list ish_joyi resposne", http.StatusOK, resp)
}

// Update ish_joyi godoc
// @ID update_ish_joyi
// @Router /ish_joyi/{id} [PUT]
// @Summary Update IshJoyi
// @Description Update IshJoyi
// @Tags IshJoyi
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param IshJoyi body models.UpdateIshJoyi true "UpdateIshJoyiRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateIshJoyi(c *gin.Context) {

	var (
		id             string = c.Param("id")
		updateIshJoyi models.UpdateIshJoyi
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateIshJoyi)
	if err != nil {
		h.handlerResponse(c, "error ish_joyi should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateIshJoyi.Id = id
	rowsAffected, err := h.strg.IshJoyi().Update(c.Request.Context(), &updateIshJoyi)
	if err != nil {
		h.handlerResponse(c, "storage.ish_joyi.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.ish_joyi.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.IshJoyi().GetByID(c.Request.Context(), &models.IshJoyiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.ish_joyi.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create ish_joyi resposne", http.StatusAccepted, resp)
}

// Delete ish_joyi godoc
// @ID delete_ish_joyi
// @Router /ish_joyi/{id} [DELETE]
// @Summary Delete IshJoyi
// @Description Delete IshJoyi
// @Tags IshJoyi
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteIshJoyi(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.IshJoyi().Delete(c.Request.Context(), &models.IshJoyiPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.ish_joyi.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create ish_joyi resposne", http.StatusNoContent, nil)
}
