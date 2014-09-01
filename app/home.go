// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kiasaki/kauthie/util"
)

func registerHomeHandlers(r *mux.Router, s *util.Server) {
	r.Handle("/", s.Handle(homeHandler)).Methods("GET").Name("index")
	r.Handle("/env.js", http.HandlerFunc(envHandler)).Methods("GET").Name("env")
}

func homeHandler(c *util.Context) error {
	http.Redirect(c.Writer, c.Request, c.RouteUrl("login"), http.StatusSeeOther)
	return nil
}

func envHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprintln(w, "var __env = __env || {};")
	fmt.Fprintf(w, "__env.stripePublishableKey = '%s';\n", util.Getenv("STRIPE_PUBLIC_KEY"))
}
