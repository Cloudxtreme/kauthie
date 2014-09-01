// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/kiasaki/kauthie/data"
	"github.com/kiasaki/kauthie/util"
	"gopkg.in/mgo.v2"
)

var (
	appServer *util.Server
	appRouter *mux.Router
	dbSession *mgo.Session
	apiUrl    string
	appUrl    string
	webUrl    string
)

var (
	templateFuncs template.FuncMap = template.FuncMap{
		"reverse": reverse,
	}
	LoginT func(string) *template.Template
	AppT   func(string) *template.Template
)

func Serve(port int, dbUrl string) {
	apiUrl = util.Getenv("API_URL")
	appUrl = util.Getenv("APP_URL")
	webUrl = util.Getenv("WEB_URL")
	authKey := []byte(util.Getenv("AUTH_KEY"))
	encKey := []byte(util.Getenv("ENC_KEY"))

	// Set up templates
	tBox := rice.MustFindBox("templates")
	LoginT = util.TemplateForLayout(tBox, "_login", templateFuncs)
	AppT = util.TemplateForLayout(tBox, "_app", templateFuncs)

	// Set up router
	appRouter = mux.NewRouter()
	httpServer := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: appRouter,
	}

	// Open databse session
	var err error
	dbSession, err = mgo.Dial(dbUrl)
	if err != nil {
		panic(err)
	}
	data.Index(dbSession.DB(""))

	appServer = &util.Server{
		SessionStore: sessions.NewCookieStore(authKey, encKey),
		SessionName:  "kauthie-app",
		Database:     dbSession,
		DatabaseName: dbSession.DB("").Name,
	}
	s := appServer
	r := appRouter

	// -----
	// App Routes
	// -----
	registerAuthHandlers(r, s)
	registerHomeHandlers(r, s)

	// Start the engines!
	fmt.Println("K ---> App running on port:", port)
	log.Fatal(httpServer.ListenAndServe())
}

func reverse(name string, things ...interface{}) string {
	// Convert the things to strings
	strs := make([]string, len(things))
	for i, th := range things {
		strs[i] = fmt.Sprint(th)
	}

	// Grab the route
	route := appRouter.GetRoute(name)
	if route != nil {
		url, err := route.URL(strs...)
		if err != nil {
			panic(err)
		}

		return url.Path
	}
	panic(errors.New("Could not find named route '" + name + "' in router."))
}
