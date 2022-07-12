package validate

import (
	"reflect"
	"testing"

	"github.com/webbmaffian/go-validate/utils"
)

func TestStructSkipTag(t *testing.T) {
	type Status byte

	const (
		Undefined Status = iota
		Null
		Present
	)

	type Text struct {
		String string
		Status Status
	}

	type User struct {
		FirstName Text `validate:"required"`
		LastName  Text `validate:"required"`
	}

	RegisterValueGetter(func(originalValue reflect.Value, opt *utils.Options) any {
		value := originalValue.Interface()

		switch val := value.(type) {
		case Text:
			if val.Status == Undefined {
				return nil
			}

			return val.String
		}

		return originalValue
	}, Text{})

	user := User{
		FirstName: Text{String: "", Status: Null},
	}
	_, errs := Struct(user, utils.Options{
		SkipNil: true,
	})

	for _, err := range errs {
		t.Log(err.Path, "-", err.Tag, "-", err.Message)
	}

	t.FailNow()
}
