package peerless

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
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

func (s *service) call(ctx context.Context, operation request) (*response, error) {

	payload, err := marshallRequest(operation)
	if err != nil {
		return nil, err
	}

	cli := s.Config
	if cli == nil {
		cli = http.DefaultClient
	}
	request, err := http.NewRequest(http.MethodPost, s.URL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx) //for cancellation support (if our client cancels, we cancel)
	// Set Headers
	if s.UserAgent != "" {
		request.Header.Add("User-Agent", s.UserAgent)
	}

	request.Header.Set("Accept", "text/xml, multipart/related")
	request.Header.Add("SOAPAction", "")
	request.Header.Set("Content-Type", "text/xml; charset=utf-8")

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
		XMLName xml.Name `xml:""`
		Body    *response
	}{Body: &response{}}

	decoder := xml.NewDecoder(resp.Body)
	//decoder.CharsetReader = charset.NewReaderLabel
	return marshalStructure.Body, decoder.Decode(&marshalStructure)
}

func marshallRequest(operation request) ([]byte, error) {
	// Envelope is a SOAP envelope.
	requestBody := struct {
		XMLName      xml.Name    `xml:"SOAP-ENV:Envelope"`
		EnvelopeAttr string      `xml:"xmlns:SOAP-ENV,attr"`
		NSAttr       string      `xml:"xmlns:ns,attr"`
		TNSAttr      string      `xml:"xmlns:tns,attr,omitempty"`
		URNAttr      string      `xml:"xmlns:urn,attr,omitempty"`
		XSIAttr      string      `xml:"xmlns:xsi,attr,omitempty"`
		Header       interface{} `xml:"SOAP-ENV:Header"`
		Body         request     `xml:"SOAP-ENV:Body"`
	}{
		NSAttr:  Namespace,
		XSIAttr: "http://www.w3.org/2001/XMLSchema-instance",
		Body:    operation,
	}

	if requestBody.EnvelopeAttr == "" {
		requestBody.EnvelopeAttr = "http://schemas.xmlsoap.org/soap/envelope/"
	}

	return xml.MarshalIndent(requestBody, "", "  ")
}
