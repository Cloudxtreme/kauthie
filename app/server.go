// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
)

func Server(port string) {
	r := gin.Default()

	// Views
	templates := loadTemplates("layouts/login", "layouts/app", "selector")
	r.SetHTMLTemplate(templates)

	// Routes
	r.GET("/", selectorRoute)

	r.GET("/authorize", authorizeRoute)
	r.GET("/token", tokenRoute)

	r.GET("/login", loginRoute)
	r.POST("/login", loginRoute)
	r.GET("/forgot", forgotRoute)
	r.POST("/forgot", forgotRoute)
	r.GET("/recover", recoverRoute)

	accountR := r.Group("/account", auth)
	{
		accountR.GET("/account/:aid/edit", homeRoute)
	}

	r.GET("/static/*filepath", staticRoute)

	// Run
	fmt.Println("K ---> App running on port:", port)
	r.Run(":" + port)
}

// Auth middleware to be used when setting up routes
// Responds with json when the request content-type is json
func auth(c *gin.Context) {
	session, _ := sessionStore.Get(r, sessionName)

	user, err := checkRequestForUser(c, session)
	if err != nil {

	}

	context.Clear(r)

}

// Home page route: redirects to login if not logged in
// Redirects to first account if user has one account
// Renders account selection if more than one account is owned
func selectorRoute(c *gin.Context) {
	// If one account -> redirect to that one

	c.HTML(200, "selector", gin.H{})
}

func loginRoute(c *gin.Context) {

}
