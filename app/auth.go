// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

var (
	sessionStore sessions.CookieStore
	sessionName  = "kauthie-session"
)

func setUpAuth() {
	authKey = viper.GetString("authkey")
	encKey = viper.GetString("enckey")
	sessionStore = sessions.NewCookieStore([]byte(authKey), []byte(encKey))
}

func checkRequestForUser(c *gin.Context, s *sessions.Session) (User, Error) {
	authorization := c.Request.Header.Get("Authorization")
	userId := session.Values["uid"]

	return nil, nil
}
