package forms

import (
	"net/http"
	"net/url"
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

// Has checks if form field is in post is not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		return false
	}

	return true
}
