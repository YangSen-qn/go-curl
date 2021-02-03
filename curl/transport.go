package curl

import (
	"net/http"
)

type Transport struct {
	Transport *http.Transport

	CAPath     string
	ForceHTTP3 bool
}

func (t *Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	if t.ForceHTTP3 {
		transport := &http3Transport{CAPath: t.CAPath}
		return transport.RoundTrip(request)
	} else {
		return t.Transport.RoundTrip(request)
	}
}

