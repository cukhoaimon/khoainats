package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtRequestFilter(rolesAllowed []PrincipalRoleType) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt, ok := ctx.Get("Authorization")
		if !ok {
			_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("authorization header must be set"))
		}

		_, ok = strings.CutPrefix(jwt.(string), "Bearer ")
		if !ok {
			_ = ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("not a bearer token"))
		}

		ctx.Next()
		//panic("TODO: not yet implemented")
	}
}
