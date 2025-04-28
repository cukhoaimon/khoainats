package auth

type JwtSecurityContext struct {
	Subject JwtPrincipal
	Secure  bool
}
