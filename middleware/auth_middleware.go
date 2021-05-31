package middleware

import (
	"net/http"
	"strings"
	"test/config"
	"test/user"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth"
)

func AuthMiddleware(userService user.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Request.Header["Authorization"]) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Not enough data (There is no token)",
			})
			ctx.Abort()
			return
		} else if len(strings.Split(ctx.Request.Header["Authorization"][0], " ")) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token issued wrong",
			})
			ctx.Abort()
			return
		}
		reqToken := strings.Split(ctx.Request.Header["Authorization"][0], " ")[1]
		jwtNew := config.NewConfig().JWT()
		token, err := jwtauth.VerifyToken(jwtNew, reqToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		userId, getted := token.Get("user_id")
		if !getted {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "can't get user_id",
			})
			ctx.Abort()
			return
		}
		obtainedUser, err := userService.GetUser(int64(userId.(float64)))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		} else if obtainedUser == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "can't find user",
			})
			ctx.Abort()
			return
		}
		if obtainedUser.Status != user.Active && obtainedUser.Status != user.UnActive {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "can't login",
			})
			ctx.Abort()
			return
		}
		ctx.Set("user", obtainedUser)
		ctx.Next()
	}
}
