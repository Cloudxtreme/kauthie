// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/kiasaki/kauthie/util"
)

func registerHomeHandlers(r *mux.Router, s *util.Server) {
	r.PathPrefix("/static").Handler(http.HandlerFunc(staticHandler))
	r.Handle("/", s.Handle(homeHandler)).Methods("GET").Name("index")

	r.NotFoundHandler = http.HandlerFunc(fourOhFourHandler)
}

func homeHandler(c *util.Context) error {
	fmt.Println(c.Request.URL.Path)
	http.Redirect(c.Writer, c.Request, reverse("login"), http.StatusSeeOther)
	return nil
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	httpBox := rice.MustFindBox("static").HTTPBox()
	fileServer := http.FileServer(httpBox)
	http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
}

func fourOhFourHandler(w http.ResponseWriter, r *http.Request) {
	box := rice.MustFindBox("templates")
	fileHandle, err := box.Open("404.html")
	if err != nil {
		panic(err)
	}

	http.ServeContent(w, r, "404.html", box.Time(), fileHandle)
}
