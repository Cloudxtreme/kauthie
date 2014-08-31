// Copyright © 2014 Frederic Gingras <frederic@gingras.cc>.
//
// Use of this source code is governed by an BSD-2-Clause
// license that can be found in the LICENSE file.

package admin

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/pat"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Oh, hai! Here lies admin interface"))
}

func Serve(port int, dbUrl string) {
	router := pat.New()
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}

	router.Get("/", homeHandler)

	fmt.Println("K ---> Admin running on port:", port)
	log.Fatal(server.ListenAndServe())
}
