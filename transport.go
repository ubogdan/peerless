package peerless

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPError is detailed soap http error
type HTTPError struct {
	StatusCode int
	Status     string
	Msg        string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%q: %q", e.Status, e.Msg)
}

func (s *service) call(ctx context.Context, soapAction string, operation request) (*response, error) {

	// Envelope is a SOAP envelope.
	payload := struct {
		XMLName      xml.Name    `xml:"SOAP-ENV:Envelope"`
		EnvelopeAttr string      `xml:"xmlns:SOAP-ENV,attr"`
		NSAttr       string      `xml:"xmlns:ns,attr"`
		TNSAttr      string      `xml:"xmlns:tns,attr,omitempty"`
		URNAttr      string      `xml:"xmlns:urn,attr,omitempty"`
		XSIAttr      string      `xml:"xmlns:xsi,attr,omitempty"`
		Header       interface{} `xml:"SOAP-ENV:Header"`
		Body         request     `xml:"SOAP-ENV:Body"`
	}{
		EnvelopeAttr: s.Envelope,
		URNAttr:      s.URNamespace,
		NSAttr:       s.Namespace,
		XSIAttr:      "http://www.w3.org/2001/XMLSchema-instance",
		Header:       s.Header,
		Body:         operation,
	}

	if payload.EnvelopeAttr == "" {
		payload.EnvelopeAttr = "http://schemas.xmlsoap.org/soap/envelope/"
	}
	if payload.NSAttr == "" {
		payload.NSAttr = s.URL
	}
	var b bytes.Buffer
	err := xml.NewEncoder(&b).Encode(payload)
	if err != nil {
		return nil, err
	}

	log.Printf("Req: %s", b.Bytes())

	cli := s.Config
	if cli == nil {
		cli = http.DefaultClient
	}
	request, err := http.NewRequest(http.MethodPost, s.URL, &b)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx) //for cancellation support (if our client cancels, we cancel)
	// Set Headers
	if s.UserAgent != "" {
		request.Header.Add("User-Agent", s.UserAgent)
	}
	var actionName string
	ct := s.ContentType
	if ct == "" {
		ct = "text/xml"
	}
	request.Header.Set("Content-Type", ct)

	if s.ExcludeActionNamespace {
		actionName = soapAction
	} else {
		actionName = fmt.Sprintf("%s/%s", s.Namespace, soapAction)
	}
	request.Header.Add("SOAPAction", actionName)

	resp, err := cli.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// read only the first MiB of the body in error case
		limReader := io.LimitReader(resp.Body, 1024*1024)
		body, _ := ioutil.ReadAll(limReader)
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Msg:        string(body),
		}
	}

	marshalStructure := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    *response
	}{Body: &response{}}

	decoder := xml.NewDecoder(resp.Body)
	//decoder.CharsetReader = charset.NewReaderLabel
	return marshalStructure.Body, decoder.Decode(&marshalStructure)
}
