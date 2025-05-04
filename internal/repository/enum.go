package repository

type PrincipalType int

const (
	USER PrincipalType = iota
	SERVICE
)

var principalType = map[PrincipalType]string{
	USER:    "user",
	SERVICE: "service",
}

func (p PrincipalType) String() string {
	return principalType[p]
}

type PrincipalAttributeType int

const (
	Email PrincipalAttributeType = iota
	MacAddress
)

var principalAttributes = map[PrincipalAttributeType]string{
	Email:      "email",
	MacAddress: "mac_address",
}

func (p PrincipalAttributeType) String() string {
	return principalAttributes[p]
}
