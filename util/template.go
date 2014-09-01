// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package util

import (
	"html/template"
	"sync"

	"github.com/GeertJohan/go.rice"
)

var cachedTemplates = map[string]*template.Template{}
var cachedMutex sync.Mutex

func TemplateForLayout(box *rice.Box, layout string, funcs template.FuncMap) func(string) *template.Template {

	return func(name string) *template.Template {
		cachedMutex.Lock()
		defer cachedMutex.Unlock()

		if t, ok := cachedTemplates[name]; ok {
			return t
		}

		t := template.New(name).Funcs(funcs)

		t.Parse(box.MustString(name + ".html"))
		t.Parse(box.MustString(layout + ".html"))
		cachedTemplates[name] = t

		return t
	}
}
