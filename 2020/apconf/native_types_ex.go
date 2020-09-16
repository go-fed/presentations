package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
)

func main() {
	// BEGIN NT1 OMIT
	example158 := `
	{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"content": "Thank you @sally for all your hard work! #givingthanks"
	}`

	// Goal:
	var note vocab.ActivityStreamsNote // HLnote
	// END NT1A OMIT
	fmt.Println(note == nil) // HLnil
	// END NT1B OMIT
	noteCallback := func(c context.Context, n vocab.ActivityStreamsNote) error {
		note = n
		return nil
	}
	// BEGIN NT2 OMIT
	var jsonld map[string]interface{}
	err := json.Unmarshal([]byte(example158), &jsonld)
	// END NT2 OMIT
	if err != nil {
		panic(err)
	}

	// BEGIN NT3 OMIT
	r, err := streams.NewJSONResolver(noteCallback)
	// END NT3 OMIT
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	// BEGIN NT4 OMIT
	err = r.Resolve(ctx, jsonld)
	fmt.Println(note == nil) // HLnil
	// END NT4 OMIT
	if err != nil {
		panic(err)
	}
}
