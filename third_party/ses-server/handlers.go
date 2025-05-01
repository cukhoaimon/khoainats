package ses_server

import (
	"log"
	"net/http"

	"github.com/cukhoaimon/khoainats/third_party/database"
	"github.com/gin-gonic/gin"
)

type V1ExchangeRequest struct {
	Email string `json:"email"`
}

type V1VerifyCodeRequest struct {
	Code string `json:"code"`
}

func v1CodeExchange(db database.AbstractDatabase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req V1ExchangeRequest
		if err := ctx.ShouldBind(&req); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, http.ErrBodyNotAllowed)
			return
		}
		code := randSeq()
		log.Printf("sending email for %s with code=%s\n", req.Email, code)

		if err := db.Write(req.Email, code); err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, http.ErrServerClosed)
			return
		}

		ctx.JSON(200, code)
	}

}

func v1Verify(db database.AbstractDatabase) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func v1GetKeys(db database.AbstractDatabase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, db.ReadAll())
	}
}
