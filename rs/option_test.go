package rs

import (
	"encoding/json"
	"github.com/pr0n1x/go-type-wrappers/assert"
	"testing"
)

func TestOptionNull(t *testing.T) {
	objectWithOptional := struct {
		Opt Option[string] `json:"opt"`
	}{}

	const nullOpt = `{"opt":null}`

	if string(assert.Must(json.Marshal(objectWithOptional))) != nullOpt {
		t.Fatal("opt field must be null")
	}
	objectWithOptional.Opt = SetNone[string]()
	if string(assert.Must(json.Marshal(objectWithOptional))) != nullOpt {
		t.Fatal("opt field must be null")
	}
}
