package peerless

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)


const XSINamespace = "http://www.w3.org/2001/XMLSchema-instance"

// HTTPError is detailed soap http error
type HTTPError struct {
	StatusCode int
	Status     string
	Msg        string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%q: %q", e.Status, e.Msg)
}

func (c *service) call(soapAction string, operation Request) (*Response, error) {

	// Envelope is a SOAP envelope.
	req := struct {
		XMLName      xml.Name    `xml:"SOAP-ENV:Envelope"`
		EnvelopeAttr string      `xml:"xmlns:SOAP-ENV,attr"`
		NSAttr       string      `xml:"xmlns:ns,attr"`
		TNSAttr      string      `xml:"xmlns:tns,attr,omitempty"`
		URNAttr      string      `xml:"xmlns:urn,attr,omitempty"`
		XSIAttr      string      `xml:"xmlns:xsi,attr,omitempty"`
		Header       interface{} `xml:"SOAP-ENV:Header"`
		Body         Request `xml:"SOAP-ENV:Body"`
	}{
		EnvelopeAttr: c.Envelope,
		URNAttr:      c.URNamespace,
		NSAttr:       c.Namespace,
		XSIAttr:      XSINamespace,
		Header:       c.Header,
		Body:         operation,
	}

	if req.EnvelopeAttr == "" {
		req.EnvelopeAttr = "http://schemas.xmlsoap.org/soap/envelope/"
	}
	if req.NSAttr == "" {
		req.NSAttr = c.URL
	}
	var b bytes.Buffer
	err := xml.NewEncoder(&b).Encode(req)
	if err != nil {
		return nil,err
	}

	log.Printf("Req: %s", b.Bytes())

	cli := c.Config
	if cli == nil {
		cli = http.DefaultClient
	}
	r, err := http.NewRequest("POST", c.URL, &b)
	if err != nil {
		return nil,err
	}
	// Set Headers
	if c.UserAgent != "" {
		r.Header.Add("User-Agent", c.UserAgent)
	}
	var actionName string
	ct := c.ContentType
	if ct == "" {
		ct = "text/xml"
	}
	r.Header.Set("Content-Type", ct)

	if c.ExcludeActionNamespace {
		actionName = soapAction
	} else {
		actionName = fmt.Sprintf("%s/%s", c.Namespace, soapAction)
	}
	r.Header.Add("SOAPAction", actionName)

	resp, err := cli.Do(r)
	if err != nil {
		return nil,err
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
		Body    *Response
	}{Body: &Response{}}

	decoder := xml.NewDecoder(resp.Body)
	//decoder.CharsetReader = charset.NewReaderLabel
	return marshalStructure.Body, decoder.Decode(&marshalStructure)
}

