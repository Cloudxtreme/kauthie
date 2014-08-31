// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
)

func staticRoute(c *gin.Context) {
	static, err := rice.FindBox("static")
	if err != nil {
		log.Fatal(err)
	}

	original := c.Request.URL.Path
	c.Request.URL.Path = c.Params.ByName("filepath")
	http.FileServer(static.HTTPBox()).ServeHTTP(c.Writer, c.Request)
	c.Request.URL.Path = original
}

func loadTemplates(list ...string) *template.Template {
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}

	templates := template.New("")

	for _, x := range list {
		templateString, err := templateBox.String(x + ".html")
		if err != nil {
			log.Fatal(err)
		}

		// get file contents as string
		_, err = templates.New(x).Parse(templateString)
		if err != nil {
			log.Fatal(err)
		}
	}

	funcMap := template.FuncMap{
		"title": func(a string) string { return strings.Title(a) },
	}

	templates.Funcs(funcMap)

	return templates
}
