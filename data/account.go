// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Name    string
	Slug    string
	Plan    string
	Created time.Time
	Updated time.Time
}
