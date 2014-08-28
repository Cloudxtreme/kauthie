// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Server(port string) {
	r := gin.Default()

	// Views
	templates := loadTemplates()
	r.SetHTMLTemplate(templates)

	// Routes
	r.GET("/", homeRoute)
	r.GET("/static/*filepath", staticRoute)

	// Run
	fmt.Println("A ---> Running on port:", port)
	r.Run(":" + port)
}

func homeRoute(c *gin.Context) {
	c.String(200, "Oh hai!")
}
