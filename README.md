# gorbl
Real-time Blackhole List (RBL) lookups for Golang.

[![GoDoc](https://godoc.org/github.com/polera/gorbl?status.svg)](https://godoc.org/github.com/polera/gorbl)  [![Build Status](https://travis-ci.org/polera/gorbl.svg?branch=master)](https://travis-ci.org/polera/gorbl)

Author
==
James Polera <james@uncryptic.com>

Dependencies
==
No external dependencies.  Uses Go's standard pacakges

Example
==

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/polera/gorbl"
)

func main() {

	r := gorbl.Lookup("b.barracudacentral.org", "smtp.gmail.com")
	json_res, _ := json.Marshal(r)
	fmt.Printf("%v\n", string(json_res))

	/*
		{
		    "host": "smtp.gmail.com",
		    "list": "b.barracudacentral.org",
		    "results": [
		        {
		            "address": "173.194.206.109",
		            "error": true,
		            "error_type": {
		                "Err": "no such host",
		                "IsTimeout": false,
		                "Name": "109.206.194.173.b.barracudacentral.org",
		                "Server": ""
		            },
		            "listed": false,
		            "text": ""
		        },
		        {
		            "address": "173.194.206.108",
		            "error": true,
		            "error_type": {
		                "Err": "no such host",
		                "IsTimeout": false,
		                "Name": "108.206.194.173.b.barracudacentral.org",
		                "Server": ""
		            },
		            "listed": false,
		            "text": ""
		        }
		    ]
			}
	*/

}

```
