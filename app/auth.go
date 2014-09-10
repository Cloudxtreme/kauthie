// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"fmt"
	"net/http"

	"github.com/bluele/gforms"
	"github.com/gorilla/mux"
	"github.com/kiasaki/kauthie/data"
	"github.com/kiasaki/kauthie/util"
)

func registerAuthHandlers(r *mux.Router, s *util.Server) {
	r.Handle("/login", s.Handle(loginHandler)).Methods("GET", "POST").Name("login")
	r.Handle("/logout", s.Handle(logoutHandler)).Methods("GET").Name("logout")
}

type userForm struct {
	Email    string `gforms:"email"`
	Password string `gforms:"password"`
}

var loginForm = gforms.DefineForm(
	gforms.NewFields(
		gforms.NewTextField(
			"username",
			gforms.Validators{
				gforms.Required(),
				gforms.MinLengthValidator(4),
				gforms.EmailValidator(),
			},
		),
		gforms.NewTextField(
			"password",
			gforms.Validators{
				gforms.Required(),
				gforms.MinLengthValidator(3),
				gforms.MaxLengthValidator(16),
			},
			gforms.PasswordInputWidget(map[string]string{}),
		),
	),
)

func loginHandler(c *util.Context) error {
	send := func(invalid bool) error {
		return c.T("l", "login").Execute(c.Writer, map[string]interface{}{
			"context":  c,
			"next":     c.Request.FormValue("next"),
			"username": c.Request.FormValue("username"),
			"invalid":  invalid,
		})
	}

	// Simple get
	if c.Request.Method == "GET" {
		fmt.Println("get")
		return send(false)
	}

	// Form validation pass
	form := loginForm(c.Request)
	if !form.IsValid() {
		fmt.Println("invalid")
		return send(true)
	}

	user, err := data.Login(c.Database,
		c.Request.FormValue("username"),
		c.Request.FormValue("password"))

	// Login failed -> bad creds
	if err != nil {
		fmt.Println(err)
		return send(true)
	}

	// Log him in!
	c.Session.Values["user"] = user.ID
	if c.Request.FormValue("next") != "" {
		c.Redirect(c.Request.FormValue("next"))
	} else {
		c.Redirect("/dashboard")
	}
	return nil
}

func logoutHandler(c *util.Context) error {
	delete(c.Session.Values, "user")
	http.Redirect(c.Writer, c.Request, c.RouteUrl("login"), http.StatusSeeOther)
	return nil
}
