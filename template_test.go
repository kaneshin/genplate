package genplate

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteTemplate(t *testing.T) {
	assert := assert.New(t)

	var buf bytes.Buffer
	assert.NoError(ExecuteImportTemplate(&buf, []ImportData{
		{
			Path: "context",
		},
		{
			Path: "path/filepath",
		},
		{
			Alias: "f",
			Path:  "fmt",
		},
		{
			Alias: ".",
			Path:  "path",
		},
		{
			Alias: "_",
			Path:  "strings",
		},
	}))
	const expected = `import (
	"context"
	"path/filepath"
	f "fmt"
	. "path"
	_ "strings"
)
`
	assert.Equal(expected, buf.String())
}
