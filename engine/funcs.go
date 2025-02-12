package engine

import (
	"fmt"
	"strings"
	"text/template"
)

func prefixFuncs(prefix string, funcs template.FuncMap) template.FuncMap {
	prefixedFuncs := template.FuncMap{}
	for name, fn := range funcs {
		if strings.HasPrefix(name, prefix) {
			continue
		}
		prefixedName := fmt.Sprintf("%s_%s", prefix, name)
		prefixedFuncs[prefixedName] = fn
	}
	return prefixedFuncs
}
