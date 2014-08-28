// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/pilu/fresh/runner/runnerutils"
)

func Server(port string) {

	r := gin.Default()

	if os.Getenv("DEV") != "" {
		r.Use(RunnerMiddleware())
	}

	templates := loadTemplates()
	r.SetHTMLTemplate(templates)

	r.GET("/", homeRoute)
	r.GET("/static/*filepath", staticServe)
	fmt.Println("Running on port:", port)
	r.Run(":" + port)
}

func staticServe(c *gin.Context) {
	static, err := rice.FindBox("static")
	if err != nil {
		log.Fatal(err)
	}
	original := c.Request.URL.Path
	c.Request.URL.Path = c.Params.ByName("filepath")
	http.FileServer(static.HTTPBox()).ServeHTTP(c.Writer, c.Request)
	c.Request.URL.Path = original
}

func RunnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if runnerutils.HasErrors() {
			runnerutils.RenderError(c.Writer)
			c.Abort(500)
		}
	}
}

func loadTemplates(list ...string) *template.Template {
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}

	templates := template.New("")

	for _, x := range list {
		templateString, err := templateBox.String(x)
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
		"html":  ProperHtml,
		"title": func(a string) string { return strings.Title(a) },
	}

	templates.Funcs(funcMap)

	return templates
}

func ProperHtml(text string) template.HTML {
	if strings.Contains(text, "content:encoded>") || strings.Contains(text, "content/:encoded>") {
		text = html.UnescapeString(text)
	}
	return template.HTML(html.UnescapeString(template.HTMLEscapeString(text)))
}

func homeRoute(c *gin.Context) {
	c.String(200, "Oh hai!")
}
