// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiasaki/kauthie/util"
)

func registerAuthHandlers(r *mux.Router, s *util.Server) {
	r.Handle("/signup", s.Handle(signupHandler)).Methods("GET").Name("signup")
	r.Handle("/login", s.Handle(loginHandler)).Methods("GET").Name("login")
	r.Handle("/logout", s.Handle(logoutHandler)).Methods("GET").Name("logout")
}

func signupHandler(c *util.Context) error {
	return c.T("s", "signup").Execute(c.Writer, map[string]interface{}{
		"pricingUrl": webUrl + "pricing",
		"docsUrl":    webUrl + "docs",
		"blogUrl":    blogUrl,
	})
}

func loginHandler(c *util.Context) error {
	return c.T("l", "login").Execute(c.Writer, map[string]interface{}{
		"context": c,
	})
}

func logoutHandler(c *util.Context) error {
	delete(c.Session.Values, "user")
	http.Redirect(c.Writer, c.Request, c.RouteUrl("login"), http.StatusSeeOther)
	return nil
}
