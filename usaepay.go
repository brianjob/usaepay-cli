package main

import (
	"github.com/hoisie/mustache"
	"bytes"
	"net/http"
	"io/ioutil"
	"strconv"
	"io"
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"encoding/base64"
	"time"
	"log"
)

type Token struct {
	ClientIP string
	SourceKey string
	Type string
	Pin string
	unixNano string
}

func (t *Token) UnixNano() string {
	if t.unixNano == "" {
		t.unixNano = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return t.unixNano
}

func (t *Token) Seed() string {
	return t.UnixNano()
}

func (t *Token) HashValue() string {
	h := sha1.New()
	io.WriteString(h, t.SourceKey)
	io.WriteString(h, t.Seed())
	io.WriteString(h, t.Pin)
	return hex.EncodeToString(h.Sum(nil))
}

func (t *Token) TokenString() string {
	body := mustache.RenderFile("token.xml", t)
	return body
}

type GetTransactionReport struct {
	StartDate string
	EndDate string
	Report string
	Format string
	Token
}

type Response struct {
	XMLName xml.Name `xml:"Envelope"`
	Body string `xml:"Body>getTransactionReportResponse>getTransactionReportReturn"`
}

func main() {
	token := &Token{
		ClientIP: "192.168.0.1",
		SourceKey: "",
		Pin: "",
		Type: "sha1",
	}
	
	report := &GetTransactionReport{
		StartDate: "07/27/2014",
		EndDate: "07/27/2014",
		Report: "check:settled by date",
		Format: "csv",
		Token: *token,
	}
	reqBody := mustache.RenderFileInLayout("getTransactionReport.xml", "envelope.xml", report)
	log.Println(reqBody)
	buffer := bytes.NewBufferString(reqBody)
	url := ""
	client := http.Client{}
	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		log.Panic(err.Error())
	}

	resp, err := client.Do(req)
/*	if err != nil {
		log.Panic(err.Error())
	}
	if resp.StatusCode != 200 {
		log.Panic(resp.Status)
	}*/

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { log.Panic(err.Error()) }
	r := &Response{}
	xml.Unmarshal(body, &r)
	data, err := base64.URLEncoding.DecodeString(r.Body)
	if err != nil { log.Panic(err.Error()) }
	log.Println(string(data))
}
