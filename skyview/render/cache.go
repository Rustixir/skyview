package render

import (
	"html/template"
	"io"
	"os"
)

var cacheFs = make(map[string]struct{})
var cache = make(map[string]*template.Template)

func Render(name string, html string, data interface{}, wr io.Writer) (err error) {
	t, ok := cache[name]
	if !ok {
		t, err = template.New(name).Parse(html)
		if err != nil {
			return err
		}
		cache[name] = t
	}

	return t.Execute(wr, data)
}

func RenderFile(path string, data interface{}, wr io.Writer) error {
	_, ok := cacheFs[path]
	if !ok {
		bytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		cacheFs[path] = struct{}{}
		return Render(path, string(bytes), data, wr)
	}

	return Render(path, "", data, wr)
}
