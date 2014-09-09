// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package data

import (
	"time"

	"code.google.com/p/go.crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Email    string
	Fullname string
	Password []byte
	StripeId string `omitempty`
	Accounts []bson.ObjectId
	Created  time.Time
	Updated  time.Time
}

func (u *User) Create(c *mgo.Collection) error {
	u.ID = bson.NewObjectId()
	u.Accounts = []bson.ObjectId{}
	u.Created = time.Now()
	u.Updated = time.Now()

	return c.Insert(u)
}

// SetPassword takes a plaintext password and hashes it with bcrypt and sets the
// password field to the hash.
func (u *User) SetPassword(password string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err) //this is a panic because bcrypt errors on invalid costs
	}
	u.Password = hpass
}

// Login validates and returns a user object if they exist in the database.
func Login(d *mgo.Database, username, password string) (u *User, err error) {
	err = d.C("users").Find(bson.M{"username": username}).One(&u)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		u = nil
	}
	return
}
