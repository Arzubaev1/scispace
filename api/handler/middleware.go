package handler

import (
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		value := c.GetHeader("Authorization")

		info, err := helper.ParseClaims(value, h.cfg.SecretKey)

		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}

		c.Set("Auth", info)
		c.Next()
	}
}
