package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create post godoc
// @ID create_post
// @Router /post [POST]
// @Summary Create Post
// @Description Create Post
// @Tags Post
// @Accept json
// @Procedure json
// @Param Post body models.CreatePost true "CreatePostRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreatePost(c *gin.Context) {

	var createPost models.CreatePost
	err := c.ShouldBindJSON(&createPost)
	if err != nil {
		h.handlerResponse(c, "error post should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Post().Create(c.Request.Context(), &createPost)
	if err != nil {
		h.handlerResponse(c, "storage.post.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Post().GetByID(c.Request.Context(), &models.PostPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.post.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create post resposne", http.StatusCreated, resp)
}

// GetByID post godoc
// @ID get_by_id_post
// @Router /post/{id} [GET]
// @Summary Get By ID Post
// @Description Get By ID Post
// @Tags Post
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdPost(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// PostData := val.(helper.TokenInfo)
	// if len(PostData) > 0 {
	// 	id = PostData.PostID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Post().GetByID(c.Request.Context(), &models.PostPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.post.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id post resposne", http.StatusOK, resp)
}

// GetList post godoc
// @ID get_list_post
// @Router /post [GET]
// @Summary Get List Post
// @Description Get List Post
// @Tags Post
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListPost(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list post offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list post limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Post().GetList(c.Request.Context(), &models.PostGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.post.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list post resposne", http.StatusOK, resp)
}

// Update post godoc
// @ID update_post
// @Router /post/{id} [PUT]
// @Summary Update Post
// @Description Update Post
// @Tags Post
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Post body models.UpdatePost true "UpdatePostRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdatePost(c *gin.Context) {

	var (
		id         string = c.Param("id")
		updatePost models.UpdatePost
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updatePost)
	if err != nil {
		h.handlerResponse(c, "error post should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updatePost.Id = id
	rowsAffected, err := h.strg.Post().Update(c.Request.Context(), &updatePost)
	if err != nil {
		h.handlerResponse(c, "storage.post.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.post.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Post().GetByID(c.Request.Context(), &models.PostPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.post.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create post resposne", http.StatusAccepted, resp)
}

// Delete post godoc
// @ID delete_post
// @Router /post/{id} [DELETE]
// @Summary Delete Post
// @Description Delete Post
// @Tags Post
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeletePost(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Post().Delete(c.Request.Context(), &models.PostPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.post.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create post resposne", http.StatusNoContent, nil)
}
