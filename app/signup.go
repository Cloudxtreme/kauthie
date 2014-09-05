// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"log"
	"time"

	"github.com/gorilla/mux"
	"github.com/kiasaki/kauthie/data"
	"github.com/kiasaki/kauthie/util"
	"github.com/stripe/stripe"
)

func registerSignupHandlers(r *mux.Router, s *util.Server) {
	r.Handle("/signup", s.Handle(signupHandler)).Methods("GET").Name("signup")
	r.Handle("/signup", s.Handle(signupPostHandler)).Methods("POST")
}

func signupHandler(c *util.Context) error {
	selectedPlan := c.Request.FormValue("plan")
	return c.T("s", "signup").Execute(c.Writer, map[string]string{
		"pricingUrl":   webUrl + "pricing",
		"docsUrl":      webUrl + "docs",
		"blogUrl":      blogUrl,
		"selectedPlan": selectedPlan,
	})
}

func signupPostHandler(c *util.Context) error {
	// Setup stripe client
	stripeClient := &stripe.Client{}
	stripeClient.Init(util.Getenv("STRIPE_PRIVATE_KEY"), nil, nil)

	// Gater form values
	fullname := c.Request.FormValue("fullname")
	accountName := c.Request.FormValue("account")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	plan := c.Request.FormValue("plan")
	token := c.Request.FormValue("stripeToken")

	// Store new user in mongo
	user := data.User{
		Email:    email,
		Fullname: fullname,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	user.SetPassword(password)
	err := c.C("users").Insert(user)
	if err != nil {
		log.Fatal(err)
	}

	// Now create his account
	account := data.Account{
		Name: accountName,
	}
	cerr := account.Create(c.C("accounts"))
	if err != nil {
		log.Fatal(cerr)
	}

	// Create and Customer, subscribe him, start trial, associate card (1-step)
	customer := &stripe.CustomerParams{
		Token: token,
		Desc:  fullname,
		Email: email,
		Plan:  plan,
		Params: stripe.Params{
			Meta: map[string]string{
				"uid": user.ID.String(),
			},
		},
	}
	_, err = stripeClient.Customers.Create(customer)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
