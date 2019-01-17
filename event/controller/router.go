package controller

import "text/template"

var (
	l letter
)

func Startup(T *template.Template) {
	l.temp = T
	l.permissionLetterHandler()
	eventCRUDHandler()
}
