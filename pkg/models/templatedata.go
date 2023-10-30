package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]int
	Data      map[string]interface{}
	CSFRToken string
	Flash     string
	Warning   string
	Error     string
}
