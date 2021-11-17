package models

import (
	"OneAuth/cfg"
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid scan source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}
func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal(j, j1)
}

// JSONTime  custom json time
type JSONTime struct {
	time.Time
}

func Now() *JSONTime {
	return &JSONTime{time.Now()}
}

// MarshalJSON 实现它的json序列化方法
func (jt JSONTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", jt.Format(cfg.CFG.TimeFormat))
	return []byte(stamp), nil
}

// UnmarshalJSON  反序列化方法
func (jt *JSONTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+cfg.CFG.TimeFormat+`"`, string(data), time.Local)
	*jt = JSONTime{now}
	return
}

// Value insert timestamp into mysql need this function.
func (jt JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if jt.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return jt.Time, nil
}

// Scan value of time.Time
func (jt *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*jt = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (jt *JSONTime) SetTime(t time.Time) {
	jt.Time = t
}

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt JSONTime  `json:"created_at"`
	UpdatedAt JSONTime  `json:"updated_at"`
	DeletedAt *JSONTime `json:"deleted_at" sql:"index"`
}
