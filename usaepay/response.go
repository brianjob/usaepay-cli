package usaepay

import (
	"encoding/base64"
	"encoding/xml"
	"net/http"
	"io/ioutil"
)

type Response interface {
	Handle(req *http.Request) (*http.Response, []byte, error)
	Decode([]byte) ([]byte, error)
	GetBody() string
	ToXML()
}

type USAePayResponse struct {
	Response
}

type RawResponse struct {
	XMLName xml.Name `xml:"Envelope"`
        Body string
        USAePayResponse
}

func (res *RawResponse) GetBody() string {
        return res.Body
}

func (res *RawResponse) Decode(respBody []byte) ([]byte, error) {
        return respBody, nil
}

func Base64Decode(body string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(body)
}

func (res *USAePayResponse) Handle(req *http.Request) (*http.Response, []byte, error) {
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil { return nil, nil, err }
	b, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return resp, b, err
	}
	if err != nil { return nil, nil, err }
	return resp, b, err
}
