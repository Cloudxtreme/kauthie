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

func registerHomeHandlers(r *mux.Router, s *util.Server) {
	r.Handle("/", s.Handle(homeHandler)).Methods("GET").Name("index")
}

func homeHandler(c *util.Context) error {
	http.Redirect(c.Writer, c.Request, c.RouteUrl("login"), http.StatusSeeOther)
	return nil
}
