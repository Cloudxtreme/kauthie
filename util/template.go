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

type CachedLayout struct {
	Box       *rice.Box
	Layout    string
	Functions template.FuncMap
}

func (cl *CachedLayout) GetTemplate(name string) *template.Template {
	cachedMutex.Lock()
	defer cachedMutex.Unlock()

	if t, ok := cachedTemplates[name]; ok {
		return t
	}

	t := template.New(name).Funcs(cl.Functions)

	t.Parse(cl.Box.MustString(cl.Layout + ".html"))
	t.Parse(cl.Box.MustString(name + ".html"))
	cachedTemplates[name] = t

	return t
}
