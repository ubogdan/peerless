package peerless

import (
	"context"
	"encoding/xml"
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

func TestService_FailActivateSOA(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte(`<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
	<S:Fault xmlns:ns4="http://www.w3.org/2003/05/soap-envelope">
		<faultcode>EP0003</faultcode>
		<faultstring>ERROR_WHILE_VALIDATING_ORDER_NUMBERS</faultstring>
	</S:Fault>
		</S:Body>
		</S:Envelope>`))
	})

	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()

	_, err := api.ActivateSOA(context.Background(), []string{"100"})
	if err == nil {
		t.Fatalf("AddNotes should fail")
	}
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

	_, err := api.AddNotes(context.Background(), "100", "Test Note")
	if err != nil {
		t.Fatalf("AddNotes: %s", err)
	}
}

func TestService_FailAddNotes(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte(`<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
	<S:Fault xmlns:ns4="http://www.w3.org/2003/05/soap-envelope">
		<faultcode>EP0003</faultcode>
		<faultstring>ERROR_WHILE_VALIDATING_ORDER_NUMBERS</faultstring>
	</S:Fault>
		</S:Body>
		</S:Envelope>`))
	})

	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()

	_, err := api.AddNotes(context.Background(), "100", "Test Note")
	if err == nil {
		t.Fatalf("AddNotes should fail")
	}
}

func TestService_CreateException(t *testing.T) {

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := struct {
			XMLName xml.Name `xml:"Envelope"`
			Body    struct {
				CreateException createException `xml:"createException,omitempty"`
			}
		}{}

		decoder := xml.NewDecoder(r.Body)

		err := decoder.Decode(&request)
		if err != nil {
			t.Fatalf("decode: %s", err)
		}

		if request.Body.CreateException.ExceptionNote.Subject == "" {
			w.WriteHeader(500)
			w.Write([]byte(`<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
		<S:Fault xmlns:ns4="http://www.w3.org/2003/05/soap-envelope">
	<faultcode>EP0003</faultcode>
	<faultstring>ERROR_WHILE_VALIDATING_ORDER_NUMBERS</faultstring>
	</S:Fault>
		</S:Body>
		</S:Envelope>`))
			return
		}

		w.Write([]byte(`<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:createExceptionResponse xmlns:ns2="http://publicapi.api.s2.peerless.com/">
         <return>true</return>
      </ns2:createExceptionResponse>
   </S:Body>
</S:Envelope>`))
	})

	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()

	exception := &ExceptionNote{
		EmailID:       "test@domain.com",
		ExceptionNote: "Exception",
		Subject:       "Exception Subject",
	}

	_, err := api.CreateException(context.Background(), exception)
	if err != nil {
		t.Fatalf("CreateException: %s", err)
	}
}

func TestService_Download(t *testing.T) {

}

func TestService_GetHierarchicalView(t *testing.T) {

}

func TestService_GetNewNumberSearchFilters(t *testing.T) {

}

func TestService_GetOrderStatus(t *testing.T) {

}

func TestService_GetOrdersByPONSearch(t *testing.T) {

}

func TestService_GetPortInRelatedOrders(t *testing.T) {

}

func TestService_GetStatusByNumberSearch(t *testing.T) {

}

func TestService_GetTnInventoryReport(t *testing.T) {

}

func TestService_PlaceOrder(t *testing.T) {

}

func TestService_DisconnectOrder(t *testing.T) {

}

func TestService_PlaceTFOrder(t *testing.T) {

}

func TestService_PlaceTFDisconnectOrder(t *testing.T) {

}

func TestService_PortabilityCheck(t *testing.T) {

}

func TestService_SearchNumbers(t *testing.T) {

}

func TestService_SearchOrderDetailsByOrderId(t *testing.T) {

}

func TestService_SupplementOrder(t *testing.T) {

}

func TestService_Upload(t *testing.T) {

}

func TestService_ValidateE911Address(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := &struct {
			XMLName xml.Name `xml:"Envelope"`
			Body    struct {
				ValidateE911Address validateE911Address `xml:"validateE911Address"`
			}
		}{}
		decoder := xml.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			t.Fatalf("decode: %s", err)
		}
		if request.Body.ValidateE911Address.Address.City == "New York" {
			w.Write([]byte(`<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
		<S:Body>
			<ns2:validateE911AddressResponse xmlns:ns2="http://publicapi.api.s2.peerless.com/"/>
		</S:Body>
			</S:Envelope>`))
			return
		}
		w.WriteHeader(500)
		w.Write([]byte(`<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
		<S:Fault xmlns:ns4="http://www.w3.org/2003/05/soap-envelope">
	<faultcode>EO0001</faultcode>
	<faultstring>FAIL</faultstring>
	</S:Fault>
		</S:Body>
		</S:Envelope>`))
	})

	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()

	_, err := api.ValidateE911Address(context.Background(), BaseAddress{City: "New York"})
	if err != nil {
		t.Fatalf("ValidateE911Address: %s", err)
	}

	_, err = api.ValidateE911Address(context.Background(), BaseAddress{})
	if err == nil {
		t.Fatalf("ValidateE911Address: expecting error")
	}

}

func TestService_ViewNumberDetails(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := &struct {
			XMLName xml.Name `xml:"Envelope"`
			Body    struct {
				ViewNumberDetails viewNumberDetails `xml:"viewNumberDetails"`
			}
		}{}

		decoder := xml.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			t.Fatalf("decode: %s", err)
		}
		if request.Body.ViewNumberDetails.TelephoneNumber[0] == "2010000000" {
			w.Write([]byte(`
			<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
		<S:Body>
			<ns2:viewNumberDetailsResponse xmlns:ns2="http://publicapi.api.s2.peerless.com/">
		<return>
			<tn>2010000000</tn>
			<routeLabel>CUSTOMER_01</routeLabel>
			<customName>CUSTOMER_01</customName>
			<cnamDelivery>false</cnamDelivery>
			<cnamStorage>false</cnamStorage>
			<peerlessMsgProvisioning>false</peerlessMsgProvisioning>
			<sms>false</sms>
			<e911>false</e911>
			</return>
			</ns2:viewNumberDetailsResponse>
		</S:Body>
			</S:Envelope>`))
			return
		}
		w.WriteHeader(500)
		w.Write([]byte(`<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
	<S:Body>
		<S:Fault xmlns:ns4="http://www.w3.org/2003/05/soap-envelope">
	<faultcode>EPCO0005</faultcode>
	<faultstring>ERROR_NUMBERS_DONOT_BELONG_TO_ENTITY</faultstring>
	</S:Fault>
		</S:Body>
		</S:Envelope>`))
	})

	api, srvShutdown := testingHTTPClient(h)
	defer srvShutdown()

	_, err := api.ViewNumberDetails(context.Background(), []string{"2010000000"})
	if err != nil {
		t.Fatalf("ViewNumberDetails: %s", err)
	}

	_, err = api.ViewNumberDetails(context.Background(), []string{"5550000000"})
	if err == nil {
		t.Fatalf("ViewNumberDetails: expecting error")
	}
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
