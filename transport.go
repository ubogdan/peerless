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

type soapRespone struct {
	XMLName xml.Name     `xml:"Envelope"`
	Header  *soapHeader  `xml:",omitempty"`
	Body    responseBody `xml:",omitempty"`
}

type soapRequest struct {
	XMLName      xml.Name    `xml:"SOAP-ENV:Envelope"`
	EnvelopeAttr string      `xml:"xmlns:SOAP-ENV,attr"`
	NSAttr       string      `xml:"xmlns:ns,attr"`
	TNSAttr      string      `xml:"xmlns:tns,attr,omitempty"`
	URNAttr      string      `xml:"xmlns:urn,attr,omitempty"`
	XSIAttr      string      `xml:"xmlns:xsi,attr,omitempty"`
	Header       interface{} `xml:"SOAP-ENV:Header"`
	Body         requestBody `xml:"SOAP-ENV:Body"`
}

type soapHeader struct {
	XMLName xml.Name    `xml:"Header"`
	Content interface{} `xml:",omitempty"`
}

type Fault struct {
	XMLName xml.Name `xml:"Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}

// HTTPError is detailed soap http error
type HTTPError struct {
	StatusCode int
	Status     string
	Msg        string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%q: %q", e.Status, e.Msg)
}

func (f *Fault) Error() string {
	return fmt.Sprintf("%q: %q", f.Code, f.String)
}

func (s *service) call(ctx context.Context, operation requestBody) (*responseBody, error) {
	envelope := &soapRequest{
		NSAttr:       Namespace,
		XSIAttr:      "http://www.w3.org/2001/XMLSchema-instance",
		EnvelopeAttr: "http://schemas.xmlsoap.org/soap/envelope/",
		Body:         operation,
	}
	payload, err := xml.MarshalIndent(envelope, "", "  ")
	if err != nil {
		return nil, err
	}

	cli := s.Config
	if cli == nil {
		cli = http.DefaultClient
	}

	// Add xml header to payload
	reqWriter := &bytes.Buffer{}
	reqWriter.WriteString(xml.Header)
	reqWriter.Write(payload)

	request, err := http.NewRequest(http.MethodPost, s.URL, reqWriter)
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

	response, err := cli.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return unmarshallResponse(response)
}

func unmarshallResponse(res *http.Response) (*responseBody, error) {
	envelope := &soapRespone{}

	if res.StatusCode != http.StatusOK {
		// read only the first MiB of the body in error case
		limReader := io.LimitReader(res.Body, 1024*1024)
		bodyBytes, _ := ioutil.ReadAll(limReader)

		err := xml.Unmarshal(bodyBytes, &envelope)
		if err != nil {
			// return HTTP Error
			return nil, &HTTPError{
				StatusCode: res.StatusCode,
				Status:     res.Status,
				Msg:        string(bodyBytes),
			}
		}
		// return SOAP Fault
		return nil, envelope.Body.Fault
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &envelope.Body, xml.Unmarshal(bodyBytes, &envelope)
}
