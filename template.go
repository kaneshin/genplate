package genplate

import (
	"io"
	"text/template"

	bundle "github.com/kaneshin/genplate/templates"
)

type (
	// An ImportData represents an import path.
	ImportData struct {
		Alias string
		Path  string
	}
)

var (
	templates *template.Template
)

func init() {
	templates = template.Must(bundle.Parse(nil))
}

// ExecuteImportTemplate executes to parse import.tmpl by given data.
func ExecuteImportTemplate(wr io.Writer, data []ImportData) error {

	return templates.ExecuteTemplate(wr, "import.tmpl", data)
}
