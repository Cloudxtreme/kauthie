// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package util

import (
	"net/http"

	"thegoods.biz/httpbuf"
)

type RouteHandler struct {
	Server  *Server
	Handler func(*Context) error
}

type HandlerFunc func(*Context) error

func (h *RouteHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Create the context
	ctx, err := h.Server.NewContext(w, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ctx.Close()

	// Run the handler and grab the error, and report it
	buf := new(httpbuf.Buffer)
	ctx.Writer = buf
	err = h.Handler(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the session
	if err = ctx.Session.Save(req, buf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Apply the buffered response to the writer
	buf.Apply(w)
}

func (s *Server) Handle(handler func(*Context) error) *RouteHandler {
	return &RouteHandler{
		Server:  s,
		Handler: handler,
	}
}
