/*
 * Khoai NATS Admin API
 *
 *
 * API version: <VERSION>
 * Contact: phuc dep trai (phucmapcaumieu@gmail.com)
 */

package openapi

import (
	"net/http"

    "github.com/cukhoaimon/khoainats/internal/auth"
	"github.com/gin-gonic/gin"
)

type JwtFilterFunc func(rolesAllowed []auth.PrincipalRoleType) gin.HandlerFunc

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name		string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method		string
	// Pattern is the pattern of the URI.
	Pattern	 	string
	// HandlerFunc is the handler function of this route.
	HandlerFunc	gin.HandlerFunc
    // RolesAllowed
    RolesAllowed []auth.PrincipalRoleType
}

// NewRouter returns a new router.
func NewRouter(handleFunctions DefaultAPI, jwtFilterFunc JwtFilterFunc) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), jwtFilterFunc, handleFunctions)
}

// NewRouterWithGinEngine add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, jwtFilterFunc JwtFilterFunc, handleFunctions DefaultAPI) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}

        handlersChain := []gin.HandlerFunc{route.HandlerFunc}
        if len(route.RolesAllowed) > 0 {
            handlersChain = append(handlersChain, jwtFilterFunc(route.RolesAllowed))
        }

		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, handlersChain...)
		case http.MethodPost:
			router.POST(route.Pattern, handlersChain...)
		case http.MethodPut:
			router.PUT(route.Pattern, handlersChain...)
		case http.MethodPatch:
			router.PATCH(route.Pattern, handlersChain...)
		case http.MethodDelete:
			router.DELETE(route.Pattern, handlersChain...)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

func getRoutes(handleFunctions DefaultAPI) []Route {
	return []Route{ 
		{
			"V1LoginExchange",
			http.MethodPost,
			"/v1/login/exchange",
			handleFunctions.V1LoginExchange,
            []auth.PrincipalRoleType{},
		},
		{
			"V1LoginStart",
			http.MethodPost,
			"/v1/login/start",
			handleFunctions.V1LoginStart,
            []auth.PrincipalRoleType{},
		},
	}
}
