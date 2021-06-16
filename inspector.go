// Copyright (c) 2021, Microsoft Corporation, Sean Hinchee
// Licensed under the MIT License.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/seh-msft/openapi"
)

var (
	inFile = flag.String("i", "", "input file name (if not stdin)")
)

// Inspector is a demonstrative tool showing calling conventions for the 'openapi' module.
func main() {
	flag.Parse()

	var f io.Reader = os.Stdin
	if *inFile != "" {
		file, err := os.Open(*inFile)
		if err != nil {
			fatal("err: could not open input file →", err)
		}
		f = file
	}

	var of io.Writer = os.Stdout

	bof := bufio.NewWriter(of)
	defer bof.Flush()

	api, err := openapi.Parse(f)
	if err != nil {
		fatal("err: api parse failed -", err)
	}

	fmt.Fprint(of, "A path ⇒ ", api.Paths["/some/kind/of/path"], "\n\n")

	fmt.Fprint(of, "A method ⇒ ", api.Paths["/another/path/{someId}/{anotherId}"]["get"], "\n\n")

	fmt.Fprint(of, "Content of an HTTP 200 OK response ⇒ ", api.Paths["/elsewhere/{someId}/dosomething"]["post"].Responses["200"].Content, "\n\n")

	fmt.Fprint(of, "Scheme of a 200 OK JSON response ⇒ ", api.Paths["/account/{account}/products/{sku}"]["get"].Responses["200"].Content["application/json"]["schema"], "\n\n")

	fmt.Fprint(of, "Parameters available for a GET request ⇒ ", api.Paths["/account/{account}/products/{sku}"]["get"].Parameters, "\n\n")

	fmt.Fprint(of, "Body for the POST method request ⇒ ", api.Paths["/account/{account}/groups/add"]["post"].RequestBody, "\n\n\n")

	fmt.Fprint(of, "A scheme definition ⇒ ", api.Components["schemas"]["Some.Kind.Of.Scheme"], "\n\n")
}

// Fatal - end program with an error message and newline
func fatal(s ...interface{}) {
	fmt.Fprintln(os.Stderr, s...)
	os.Exit(1)
}
