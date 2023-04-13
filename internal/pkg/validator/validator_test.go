package validator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_Check(t *testing.T) {
	type TestCase struct {
		Name         string
		DataNotValid bool
		ReqBody      string
	}

	type DummyReq struct {
		Name        string `validate:"required" json:"name"`
		Description string `validate:"required" json:"description"`
	}

	cases := []TestCase{
		{
			Name:         "when name not presence",
			DataNotValid: true,
			ReqBody:      `{"description": "foobar"}`,
		},
		{
			Name:         "when description not presence",
			DataNotValid: true,
			ReqBody:      `{"name": "foo"}`,
		},
		{
			Name:         "when name and description presence",
			DataNotValid: false,
			ReqBody:      `{"name": "foo", "description": "foobar"}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var req DummyReq
			_ = json.Unmarshal([]byte(tc.ReqBody), &req)
			isError := Check(&req)
			assert.Equal(t, tc.DataNotValid, isError)
		})
	}
}
