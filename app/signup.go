// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/kiasaki/kauthie/data"
	"github.com/kiasaki/kauthie/util"
	"github.com/stripe/stripe"
	"gopkg.in/mgo.v2/bson"
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
	}
	user.SetPassword(password)
	err := user.Create(c.C("users"))
	if err != nil {
		log.Print(err)
		return err
	}

	// Now create his account
	account := data.Account{
		Name: accountName,
		Plan: plan,
	}
	account.Users = []bson.ObjectId{user.ID}
	err = account.Create(c.C("accounts"))
	if err != nil {
		log.Print(err)
		return err
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
	newCustomer, err := stripeClient.Customers.Create(customer)

	// Save new imformation (stripe id) and add account id to user accounts
	user.StripeId = newCustomer.Id
	user.Accounts = []bson.ObjectId{account.ID}
	c.C("users").UpdateId(user.ID, user)

	if err != nil {
		log.Print(err)
		return err
	}

	c.Redirect(c.RouteUrl("login"))
	return nil
}
