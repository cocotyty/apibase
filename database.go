package apibase

import (
	"encoding/json"
	"errors"
	"strings"
)

type Database struct {
	Type DatabaseType `json:"type"`
}
type Databases map[string]Database

type DatabaseType int

const (
	MYSQL DatabaseType = iota
	SQLITE
)

var ErrDatabaseType error = errors.New("ErrDatabaseType")

func (this DatabaseType) String() string {
	switch this {
	case MYSQL:
		return "MYSQL"
	case SQLITE:
		return "SQLITE"
	default:
		return "NOTSUPPORT"
	}
}
func (this *DatabaseType) UnmarshalJSON(data []byte) error {
	var str string
	json.Unmarshal(data, &str)
	str = strings.ToLower(str)
	switch str {
	case "mysql":
		*this = MYSQL
	case "sqlite":
		*this = SQLITE
	default:
		return ErrDatabaseType
	}
	return nil
}
