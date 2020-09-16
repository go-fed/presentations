package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-fed/activity/streams"
)

func mustParse(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

func main() {
	attributedTo := streams.NewActivityStreamsAttributedToProperty()
	attributedTo.AppendIRI(mustParse("https://example.com/peyton"))
	assignedTo := streams.NewForgeFedAssignedToProperty()
	assignedTo.SetIRI(mustParse("https://example.com/jessie"))
	content := streams.NewActivityStreamsContentProperty()
	content.AppendXMLSchemaString("I found a bug in my computer!")

	// BEGIN FFE1 OMIT
	ticket := streams.NewForgeFedTicket() // HLff
	ticket.SetActivityStreamsAttributedTo(attributedTo)
	ticket.SetForgeFedAssignedTo(assignedTo)
	ticket.SetActivityStreamsContent(content)
	m, err := streams.Serialize(ticket)
	// END FFE1 OMIT
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
