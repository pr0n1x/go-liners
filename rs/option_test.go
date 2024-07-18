package rs

import (
	"encoding/json"
	"testing"
)

func TestOptionNull(t *testing.T) {
	objectWithOptional := struct {
		Opt Option[string] `json:"opt"`
	}{}

	const nullOpt = `{"opt":null}`

	if string(Must(json.Marshal(objectWithOptional))) != nullOpt {
		t.Fatal("opt field must be null")
	}
	objectWithOptional.Opt = SetNone[string]()
	if string(Must(json.Marshal(objectWithOptional))) != nullOpt {
		t.Fatal("opt field must be null")
	}
}
