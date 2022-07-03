package pkg

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig("../config.yaml")
	if err != nil {
		t.Error(err)
	}
	t.Logf("get config %+v", config)
}
