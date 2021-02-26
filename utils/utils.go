package utils

import (
	"encoding/json"
	"fmt"
	httpext "github.com/go-playground/pkg/net/http"
	"github.com/gorilla/schema"
	"github.com/todo/types/apiFunc"
	validator "gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

var validate *validator.Validate
var decoder *schema.Decoder

func init() {
	decoder = schema.NewDecoder()
	decoder.SetAliasTag("json")
	decoder.IgnoreUnknownKeys(true)
	decoder.ZeroEmpty(true)

	validate = validator.New()
}

func SetStatusCode(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func Prepare(dtoDefinition interface{}, apiFunc apiFunc.APIFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// copy the dtoDefinition so we don't use the same object for each request
		dto := reflect.New(reflect.TypeOf(dtoDefinition)).Interface()

		json.NewDecoder(r.Body).Decode(&dto)
		// validate request
		err := validate.Struct(dto)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				writeResponse(w, http.StatusInternalServerError, nil)
				return
			}

			validationErrs := err.(validator.ValidationErrors)

			// log everything and build the error messages
			var errMsgs []string
			for _, err := range validationErrs {
				errMsgs = append(errMsgs,getValidationErrorMessage(err))
			}

			// just return 400 if any validation errors found
			if len(validationErrs) > 0 {
				writeResponse(w, http.StatusBadRequest, errMsgs)
				return
			}
		}

		// call the api then write response
		statusCode, output := apiFunc(dto)
		writeResponse(w, statusCode, output)
	}
}

func writeResponse(w http.ResponseWriter, statusCode int, output interface{}) {
	if output == nil {
		w.WriteHeader(statusCode)
		return
	}

	// if it is a slice of bytes, then write directly
	if outBytes, ok := output.([]byte); ok {
		w.Header().Set(httpext.ContentLength, strconv.FormatInt(int64(len(outBytes)), 10))
		w.WriteHeader(statusCode)
		w.Write(outBytes)
	} else {
		outputJSON, _ := json.Marshal(output)
		w.Header().Set(httpext.ContentType, httpext.ApplicationJSON)
		w.Header().Set(httpext.ContentLength, strconv.FormatInt(int64(len(outputJSON)), 10))
		w.WriteHeader(statusCode)
		w.Write(outputJSON)
	}
}

func getValidationErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}