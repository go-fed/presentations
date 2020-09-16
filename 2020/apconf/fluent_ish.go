package main

func main() {
	example158 := `
	{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Note",
		"content": "Thank you @sally for all your hard work! #givingthanks"
	}`

	// BEGIN FI1 OMIT
	// Thin wrapper
	i := streams.NewInterpreter(example158) // HLctr

	// Each call inspects the bytes, each could generate an error
	content, err := i.GetContent()
	iri, err := i.GetContent().GetIRI() // Error!

	// Further uses of `i` may lazily discover malformed content

	// Also, could permit writing inappropriate code
	current, err := i.GetCurrent()
	// END FI1 OMIT

	// BEGIN FI2 OMIT
	// Typical JSON processing
	var jsonld map[string]interface{}
	err := json.Unmarshal(example158, &jsonld)

	var note vocab.ActivityStreamsNote // Goal: populate
	noteCallback := func(n vocab.ActivityStreamsNote) { /* Set `note`*/ }

	// Ensure jsonld schema is not malformed.
	r, err := streams.NewJSONResolver(noteCallback) // HLctr
	err = r.Resolve(ctx, jsonld) // One-time error
	// END FI2 OMIT
}
