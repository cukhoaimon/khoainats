/*
 * Khoai NATS Admin API
 *
 *
 * API version: <VERSION>
 * Contact: phuc dep trai (phucmapcaumieu@gmail.com)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
)

type DefaultAPI interface {

    // V1LoginExchange Post /v1/login/exchange
     V1LoginExchange(c *gin.Context)

    // V1LoginStart Post /v1/login/start
     V1LoginStart(c *gin.Context)

}
