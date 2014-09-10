// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package main

import (
	"encoding/gob"
	"flag"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/kiasaki/kauthie/admin"
	"github.com/kiasaki/kauthie/app"
	"github.com/kiasaki/kauthie/util"
	"github.com/kiasaki/kauthie/work"
	"gopkg.in/mgo.v2/bson"
)

var (
	DatabaseUrl = flag.String("database", util.Getenv("MONGOHQ_URL", util.Getenv("DATABASE_URL", "mongodb://localhost:27017/kauthie")), "Mongo database url")
	Port        = flag.Int("port", 1337, "Port to start the app server on")
	AdminPort   = flag.Int("admin_port", 1334, "Port to start the admin server on")

	StartApp   = flag.Bool("app", false, "Should kauthie start her app server?")
	StartWork  = flag.Bool("work", false, "Should kauthie start some workers?")
	StartAdmin = flag.Bool("admin", false, "Should kauthie start her admin server?")
)

func init() {
	gob.Register(bson.ObjectId(""))
}

func main() {
	godotenv.Load()
	flag.Parse()

	if *StartApp {
		go app.Serve(*Port, *DatabaseUrl)
	}
	if *StartWork {
		go work.Work(*DatabaseUrl)
	}
	if *StartAdmin {
		go admin.Serve(*AdminPort, *DatabaseUrl)
	}
	if !*StartApp && !*StartWork && !*StartAdmin {
		go app.Serve(*Port, *DatabaseUrl)
		go work.Work(*DatabaseUrl)
		go admin.Serve(*AdminPort, *DatabaseUrl)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}
