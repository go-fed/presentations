package main

import (
	"net/url"
)

// BEGIN NT1 OMIT
type ActivityStreamsNote interface {
	// "content"
	GetActivityStreamsContent() ActivityStreamsContentProperty
}

// END NT1 OMIT

// BEGIN NT2 OMIT
type ActivityStreamsContentProperty interface {
	// Get an iterator, since there could be more than one
	Begin() ActivityStreamsContentPropertyIterator
	End() ActivityStreamsContentPropertyIterator
}

type ActivityStreamsContentPropertyIterator interface {
	// Return the IRI of the property
	GetIRI() *url.URL
	// Other getters for different value types...
	// Get the content as a "string"
	GetXMLSchemaString() string
}

// END NT2 OMIT
