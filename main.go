// Copyright 2018 Shestakov Denis. All rights reserved.
// Use of this source code is governed by a Lil
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	//"net/http"
	. "./logparser"
	"github.com/gocraft/web"
	"bytes"
)

type Context struct {
}

func (this *Context) Root(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(rw, "URL = %s", req.URL)
}

func main() {
	testData := make([][]byte, 0)
	for i := 0; i < 5; i++ {
		testData = append(testData, []byte(fmt.Sprintf("%s \n", "klfdsk1111ksldfds")))
	}
	for i := 0; i < 3; i++ {
		testData = append(testData, []byte(fmt.Sprintf("%s \n", "klfdskksldfds")))
	}

	_, err := ParseLogObjects(bytes.Join(testData, nil), "./logparser/test_config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	/*var ctx Context
	router := web.New(ctx).Get("/", (*Context).Root)
	http.ListenAndServe(":8000", router)*/
}
