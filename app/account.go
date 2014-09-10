// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"github.com/gorilla/mux"
	"github.com/kiasaki/kauthie/util"
)

func registerAccountHandlers(r *mux.Router, s *util.Server) {
	r.Handle("/dashboard", s.Handle(Protect(dashboardHandler))).Methods("GET").Name("dashboard")
}

func dashboardHandler(c *util.Context) error {
	return c.T("a", "account/settings").Execute(c.Writer, map[string]interface{}{
		"context": c,
	})
}
