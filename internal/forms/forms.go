package forms

import (
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

//New initializes a Form structure
func New(data url.Values) *Form {
	return &Form{
		data,
		// Declaring this errors string as empty
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be left blank")
		}
	}
}

// Has checks if form field is in post is not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}

// If the Error array is empty it will be valid, elsewise invalid
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
