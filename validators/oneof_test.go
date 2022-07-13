package validators

import (
	"reflect"
	"testing"

	"github.com/webbmaffian/go-validate/utils"
)

func TestOneof(t *testing.T) {
	value := []string{"a"}

	if err := Oneof(reflect.ValueOf(value), reflect.Value{}, "a b", &utils.Options{}); err != nil {
		t.Error(err)
	}
}
