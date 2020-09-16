package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
)

func main() {
	// BEGIN NTE1 OMIT
	ex := `
	{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"id": "https://example.com/note",
		"object": [
			"https://example.com/iri",
			{
				"type": "Person",
				"id": "https://example.com/person"
			}
		]
	}`
	// END NTE1 OMIT

	var note vocab.ActivityStreamsNote
	noteCallback := func(c context.Context, n vocab.ActivityStreamsNote) error {
		note = n
		return nil
	}
	var jsonld map[string]interface{}
	err := json.Unmarshal([]byte(ex), &jsonld)
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
	// BEGIN NTE2 OMIT
	fmt.Println("Note id:", note.GetJSONLDId().GetIRI()) // HLfmt
	object := note.GetActivityStreamsObject()

	for iter := object.Begin(); iter != object.End(); iter = iter.Next() {
		if iter.IsIRI() {
			fmt.Println("IRI:", iter.GetIRI()) // HLfmt
		} else if streams.IsOrExtendsActivityStreamsPerson(iter.GetType()) {
			person := iter.GetActivityStreamsPerson()
			fmt.Println("Person ID:", person.GetJSONLDId().GetIRI()) // HLfmt
		}
	}
	// END NTE2 OMIT
}
