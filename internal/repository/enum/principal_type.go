package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type PrincipalType string

const (
	PrincipalUser PrincipalType = "USER"
	PrincipalService PrincipalType = "SERVICE"
)

func (p PrincipalType) IsValid() bool {
	switch p {
	case PrincipalUser:
	case PrincipalService:
		return true
	}
	return false
}


func (p *PrincipalType) MarshalJSON() ([]byte, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid PrincipalType: %s", *p)
	}
	return json.Marshal(string(*p))
}

func (p *PrincipalType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	pt := PrincipalType(s)
	if !pt.IsValid() {
		return fmt.Errorf("invalid PrincipalType: %s", s)
	}
	*p = pt
	return nil
}

func (p *PrincipalType) Value() (driver.Value, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid PrincipalType: %s", *p)
	}
	return string(*p), nil
}

func (p *PrincipalType) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("PrincipalType should be a string")
	}
	*p = PrincipalType(str)
	if !p.IsValid() {
		return fmt.Errorf("invalid PrincipalType: %s", str)
	}
	return nil
}
