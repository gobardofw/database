package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// NullInt nullable string keeper
type NullInt struct {
	Int   int
	Valid bool
}

// Scan implements the Scanner interface.
func (me *NullInt) Scan(value interface{}) error {
	me.Int, me.Valid = 0, false
	if value != nil {
		temp := sql.NullInt64{}
		err := temp.Scan(value)
		if err != nil {
			return err
		}
		me.Int, me.Valid = int(temp.Int64), temp.Valid
	}
	return nil
}

// Value implements the driver Valuer interface.
func (me *NullInt) Value() (driver.Value, error) {
	if !me.Valid {
		return nil, nil
	}
	return me.Int, nil
}

// Val get nullable value
func (me *NullInt) Val() interface{} {
	if !me.Valid {
		return nil
	}
	return me.Int
}

// MarshalJSON convert to json
func (me NullInt) MarshalJSON() ([]byte, error) {
	if me.Valid {
		return json.Marshal(me.Int)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON parse from json
func (me *NullInt) UnmarshalJSON(data []byte) error {
	var v *int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v != nil {
		me.Valid = true
		me.Int = *v
	} else {
		me.Valid = false
	}
	return nil
}
