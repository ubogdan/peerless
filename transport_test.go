package peerless

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

// TODO test
func TestService_Call(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
		<S:Body>
			<ns2:viewNumberDetailsResponse xmlns:ns2="http://publicapi.api.s2.peerless.com/">
			</ns2:viewNumberDetailsResponse>
		</S:Body>
			</S:Envelope>`))

	})

	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()

	service, ok := api.(*service)
	if !ok {
		t.Fatalf("Invalid service type")
	}

	_, err := service.call(context.Background(), requestBody{})
	if err != nil {
		t.Fatalf("ViewNumberDetails: %s", err)
	}
}

func TestService_UserAgent(t *testing.T) {
	expectUserAgent := "Custom User Agent"
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") != expectUserAgent {
			t.Fatalf("Expected %s Got %s", expectUserAgent, r.Header.Get("User-Agent"))
		}
	})
	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()

	service, ok := api.(*service)
	if !ok {
		t.Fatalf("Invalid service type")
	}
	service.UserAgent = expectUserAgent
	service.call(context.Background(), requestBody{})
}

func TestService_FaultResponse(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte(`<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
<S:Body><S:Fault xmlns:ns4="http://www.w3.org/2003/05/soap-envelope"><faultcode>EO0000</faultcode>
<faultstring>GENERAL_ERROR</faultstring></S:Fault></S:Body></S:Envelope>`))
	})
	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()
	service, ok := api.(*service)
	if !ok {
		t.Fatalf("Invalid service type")
	}
	_, err := service.call(context.Background(), requestBody{})
	if err != nil {
		faultErr, ok := err.(*Fault)
		if !ok {
			t.Fatalf("Must fail with SOAP FAUL error")
		}
		expected := fmt.Sprintf("%q: %q", faultErr.Code, faultErr.String)
		if faultErr.Error() != expected {
			t.Fatalf("Error expexted `%s` Got `%s`", expected, faultErr.Error())
		}
	}

}

func TestService_InvalidResponse(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte(`500 Internal server error`))
	})

	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()

	service, ok := api.(*service)
	if !ok {
		t.Fatalf("Invalid service type")
	}

	_, err := service.call(context.Background(), requestBody{})
	if err != nil {
		httpErr, ok := err.(*HTTPError)
		if !ok {
			t.Fatalf("Must fail with HTTP error")
		}
		expected := fmt.Sprintf("%q: %q", httpErr.Status, httpErr.Msg)
		if httpErr.Error() != expected {
			t.Fatalf("Error expexted `%s` Got `%s`", expected, httpErr.Error())
		}
	}
}

func TestService_EmptyResponse(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})
	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()
	service, ok := api.(*service)
	if !ok {
		t.Fatalf("Invalid service type")
	}
	_, err := service.call(context.Background(), requestBody{})
	if err != nil {
		httpErr, ok := err.(*HTTPError)
		if !ok {
			t.Fatalf("Must fail with HTTP error")
		}
		if httpErr.Msg != "" {
			t.Fatalf("Empty response expected")
		}
	}
}
