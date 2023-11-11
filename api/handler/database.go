package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create database godoc
// @ID create_database
// @Router /database [POST]
// @Summary Create Database
// @Description Create Database
// @Tags Database
// @Accept json
// @Procedure json
// @Param Database body models.CreateDatabase true "CreateDatabaseRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateDatabase(c *gin.Context) {

	var createDatabase models.CreateDatabase
	err := c.ShouldBindJSON(&createDatabase)
	if err != nil {
		h.handlerResponse(c, "error database should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Database().Create(c.Request.Context(), &createDatabase)
	if err != nil {
		h.handlerResponse(c, "storage.database.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Database().GetByID(c.Request.Context(), &models.DatabasePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.database.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create database resposne", http.StatusCreated, resp)
}

// GetByID database godoc
// @ID get_by_id_database
// @Router /database/{id} [GET]
// @Summary Get By ID Database
// @Description Get By ID Database
// @Tags Database
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdDatabase(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// DatabaseData := val.(helper.TokenInfo)
	// if len(DatabaseData) > 0 {
	// 	id = DatabaseData.DatabaseID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Database().GetByID(c.Request.Context(), &models.DatabasePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.database.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id database resposne", http.StatusOK, resp)
}

// GetList database godoc
// @ID get_list_database
// @Router /database [GET]
// @Summary Get List Database
// @Description Get List Database
// @Tags Database
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListDatabase(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list database offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list database limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Database().GetList(c.Request.Context(), &models.DatabaseGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.database.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list database resposne", http.StatusOK, resp)
}

// Update database godoc
// @ID update_database
// @Router /database/{id} [PUT]
// @Summary Update Database
// @Description Update Database
// @Tags Database
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Database body models.UpdateDatabase true "UpdateDatabaseRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateDatabase(c *gin.Context) {

	var (
		id             string = c.Param("id")
		updateDatabase models.UpdateDatabase
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateDatabase)
	if err != nil {
		h.handlerResponse(c, "error database should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateDatabase.Id = id
	rowsAffected, err := h.strg.Database().Update(c.Request.Context(), &updateDatabase)
	if err != nil {
		h.handlerResponse(c, "storage.database.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.database.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Database().GetByID(c.Request.Context(), &models.DatabasePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.database.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create database resposne", http.StatusAccepted, resp)
}

// Delete database godoc
// @ID delete_database
// @Router /database/{id} [DELETE]
// @Summary Delete Database
// @Description Delete Database
// @Tags Database
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteDatabase(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Database().Delete(c.Request.Context(), &models.DatabasePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.database.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create database resposne", http.StatusNoContent, nil)
}
