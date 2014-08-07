package usaepay

import (
	"encoding/base64"
	"net/http"
	"io/ioutil"
	"log"
)

type Response interface {
	Handle(req *http.Request) ([]byte, error)
	Decode([]byte) ([]byte, error)
	GetBody() string
	ToXML() 
}

type USAePayResponse struct {
	Response
}

func Base64Decode(body string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(body)
}

func (res *USAePayResponse) Handle(req *http.Request) ([]byte, error) {
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil { log.Println(err.Error()) }
	if resp.StatusCode != 200 { 
		log.Println(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	}

	return ioutil.ReadAll(resp.Body)
}
