package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/s-owl/sowl_manager_backend/firebaseapp"
)

// AdminUserMiddleware - HTTP Header 에서 Firebase ID Token으로 유저를 식별하는 미들웨어
func AdminUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authClient := firebaseapp.App().Auth
		idToken := c.GetHeader("Firebase-ID-Token")
		if idToken == "" {
			c.AbortWithStatusJSON(UserTokenInvalidError.StatusCode, UserTokenInvalidError)
			return
		}

		decoded, err := authClient.VerifyIDTokenAndCheckRevoked(c, idToken)
		if err != nil {
			c.AbortWithStatusJSON(UserTokenInvalidError.StatusCode, UserTokenInvalidError)
			return
		}

		user, err := authClient.GetUser(c, decoded.UID)
		if err != nil {
			c.AbortWithStatusJSON(UserTokenInvalidError.StatusCode, UserTokenInvalidError)
			return
		}
		if user.UserInfo.Email == "" {
			c.AbortWithStatusJSON(NoPermissionError.StatusCode, NoPermissionError)
			return
		}

		c.Set("user", user)

		c.Next()
		return
	}
}
