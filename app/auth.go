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
	r.Handle("/login", s.Handle(loginHandler)).Methods("GET").Name("login")
	r.Handle("/logout", s.Handle(logoutHandler)).Methods("GET").Name("logout")
}

func loginHandler(c *util.Context) error {
	return LoginT("login").Execute(c.Writer, map[string]interface{}{
		"context": c,
	})
}

func logoutHandler(c *util.Context) error {
	delete(c.Session.Values, "user")
	http.Redirect(c.Writer, c.Request, reverse("login"), http.StatusSeeOther)
	return nil
}
