package models

type TemplateData struct {
	// Data to pass to a template
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	// For the unsure items you wanna pass to a template
	Data     map[string]interface{}
	CSRToken string
	Flash    string
	Warning  string
	Error    string
}
