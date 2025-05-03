/*
 * Khoai NATS Admin API
 *
 *
 * API version: <VERSION>
 * Contact: phuc dep trai (phucmapcaumieu@gmail.com)
 */

package openapi

import (
	"time"
)

// V1AccessToken - Access Token
type V1AccessToken struct {

	Id string `json:"id"`

	PrincipalId string `json:"principalId"`

	PrincipalType V1PrincipalType `json:"principalType"`

	OrganizationId string `json:"organizationId"`

	RevokedBy string `json:"revokedBy,omitempty"`

	RevokedAt time.Time `json:"revokedAt,omitempty"`

	CreatedAt time.Time `json:"createdAt"`

	UpdatedAt time.Time `json:"updatedAt"`

	Roles []V1PrincipalRoleType `json:"roles"`

	JwkId string `json:"jwkId,omitempty"`
}
