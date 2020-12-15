package dst

import (
	"bytes"
	"io"
	"text/template"
)

type Template struct {
	T *template.Template
}

func NewTemplate(b string) *Template {
	t := template.Must(template.New("").Parse(b))
	return &Template{
		T: t,
	}
}

func (t *Template) Prepare(d interface{}) {
	var b bytes.Buffer
	t.T.Execute(&b, d)
	t.T = template.Must(template.New("").Parse(b.String()))
}

func (t *Template) Execute(w io.Writer, d interface{}) {
	t.T.Execute(w, d)
}
