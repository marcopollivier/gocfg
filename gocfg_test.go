package gocfg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadYaml(t *testing.T) {
	var testObject = struct {
		SomeKey    string `cfg:"some_key"`
		AnotherKey int    `cfg:"another_key"`
		SubObject  struct {
			SomeThing                      string  `cfg:"something"`
			SomethingElse                  int     `cfg:"something_else"`
			SomethingElseWithInterpolation float64 `cfg:"something_else_with_interpolation"`
		} `cfg:"sub_object"`
		LastKey           string `cfg:"last_key"`
		LastKeyWithNumber uint   `cfg:"last_key_with_number"`
	}{}

	err := Load(&testObject, YAML, "test.yml")
	assert.Nil(t, err)
	fmt.Printf("%+v", testObject)
}

func TestLoadJSON(t *testing.T) {
	var testObject = struct {
		SomeKey    string `cfg:"some_key"`
		AnotherKey int    `cfg:"another_key"`
		SubObject  struct {
			SomeThing                      string  `cfg:"something"`
			SomethingElse                  int     `cfg:"something_else"`
			SomethingElseWithInterpolation float64 `cfg:"something_else_with_interpolation"`
		} `cfg:"sub_object"`
		LastKey           string `cfg:"last_key"`
		LastKeyWithNumber uint   `cfg:"last_key_with_number"`
	}{}

	err := Load(&testObject, JSON, "test.json")
	assert.Nil(t, err)
	fmt.Printf("%+v", testObject)
}