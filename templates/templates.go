package templates

import "text/template"

var templates = map[string]string{"import.tmpl": `import ({{range .}}
	{{if ne "" .Alias}}{{printf "%s " .Alias}}{{end}}"{{.Path}}"{{end}}
)
`,
}

// Parse parses declared templates.
func Parse(t *template.Template) (*template.Template, error) {
	for name, s := range templates {
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		if _, err := tmpl.Parse(s); err != nil {
			return nil, err
		}
	}
	return t, nil
}
