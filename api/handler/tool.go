package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create tool godoc
// @ID create_tool
// @Router /tool [POST]
// @Summary Create Tool
// @Description Create Tool
// @Tags Tool
// @Accept json
// @Procedure json
// @Param Tool body models.CreateTool true "CreateToolRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateTool(c *gin.Context) {

	var createTool models.CreateTool
	err := c.ShouldBindJSON(&createTool)
	if err != nil {
		h.handlerResponse(c, "error tool should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Tool().Create(c.Request.Context(), &createTool)
	if err != nil {
		h.handlerResponse(c, "storage.tool.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Tool().GetByID(c.Request.Context(), &models.ToolPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.tool.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create tool resposne", http.StatusCreated, resp)
}

// GetByID tool godoc
// @ID get_by_id_tool
// @Router /tool/{id} [GET]
// @Summary Get By ID Tool
// @Description Get By ID Tool
// @Tags Tool
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdTool(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// ToolData := val.(helper.TokenInfo)
	// if len(ToolData) > 0 {
	// 	id = ToolData.ToolID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Tool().GetByID(c.Request.Context(), &models.ToolPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.tool.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id tool resposne", http.StatusOK, resp)
}

// GetList tool godoc
// @ID get_list_tool
// @Router /tool [GET]
// @Summary Get List Tool
// @Description Get List Tool
// @Tags Tool
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListTool(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list tool offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list tool limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Tool().GetList(c.Request.Context(), &models.ToolGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.tool.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list tool resposne", http.StatusOK, resp)
}

// Update tool godoc
// @ID update_tool
// @Router /tool/{id} [PUT]
// @Summary Update Tool
// @Description Update Tool
// @Tags Tool
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Tool body models.UpdateTool true "UpdateToolRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateTool(c *gin.Context) {

	var (
		id         string = c.Param("id")
		updateTool models.UpdateTool
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateTool)
	if err != nil {
		h.handlerResponse(c, "error tool should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateTool.Id = id
	rowsAffected, err := h.strg.Tool().Update(c.Request.Context(), &updateTool)
	if err != nil {
		h.handlerResponse(c, "storage.tool.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.tool.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Tool().GetByID(c.Request.Context(), &models.ToolPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.tool.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create tool resposne", http.StatusAccepted, resp)
}

// Delete tool godoc
// @ID delete_tool
// @Router /tool/{id} [DELETE]
// @Summary Delete Tool
// @Description Delete Tool
// @Tags Tool
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteTool(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Tool().Delete(c.Request.Context(), &models.ToolPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.tool.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create tool resposne", http.StatusNoContent, nil)
}
