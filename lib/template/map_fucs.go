package template

import (
	"html/template"

	"github.com/jeffprestes/go-example-vps-locaweb/lib/context"
)

// FuncMaps to view
func FuncMaps() []template.FuncMap {
	return []template.FuncMap{
		map[string]interface{}{
			"Tr": context.I18n,
		}}
}
