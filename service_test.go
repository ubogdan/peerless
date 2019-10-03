package peerless

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_ActivateSOA(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:activateSOAResponse xmlns:ns2="http://publicapi.api.s2.peerless.com/">
         <return>true</return>
      </ns2:activateSOAResponse>
   </S:Body>
</S:Envelope>`))
	})

	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()

	res, err := api.ActivateSOA(context.Background(), []string{"100"})
	if err != nil {
		t.Fatalf("ActivateSOA: %s", err)
	}
	t.Logf("ActivateSOA: %v", res)
}

func TestService_AddNotes(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:addNotesResponse xmlns:ns2="http://publicapi.api.s2.peerless.com/">
         <return>true</return>
      </ns2:addNotesResponse>
   </S:Body>
</S:Envelope>`))
	})

	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()

	res, err := api.AddNotes(context.Background(), "100", "Test Note")
	if err != nil {
		t.Fatalf("AddNotes: %s", err)
	}
	t.Logf("AddNotes: %v", res)
}

func testingHTTPClient(handler http.Handler) (APIService, func()) {
	s := httptest.NewServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
	}

	return &service{
		URL:    "http://dummy.peerless.com",
		Config: cli,
	}, s.Close
}
