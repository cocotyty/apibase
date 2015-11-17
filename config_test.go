package apibase_test

import "testing"

import (
	"github.com/cocotyty/apibase"
)

func TestLoadConfig(t *testing.T) {
	config, err := apibase.TransferConfig([]byte(`{"version":20.2,"database":{"":{"type":"MYSQL"}}}`))
	if err != nil {
		t.Error(err)
	}
	t.Log("[config]:", config)
}
