// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package util

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/kiasaki/kauthie/data"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Context struct {
	Database *mgo.Database
	Session  *sessions.Session
	User     *data.User
	Request  *http.Request
	Writer   http.ResponseWriter
}

func (c *Context) Close() {
	c.Database.Session.Close()
}

//C is a convenience function to return a collection from the context database.
func (c *Context) C(name string) *mgo.Collection {
	return c.Database.C(name)
}

func (s Server) NewContext(w http.ResponseWriter, req *http.Request) (*Context, error) {
	session, err := s.SessionStore.Get(req, s.SessionName)
	ctx := &Context{
		Database: s.Database.Clone().DB(s.DatabaseName),
		Session:  session,
		Request:  req,
		Writer:   w,
	}
	if err != nil {
		// There was no session, do not bother fetching user
		return ctx, err
	}

	// Try to fill in the user from the session
	if uid, ok := session.Values["user"].(bson.ObjectId); ok {
		err = ctx.C("users").Find(bson.M{"_id": uid}).One(&ctx.User)
	}

	return ctx, err
}
