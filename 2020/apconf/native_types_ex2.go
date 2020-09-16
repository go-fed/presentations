package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
)

func main() {
	example158 := `
	{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"content": "Thank you @sally for all your hard work! #givingthanks"
	}`

	var note vocab.ActivityStreamsNote
	noteCallback := func(c context.Context, n vocab.ActivityStreamsNote) error {
		note = n
		return nil
	}
	var jsonld map[string]interface{}
	err := json.Unmarshal([]byte(example158), &jsonld)
	if err != nil {
		panic(err)
	}

	r, err := streams.NewJSONResolver(noteCallback)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	err = r.Resolve(ctx, jsonld)
	if err != nil {
		panic(err)
	}
	// BEGIN NTE1 OMIT
	// Get "content" property
	content := note.GetActivityStreamsContent()

	// Inspect its length:
	fmt.Println("Len:", content.Len())

	// Iterate across all of its values
	for iter := content.Begin(); iter != content.End(); iter = iter.Next() {

		// Is it an IRI?
		fmt.Println("IsIRI:", iter.IsIRI())

		// Get a string value
		fmt.Println("String (xml-flavor):", iter.GetXMLSchemaString())
	}
	// END NTE1 OMIT
}
