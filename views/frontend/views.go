package views

import "text/template"

type View struct {
	Template *template.Template
}

func NewView(files ...string) *View {
	files = append(files, "views/frontend/templates/header.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}
