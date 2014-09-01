// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"os"
)

func Getenv(params ...string) string {
	if len(params) == 0 {
		panic(errors.New("No params passed to util.Getenv"))
	}

	value := os.Getenv(params[0])
	if value != "" {
		return value
	}

	// Panic if we have no default else return default
	if len(params) == 1 {
		panic(errors.New("There is no environment variable named: " + params[0]))
	} else {
		return params[1]
	}
}
