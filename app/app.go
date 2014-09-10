// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"log"
	"net/url"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
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
	blogUrl   string
)

func Serve(port int, dbUrl string) {
	apiUrl = util.Getenv("API_URL")
	appUrl = util.Getenv("APP_URL")
	webUrl = util.Getenv("WEB_URL")
	webUrl = util.Getenv("BLOG_URL")

	appServer = util.NewServer()
	// Set up templates
	tBox := rice.MustFindBox("templates")
	sBox := rice.MustFindBox("static")
	appServer.RegisterTemplateLayout("l", "_login", tBox)
	appServer.RegisterTemplateLayout("s", "_signup", tBox)
	appServer.RegisterTemplateLayout("a", "_app", tBox)
	// Set up static files & 404
	appServer.Register404(tBox, "404.html")
	appServer.RegisterStaticHandler(sBox, "/static")

	// Open databse session
	var err error
	dbSession, err = mgo.Dial(dbUrl)
	if err != nil {
		panic(err)
	}
	data.Index(dbSession.DB(""))
	appServer.Database = dbSession
	appServer.DatabaseName = dbSession.DB("").Name

	// -----
	// App Routes
	// -----
	registerAccountHandlers(appServer.Router, appServer)
	registerSignupHandlers(appServer.Router, appServer)
	registerAuthHandlers(appServer.Router, appServer)
	registerHomeHandlers(appServer.Router, appServer)

	// Start the engines!
	fmt.Println("K ---> App running on port:", port)
	log.Fatal(appServer.Serve(port))
}

func Protect(handler util.HandlerFunc) util.HandlerFunc {
	return func(c *util.Context) error {
		if c.User != nil {
			return handler(c)
		} else {
			c.Redirect(c.RouteUrl("login") + "?next=" +
				url.QueryEscape(c.Request.URL.String()))
			return nil
		}
	}
}
