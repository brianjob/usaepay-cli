package main

import (
	"log"
	"flag"
	"io/ioutil"	
//	"encoding/xml"
	"usaepay-cli/usaepay"
	"os"
)

func main() {
	location := flag.String("location", "", "usaepay endpoint")
	action := flag.String("action", "", "API request action")
	key := flag.String("key", "", "gateway source key")
	pin := flag.String("pin", "", "gateway pin")
	reqFile := flag.String("req", "", "path to request file (json)")
	outFile := flag.String("out", "", "path to output file")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Read req file
	reqData, err := ioutil.ReadFile(*reqFile)
	if err != nil { log.Panic(err.Error()) }

	token := usaepay.NewToken(*key, *pin)

	var req usaepay.Request
	var res usaepay.Response
	var body []byte
	if (*action == "getTransactionReport") {
		req = new(usaepay.GetTransactionReportRequest)
		req.SetToken(token)
		body, err = usaepay.JSONToXML(req, reqData)
		if err != nil { log.Panic(err.Error()) }
		res = new(usaepay.GetTransactionReportResponse)
	} else if (*action == "searchTransactionsCustom") {
		req = new(usaepay.SearchTransactionsCustomRequest)
		req.SetToken(token)
		body, err = usaepay.JSONToXML(req, reqData)
		if err != nil { log.Panic(err.Error()) }
		res = new(usaepay.SearchTransactionsCustomResponse)
	}

	if *debug { log.Println(string(body)) }

	httpReq, err := usaepay.NewRequest(*location, string(body))
	if err != nil { log.Panic(err.Error()) }
	fullBody, err := res.Handle(httpReq)
	if err != nil { log.Panic(err.Error()) }
	b, err := res.Decode(fullBody)
	// write whole the body
	err = ioutil.WriteFile(*outFile, b, 0644)
	if err != nil { panic(err) }
}
