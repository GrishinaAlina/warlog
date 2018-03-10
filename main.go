// Copyright 2018 Shestakov Denis. All rights reserved.
// Use of this source code is governed by a Lil
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

type Context struct {
}

func (this *Context) Root(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(rw, "URL = %s", req.URL)
}

func main() {
	var ctx Context
	router := web.New(ctx).Get("/", (*Context).Root)
	http.ListenAndServe(":8000", router)
}
