package apibase

import (
	"encoding/json"
)

type ApiBaseConfig struct {
	Version  float64         `json:"version"`
	Database Databases       `json:"database"`
	Own      json.RawMessage `json:"own"`
}

func TransferConfig(config []byte) (conf *ApiBaseConfig, err error) {
	err = json.Unmarshal(config, &conf)
	return
}
