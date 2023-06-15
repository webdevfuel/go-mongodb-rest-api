package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NotEmpty(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(field.String())) > 0
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

func getValidator() (*validator.Validate, error) {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	err := v.RegisterValidation("notempty", NotEmpty)
	if err != nil {
		return v, nil
	}
	return v, nil
}

func validationTagToHumanReadable(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	default:
		return ""
	}
}

func validatorError(err error, w http.ResponseWriter) {
	errs := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()
			tag := validationTagToHumanReadable(fieldError.Tag())
			errs[field] = tag
		}
	}

	j, _ := json.Marshal(errs)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(j)
}

func decoderError(err error, w http.ResponseWriter) {
	errs := make(map[string]string)

	if err, ok := err.(*json.UnmarshalTypeError); ok {
		message := fmt.Sprintf("This field requires a value of type %s.", err.Type)
		errs[err.Field] = message

		j, _ := json.Marshal(errs)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
