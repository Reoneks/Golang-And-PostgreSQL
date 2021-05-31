package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
 *ctx.Errors = append(ctx.Errors, &gin.Error{
 *	Err:  errors.New("some error"),
 *	Type: http.StatusBadRequest,
 *	Meta: "Some message",
 *})
 */
func ErrorHandlingMiddleware(numberOfFunctions int64, log *logrus.Entry) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for i := int64(0); i < numberOfFunctions; i++ {
			ctx.Next()
			errorFounded := false
			for err := ctx.Errors.Last(); err != nil; err = ctx.Errors.Last() {
				errInfo := err.Meta.(string)

				log.WithFields(logrus.Fields{
					"time":    time.Now(),
					"message": errInfo,
				}).WithError(err.Err)

				ctx.JSON(int(err.Type), gin.H{
					"Message:": errInfo,
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
