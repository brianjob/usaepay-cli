package main

import (
	"bytes"
	"net/http"
	"log"
	"flag"
	"encoding/json"
	"io/ioutil"
	"usaepay-cli/usaepay"
)

func main() {
	location := flag.String("location", "", "usaepay endpoint")
	key := flag.String("key", "", "gateway source key")
	pin := flag.String("pin", "", "gateway pin")
	reqFile := flag.String("req", "", "path to request file (json)")
	outFile := flag.String("out", "", "path to output file")
	flag.Parse()

	token := &usaepay.Token{
		ClientIP: "192.168.0.1",
		SourceKey: *key,
		Pin: *pin,
		Type: "sha1",
	}

	reqData, err := ioutil.ReadFile(*reqFile)
	if err != nil { log.Panic(err.Error()) }
	reportReq := &usaepay.GetTransactionReportRequest{}
	err = json.Unmarshal(reqData, reportReq)
	if err != nil { log.Panic(err.Error()) }
	reportReq.Token = *token
	
	buffer := bytes.NewBufferString(reportReq.String())
	client := http.Client{}
	req, err := http.NewRequest("POST", *location, buffer)
	if err != nil {
		log.Panic(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil { log.Println(err.Error()) }
	if resp.StatusCode != 200 { log.Println(resp.Status) }
	log.Println(resp)

	repRes, err := usaepay.NewGetTransactionReportResponse(resp.Body)
	if err != nil { log.Panic(err.Error()) }
	b, err := repRes.Decode()
	if err != nil { log.Panic(err.Error()) }
	// write whole the body
	err = ioutil.WriteFile(*outFile, b, 0644)
	if err != nil { panic(err) }
}
