package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type PrincipalAttributeType string

const (
	Email PrincipalAttributeType = "EMAIL"
	MacAddress PrincipalAttributeType = "MAC_ADDRESS"
)

func (p *PrincipalAttributeType) IsValid() bool {
	switch *p {
	case Email:
	case MacAddress:
		return true
	}
	return false
}


func (p *PrincipalAttributeType) MarshalJSON() ([]byte, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid PrincipalAttributeType: %s", *p)
	}
	return json.Marshal(string(*p))
}

// UnmarshalJSON Implement json.Unmarshaler
func (p *PrincipalAttributeType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	pt := PrincipalAttributeType(s)
	if !pt.IsValid() {
		return fmt.Errorf("invalid PrincipalAttributeType: %s", s)
	}
	*p = pt
	return nil
}

// Value Implement driver.Valuer for SQL
func (p *PrincipalAttributeType) Value() (driver.Value, error) {
	if !p.IsValid() {
		return nil, fmt.Errorf("invalid PrincipalAttributeType: %s", *p)
	}
	return string(*p), nil
}

func (p *PrincipalAttributeType) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("PrincipalAttributeType should be a string")
	}
	*p = PrincipalAttributeType(str)
	if !p.IsValid() {
		return fmt.Errorf("invalid PrincipalAttributeType: %s", str)
	}
	return nil
}
