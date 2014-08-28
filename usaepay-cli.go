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

var Usage = func(flags *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, "Usage of %s [command]:\n", os.Args[0])
	flags.PrintDefaults()
	os.Exit(0)
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
	inPath := flags.String("in", "", "grab input from file instead of stdin")
	out := flags.String("out", "", "write output to file instead of stdout")
	debug := flags.Bool("debug", false, "debug mode")

	if len(os.Args) > 1 {
		flags.Parse(os.Args[2:])
	}

	// Command Required
	if cmd == "" {
		Usage(flags)
	}
	
	// Required Flags
	if *location == "" || *key == "" || *pin == "" {
		Usage(flags)
	}

	// Input
	var in []byte
	var err error
	if *inPath == "" {
		in, err = ioutil.ReadAll(os.Stdin)
		if err != nil { log.Panic(err.Error()) }
	} else {
		in, err = ioutil.ReadFile(*inPath)
		if err != nil { log.Panic(err.Error()) }
	}

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
	case "createBatchUpload":
		req = new(usaepay.CreateBatchUploadRequest)
		req.SetToken(token)
		_, err = usaepay.JSONToXML(req, in)
		err := req.PostProcess()
		body, err = usaepay.JSONToXML(req, in)
		if err != nil { log.Panic(err.Error()) }
		res = new(usaepay.CreateBatchUploadResponse)
	}

	if *debug { log.Println(string(body)) }

	httpReq, err := usaepay.NewRequest(*location, string(body))
	if err != nil { log.Panic(err.Error()) }
	fullBody, err := res.Handle(httpReq)
	if err != nil { log.Panic(err.Error()) }
	b, err := res.Decode(fullBody)
	if err != nil { log.Panic(err.Error()) }
	// write whole the body
	if *out == "" {
		os.Stdout.Write(b)
	} else {
		err = ioutil.WriteFile(*out, b, 0644)
		if err != nil { log.Panic(err) }
	}
}
