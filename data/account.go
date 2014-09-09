// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package data

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	UID     string
	Name    string
	Plan    string
	Users   []bson.ObjectId
	Created time.Time
	Updated time.Time
}

func (a *Account) Create(c *mgo.Collection) error {
	a.ID = bson.NewObjectId()
	a.Created = time.Now()
	a.Updated = time.Now()
	a.Users = []bson.ObjectId{}
	a.UID = a.GenerateUID(c)

	return c.Insert(a)
}

func (a *Account) GenerateUID(c *mgo.Collection) string {
	candidate := RandomAlNum(8)

	// Verify it's unique
	n, err := c.Find(&Account{UID: candidate}).Count()
	if err != nil {
		panic(err)
	}
	// If we had a result -> try again
	if n > 0 {
		return a.GenerateUID(c)
	}

	return candidate
}
