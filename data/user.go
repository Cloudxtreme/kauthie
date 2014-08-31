// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package data

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type Entry struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Timestamp time.Time
	Name      string
	Message   string
}

func NewEntry() *Entry {
	return &Entry{
		Timestamp: time.Now(),
	}
}
