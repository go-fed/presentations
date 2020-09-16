type myTransport struct { // implements pub.Transport
	signer httpsig.Signer // HLsig
}

func (myTransport) Dereference(c context.Context, iri *url.URL) ([]byte, error) {
	request := http.NewRequest(...)
	// ...
	signer.SignRequest(privKey, pubKeyID, request, nil) // HLsig
	httpClient.Do(request)
}
