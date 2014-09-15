package main

import (
	"log"
	"fmt"
	"flag"
	"bytes"
	"io/ioutil"
	"usaepay-cli/usaepay"
	"os"
)

var Usage = func(flags *flag.FlagSet) {
	fmt.Fprintf(os.Stderr, "Usage of %s [command]:\n", os.Args[0])
	flags.PrintDefaults()
	os.Exit(0)
}

func Error(err error, errPath* string) {
	if err == nil { return }
	if *errPath != "" {
		err = ioutil.WriteFile(*errPath, []byte(err.Error()), 0644)
		if err != nil { panic(err) }
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
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
	errPath := flags.String("error", "", "write errors to file instead of stderr (excluding USAePay errors)")
	debug := flags.Bool("debug", false, "debug mode")
	legacy := flags.Bool("legacy", false, "support legacy JSON API. (Only Applies to runTransction)")

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
		if err != nil { Error(err, errPath) }
	} else {
		in, err = ioutil.ReadFile(*inPath)
		if err != nil { Error(err, errPath) }
	}

	token := usaepay.NewToken(*key, *pin)

	var req usaepay.Request
	var res usaepay.Response
	var body bytes.Buffer
	var legacyBody []byte

	if *legacy != true {
		body.WriteString("<ns1:")
		body.WriteString(cmd)
		body.WriteString(">\n")
	}

	switch cmd {
	case "getTransactionReport":
		req = new(usaepay.GetTransactionReportRequest)
		req.SetToken(token)
		res = new(usaepay.GetTransactionReportResponse)
	case "searchTransactionsCustom":
		req = new(usaepay.SearchTransactionsCustomRequest)
		req.SetToken(token)
		res = new(usaepay.SearchTransactionsCustomResponse)
	case "searchCustomers":
		req = new(usaepay.SearchCustomersRequest)
		req.SetToken(token)
		res = new(usaepay.SearchCustomersResponse)
	case "createBatchUpload":
		req = new(usaepay.CreateBatchUploadRequest)
		req.SetToken(token)
		res = new(usaepay.CreateBatchUploadResponse)
	case "runTransaction":
		if *legacy {
			req = new(usaepay.RunTransactionRequest)
			req.SetToken(token)
			legacyBody, err = usaepay.JSONToXML(req, in)
			if err != nil { Error(err, errPath) }
		}
		res = new(usaepay.RawResponse)
	default:
		res = new(usaepay.RawResponse)
	}

	if *legacy != true {
		body.Write(in)
		body.WriteString(token.XMLString())
		body.WriteString("\n</ns1:")
		body.WriteString(cmd)
		body.WriteString(">")
	} else {
		body.Write(legacyBody)
	}

	if *debug { log.Println(body.String()) }

	httpReq, err := usaepay.NewRequest(*location, body.String())
	if err != nil { Error(err, errPath) }
	fullBody, err := res.Handle(httpReq)
	if err != nil { Error(err, errPath) }
	b, err := res.Decode(fullBody)
	if err != nil { Error(err, errPath) }
	// write whole the body
	if *out == "" {
		os.Stdout.Write(b)
	} else {
		err = ioutil.WriteFile(*out, b, 0644)
		if err != nil { Error(err, errPath) }
	}
}
