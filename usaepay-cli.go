package main

import (
	"log"
	"fmt"
	"flag"
	"io/ioutil"	
//	"encoding/xml"
	"usaepay-cli/usaepay"
	"os"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s [command]:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	flags := flag.NewFlagSet(cmd, flag.ExitOnError)

	location := flags.String("location", "", "usaepay endpoint")
	key := flags.String("key", "", "gateway source key")
	pin := flags.String("pin", "", "gateway pin")
	debug := flags.Bool("debug", false, "debug mode")

	flags.Parse(os.Args[2:])

	// Read req file
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil { log.Panic(err.Error()) }
	token := usaepay.NewToken(*key, *pin)

	var req usaepay.Request
	var res usaepay.Response
	var body []byte
	switch cmd {
	case "getTransactionReport":
		req = new(usaepay.GetTransactionReportRequest)
		req.SetToken(token)
		body, err = usaepay.JSONToXML(req, in)
		if err != nil { log.Panic(err.Error()) }
		res = new(usaepay.GetTransactionReportResponse)
	case "searchTransactionsCustom":
		req = new(usaepay.SearchTransactionsCustomRequest)
		req.SetToken(token)
		body, err = usaepay.JSONToXML(req, in)
		if err != nil { log.Panic(err.Error()) }
		res = new(usaepay.SearchTransactionsCustomResponse)
	case "searchCustomers":
		req = new(usaepay.SearchCustomersRequest)
		req.SetToken(token)
		body, err = usaepay.JSONToXML(req, in)
		if err != nil { log.Panic(err.Error()) }
		res = new(usaepay.SearchCustomersResponse)
	case "runTransaction":
		req = new(usaepay.RunTransactionRequest)
		req.SetToken(token)
		body, err = usaepay.JSONToXML(req, in)
		if err != nil { log.Panic(err.Error()) }
		res = new(usaepay.RunTransactionResponse)
	}

	if *debug { log.Println(string(body)) }

	httpReq, err := usaepay.NewRequest(*location, string(body))
	if err != nil { log.Panic(err.Error()) }
	fullBody, err := res.Handle(httpReq)
	if err != nil { log.Panic(err.Error()) }
	b, err := res.Decode(fullBody)
	if err != nil { log.Panic(err.Error()) }
	// write whole the body
	os.Stdout.Write(b) 
}
