// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package data

import (
	mgo "gopkg.in/mgo.v2"
)

// Ensure database indexes are respected for given mongo database
func Index(db *mgo.Database) {
	if err := db.C("users").EnsureIndex(mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	}); err != nil {
		panic(err)
	}
}
