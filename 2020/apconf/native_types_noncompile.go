package main

import (
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
	// BEGIN NTN1 OMIT
	var current vocab.ActivityStreamsCurrentProperty
	note.SetActivityStreamsCurrent(current) // HLerr
	// END NTN1 OMIT
}
