package main

import (
	"log"
	"flag"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"usaepay-cli/usaepay"
	"os"
)

func main() {
	location := flag.String("location", "", "usaepay endpoint")
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

	token := usaepay.NewToken(*key, *pin)

	// Read req file
	reqData, err := ioutil.ReadFile(*reqFile)
	if err != nil { log.Panic(err.Error()) }


	reportReq := &usaepay.GetTransactionReportRequest{}
	err = json.Unmarshal(reqData, reportReq)
	if err != nil { log.Panic(err.Error()) }
	reportReq.Token = token

	body, err := xml.MarshalIndent(reportReq, "", "   ")
	if *debug { log.Println(string(body)) }
	req, err := usaepay.NewRequest(*location, string(body))
	if err != nil { log.Panic(err.Error()) }
	usaepay.HandleResponse(req, *outFile)
}
