package main

import (
	"path/filepath"

	"github.com/gobuffalo/packr"
	"github.com/pkg/errors"
)

type Templ map[string]string

// Load all .html files into an array
func loadTemplates(box packr.Box) (Templ, error) {
	list := map[string]string{}

	err := box.Walk(func(s string, file packr.File) error {
		ext := filepath.Ext(s)
		if ext == ".html" {
			t := box.String(s)
			list[filepath.Base(s)] = t
		}
		return nil
	})
	if err != nil {
		debug("got an error")
		return Templ{}, errors.Wrap(err, "loadTemplates")
	}

	return list, nil
}
