package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"perisie.com/koran/constants"
	"perisie.com/koran/managers"
	"perisie.com/koran/utils"
)

func UserAuth(userManager managers.UserManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(constants.XAccessToken())

		if token == "" {
			utils.JsonError(c, http.StatusUnauthorized, errors.New(`missing x-access-token header`))
			c.Abort()
			return
		}

		_, err := userManager.GetUser(token)

		if err != nil {
			utils.JsonError(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}
	}
}
