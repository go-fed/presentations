// BEGIN ASS1 OMIT
createLogic := func(c context.Context, create vocab.ActivityStreamsCreate) error {
	// Big Boring Business Logic - assume AP logic already done.
	phone := database.FetchPhoneForCreatedMessage(c)
	phone.BananaPhoneRingtone()
}
// END ASS1 OMIT

// BEGIN ASS2 OMIT
// This function satisfies an interface for the middleware to call.
func (myService) FederatingCallbacks(c context.Context) (pub.FederatingWrappedCallbacks, []interface, error) {
	return pub.FederatingWrappedCallbacks{
		Create: createLogic, // HLfun
		// Delete: ...,
		// Update: ...,
		// Like: ...,
		// etc
	}, nil, nil
}
// END ASS2 OMIT

// BEGIN ASS3 OMIT
func (myService) FederatingCallbacks(c context.Context) (pub.FederatingWrappedCallbacks, []interface, error) {
	return pub.FederatingWrappedCallbacks{},
		[]interface{}{createInDBAndBznsLogic}, // HLint
		nil
}
// END ASS3 OMIT
