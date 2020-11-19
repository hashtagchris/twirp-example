// Copyright 2018 Twitch Interactive, Inc.  All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the License is
// located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/twitchtv/twirp"
	"github.com/twitchtv/twirp-example/rpc/haberdasher"
)

func main() {
	// Create a client capable of talking to a Haberdasher server running on
	// localhost. This is a generated function call.

	// If the server returns an unknown enum value, the json client will fail with:
	// twirp error internal: failed to unmarshal json response: unknown value "\"FRUIT_LEATHER\"" for enum twitch.twirp.example.haberdasher.Hat.Fabric
	// client := haberdasher.NewHaberdasherJSONClient("http://localhost:8080", &http.Client{})

	client := haberdasher.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})

	var (
		hat *haberdasher.Hat
		err error
	)

	// Call the client's 'MakeHat' method, retrying up to five times.
	for i := 0; i < 5; i++ {
		hat, err = client.MakeHat(context.Background(), &haberdasher.Size{Inches: 12})
		if err != nil {
			// We got an error. Is it a twirp Error?
			if twerr, ok := err.(twirp.Error); ok {
				// Twirp errors support custom, arbitrary metadata. For example, a
				// server could inform a client that a particular error is retryable.
				if twerr.Meta("retryable") != "" {
					log.Printf("got error %q, retrying", twerr)
					continue
				}
			}
			log.Fatal(err)
		} else {
			break
		}
	}

	// Print out the response.
	fmt.Printf("%+v\n", hat)

	switch hat.Fabric {
	case haberdasher.Hat_COTTON:
		fmt.Println("Cotton")
	case haberdasher.Hat_SATIN:
		fmt.Println("Satin")
	case haberdasher.Hat_UNKNOWN:
		fmt.Println("Unknown fabric")
	case haberdasher.Hat_VINYL:
		fmt.Println("Vinyl")
	default:
		fmt.Printf("Can't map %v to a fabric, sorry\n", hat.Fabric)
	}
}
