package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// @Security ApiKeyAuth
// Create question godoc
// @ID create_question
// @Router /question [POST]
// @Summary Create Question
// @Description Create Question
// @Tags Question
// @Accept json
// @Procedure json
// @Param Question body models.CreateQuestion true "CreateQuestionRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateQuestion(c *gin.Context) {

	var createQuestion models.CreateQuestion
	err := c.ShouldBindJSON(&createQuestion)
	if err != nil {
		h.handlerResponse(c, "error question should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Question().Create(c.Request.Context(), &createQuestion)
	if err != nil {
		h.handlerResponse(c, "storage.question.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Question().GetByID(c.Request.Context(), &models.QuestionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.question.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create question resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID question godoc
// @ID get_by_id_question
// @Router /question/{id} [GET]
// @Summary Get By ID Question
// @Description Get By ID Question
// @Tags Question
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdQuestion(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// questionData := val.(helper.TokenInfo)
	// if len(questionData) > 0 {
	// 	id = questionData.QuestionID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Question().GetByID(c.Request.Context(), &models.QuestionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.question.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id question resposne", http.StatusOK, resp)
}

// GetList question godoc
// @ID get_list_question
// @Router /question [GET]
// @Summary Get List Question
// @Description Get List Question
// @Tags Question
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListQuestion(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list question offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list question limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Question().GetList(c.Request.Context(), &models.QuestionGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.question.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list question resposne", http.StatusOK, resp)
}

// Update question godoc
// @ID update_question
// @Router /question/{id} [PUT]
// @Summary Update Question
// @Description Update Question
// @Tags Question
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Question body models.UpdateQuestion true "UpdateQuestionRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateQuestion(c *gin.Context) {

	var (
		id             string = c.Param("id")
		updateQuestion models.UpdateQuestion
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateQuestion)
	if err != nil {
		h.handlerResponse(c, "error question should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateQuestion.Id = id
	rowsAffected, err := h.strg.Question().Update(c.Request.Context(), &updateQuestion)
	if err != nil {
		h.handlerResponse(c, "storage.question.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.question.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Question().GetByID(c.Request.Context(), &models.QuestionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.question.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create question resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete question godoc
// @ID delete_question
// @Router /question/{id} [DELETE]
// @Summary Delete Question
// @Description Delete Question
// @Tags Question
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteQuestion(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Question().Delete(c.Request.Context(), &models.QuestionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.question.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create question resposne", http.StatusNoContent, nil)
}
