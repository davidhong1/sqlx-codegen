// GENERATED BY sqlx-codegen; DO NOT EDIT.

package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"errors"
)


func (s *User) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &s)
		return nil
	case string:
		json.Unmarshal([]byte(v), &s)
		return nil
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (s User) Value() (driver.Value, error) {
	return json.Marshal(s)
}


