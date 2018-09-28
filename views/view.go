package views

import (
	"html/template"
	"io"
)

func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/semantic.gohtml",
		"views/layouts/footer.gohtml")

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(wr io.Writer, data interface{}) error {
	err := v.Template.ExecuteTemplate(wr, v.Layout, nil)
	if err != nil {
		return err
	}
	return nil
}
