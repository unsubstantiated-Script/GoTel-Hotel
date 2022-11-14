package forms

//Map of strings with a slice of strings
type errors map[string][]string

// Adds an error message for a given field
func (e errors) Add(field, message string) {
	//Appending the error to a particular slice
	e[field] = append(e[field], message)
}

// Get returns the first error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	//Testing to see if a given field has an error and returning if so
	return es[0]
}
