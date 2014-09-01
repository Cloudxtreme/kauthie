// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package util

import (
	"github.com/gorilla/sessions"
	mgo "gopkg.in/mgo.v2"
)

type Server struct {
	SessionStore *sessions.CookieStore
	SessionName  string
	Database     *mgo.Session
	DatabaseName string
}
