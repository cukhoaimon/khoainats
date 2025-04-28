package auth

type PrincipalRoleType int

const (
	Admin PrincipalRoleType = iota
	User
	Service
	CustomerAdmin
)

var roleName = map[PrincipalRoleType]string{
	Admin:         "admin",
	User:          "user",
	Service:       "service",
	CustomerAdmin: "customer_admin",
}

func (p PrincipalRoleType) String() string {
	return roleName[p]
}
