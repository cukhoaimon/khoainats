/*
 * Khoai NATS Admin API
 *
 *
 * API version: <VERSION>
 * Contact: phuc dep trai (phucmapcaumieu@gmail.com)
 */

package openapi

type V1LoginExchangeRequest struct {

	Email string `json:"email"`

	PrincipalType V1PrincipalType `json:"principalType"`

	PasswordOrCode string `json:"passwordOrCode"`
}
