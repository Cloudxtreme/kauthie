// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	mgo "gopkg.in/mgo.v2"
)

type Server struct {
	SessionStore *sessions.CookieStore
	SessionName  string
	Database     *mgo.Session
	DatabaseName string
	Router       *mux.Router
	Server       *http.Server
	Layouts      map[string]CachedLayout
}

func NewServer() *Server {
	authKey := []byte(Getenv("AUTH_KEY"))
	encKey := []byte(Getenv("ENC_KEY"))

	return &Server{
		SessionStore: sessions.NewCookieStore(authKey, encKey),
		SessionName:  "kauthie-app",
		Router:       mux.NewRouter(),
		Layouts:      make(map[string]CachedLayout),
	}
}

func (s *Server) Serve(port int) error {
	s.Server = &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: s.Router,
	}
	return s.Server.ListenAndServe()
}

func (s *Server) RegisterTemplateLayout(shorthand string, layout string, box *rice.Box) {
	funcs := template.FuncMap{
		"routeurl": s.RouteUrl,
	}

	s.Layouts[shorthand] = CachedLayout{
		Box:       box,
		Layout:    layout,
		Functions: funcs,
	}
}

func (s *Server) Register404(box *rice.Box, filename string) {
	s.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fileHandle, err := box.Open(filename)
		if err != nil {
			panic(err)
		}

		http.ServeContent(w, r, filename, box.Time(), fileHandle)
	})
}

func (s *Server) RegisterStaticHandler(box *rice.Box, pathPrefix string) {
	s.Router.PathPrefix(pathPrefix).Handler(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fileServer := http.FileServer(box.HTTPBox())
			http.StripPrefix(pathPrefix, fileServer).ServeHTTP(w, r)
		},
	))
}

func (s *Server) RouteUrl(name string, things ...interface{}) string {
	// Convert the things to strings
	strs := make([]string, len(things))
	for i, th := range things {
		strs[i] = fmt.Sprint(th)
	}

	// Grab the route
	route := s.Router.GetRoute(name)
	if route != nil {
		url, err := route.URL(strs...)
		if err != nil {
			panic(err)
		}

		return url.Path
	}
	panic(errors.New("Could not find named route '" + name + "' in router."))
}
