package usaepay

import (
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"bytes"
	"log"
)

type Body struct {
	XMLName xml.Name `xml:"SOAP-ENV:Body"`
	Content string `xml:",innerxml"`
}

type Envelope struct {
	XMLName xml.Name `xml:"SOAP-ENV:Envelope"`
	SoapEnv string `xml:"xmlns:SOAP-ENV,attr"`
	Ns1 string `xml:"xmlns:ns1,attr"`
	Xsd string `xml:"xmlns:xsd,attr"`
	Xsi string `xml:"xmlns:xsi,attr"`
	SoapEnc string `xml:"xmlns:SOAP-ENC,attr"`
	EncodingStyle string `xml:"SOAP-ENV:encodingStyle,attr"`
	Body Body
}

func NewEnvelope(body string) *Envelope {
	return &Envelope{
		Xsd: "http://www.w3.org/2001/XMLSchema",
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Ns1: "urn:usaepay",
		Xsi: "http://www.w3.org/2001/XMLSchema-instance",
		SoapEnc: "http://schemas.xmlsoap.org/soap/encoding/",
		EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: Body{Content: body},
	}
}

func NewRequest(location, body string) (*http.Request, error) {
	envelope := NewEnvelope(body)
	m, err := xml.MarshalIndent(envelope, "", "   ")
	if err != nil { return nil, err }
	log.Println(string(m))
	buffer := bytes.NewBufferString(string(m))
	return http.NewRequest("POST", location, buffer)
}

func HandleResponse(req *http.Request, outFile string) {
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil { log.Println(err.Error()) }
	if resp.StatusCode != 200 { 
		log.Println(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	}

	repRes, err := NewGetTransactionReportResponse(resp.Body)
	if err != nil { log.Panic(err.Error()) }
	b, err := repRes.Decode()
	if err != nil { log.Panic(err.Error()) }
	// write whole the body
	err = ioutil.WriteFile(outFile, b, 0644)
	if err != nil { panic(err) }
}