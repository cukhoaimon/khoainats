/*
 * Khoai NATS Admin API
 *
 *
 * API version: <VERSION>
 * Contact: phuc dep trai (phucmapcaumieu@gmail.com)
 */

package openapi

type V1PrincipalRoleType string

// List of V1PrincipalRoleType
const (
	ADMIN V1PrincipalRoleType = "Admin"
	USER V1PrincipalRoleType = "User"
	SERVICE V1PrincipalRoleType = "Service"
	CUSTOMER_ADMIN V1PrincipalRoleType = "CustomerAdmin"
)
