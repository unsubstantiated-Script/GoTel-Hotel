package models

import "GoTel/internal/forms"

type TemplateData struct {
	// Data to pass to a template
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	// For the unsure items you wanna pass to a template
	Data      map[string]interface{}
	Flash     string
	Warning   string
	Error     string
	CSRFToken string
	Form      *forms.Form
}
