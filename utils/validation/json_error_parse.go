package validation

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

func DecodeJSON(body io.Reader, input interface{}) error {
	d := json.NewDecoder(body)
	d.DisallowUnknownFields()

	e := d.Decode(input)
	if e != nil {
		var (
			se  *json.SyntaxError
			ute *json.UnmarshalTypeError
		)

		switch {
		case
			errors.Is(e, io.EOF),
			errors.Is(e, io.ErrUnexpectedEOF),
			errors.As(e, &se):
			return errors.New("Malformed JSON payload.")
		case
			errors.As(e, &ute):
			return errors.New("Unexpected field value.")
		case
			strings.HasPrefix(e.Error(), "json: unknown field "):
			return errors.New("Unexpected field key.")
		default:
			return errors.New("Invalid request payload.")
		}
	}

	return nil
}