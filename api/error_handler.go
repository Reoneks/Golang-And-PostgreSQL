package api

import (
	"test/config"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Title   string `json:"title"`
	Details string `json:"details"`
}

/*
 *ctx.Errors = append(ctx.Errors, &gin.Error{
 *	Err:  errors.New("some error"),
 *	Type: http.StatusBadRequest,
 *	Meta: Error{},
 *})
 */
func ErrorHandlingMiddleware(numberOfFunctions int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conf := config.NewConfig()
		log := conf.Log()
		for i := int64(0); i < numberOfFunctions; i++ {
			ctx.Next()
			errorFounded := false
			for err := ctx.Errors.Last(); err != nil; err = ctx.Errors.Last() {
				errInfo := err.Meta.(Error)
				log.Error(errInfo.Title, errInfo.Details, err)
				ctx.JSON(int(err.Type), gin.H{
					"Title":   errInfo.Title,
					"Details": errInfo.Details,
				})
				errorFounded = true
				ctx.Abort()
			}
			if errorFounded {
				return
			}
		}
	}
}
