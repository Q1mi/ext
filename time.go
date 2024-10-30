package ext

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/migrator"
	"gorm.io/gorm/schema"
)

const (
	quote = `"`
	null  = "null"
)

var _ json.Marshaler = (*Time)(nil)
var _ json.Unmarshaler = (*Time)(nil)

var _ migrator.GormDataTypeInterface = (*Time)(nil)
var _ schema.GormDataTypeInterface = (*Time)(nil)
var _ driver.Valuer = (*Time)(nil)
var _ sql.Scanner = (*Time)(nil)

type Time time.Time

// UnmarshalJSON unmarshal from value
func (t *Time) UnmarshalJSON(bs []byte) error {
	s := strings.Trim(string(bs), quote)
	if len(s) == 0 || s == null {
		return nil
	}
	v, err := time.Parse(time.DateTime, s)
	if err != nil {
		return nil
	}

	*t = Time(v)
	return nil
}

// MarshalJSON marshal to json
func (t Time) MarshalJSON() ([]byte, error) {
	rt := time.Time(t)
	if rt.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(strconv.Quote(rt.Format(time.DateTime))), nil
}

func (t Time) GormDataType() string {
	return "time"
}

func (t Time) GormDBDataType(*gorm.DB, *schema.Field) string {
	return "datetime"
}

func (t Time) Value() (driver.Value, error) {
	rt := time.Time(t)
	if rt.IsZero() {
		return nil, nil
	}
	return rt, nil
}

func (t *Time) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	st, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("fail to scan Time value:%v", value)
	}
	*t = Time(st)
	return nil
}
