package middleware

import (
	"errors"
	"net/http"

	"github.com/arikama/koran-backend/constants"
	"github.com/arikama/koran-backend/managers"
	"github.com/arikama/koran-backend/utils"
	"github.com/gin-gonic/gin"
)

func UserAccessMiddleware(userManager managers.UserManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get(constants.XAccessToken())

		if accessToken == "" {
			utils.JsonError(c, http.StatusUnauthorized, errors.New(`missing x-access-token header`))
			c.Abort()
			return
		}

		_, err := userManager.GetUser(accessToken)

		if err != nil {
			utils.JsonError(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}
	}
}
