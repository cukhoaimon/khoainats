/*
 * Khoai NATS Admin API
 *
 *
 * API version: <VERSION>
 * Contact: phuc dep trai (phucmapcaumieu@gmail.com)
 */

package openapi

type V1PrincipalType string

// List of V1PrincipalType
const (
	EMAIL_CODE V1PrincipalType = "EmailCode"
	PASSWORD V1PrincipalType = "Password"
)
