package usaepay

import (
	"encoding/xml"
	"time"
	"testing"
	"log"
)

func TestNewEnvelope(t *testing.T) {
	e := NewEnvelope("test")
	m, _ := xml.MarshalIndent(e, "", "   ")
	log.Println(string(m))
}

func TestNewRequest(t *testing.T) {
	token := NewToken("Z4P2L0O4hW96bKeEvb6YB7j76Z12aBsg", "healpayreports")
	start := time.Now()
	end := time.Now()
	report := "check:settled by date"
	format := "csv"
	query := NewGetTransactionReportRequest(token, start, end, report, format)
	body, err := xml.MarshalIndent(query, "", "   ")
	req, err := NewRequest("https://www.usaepay.com/soap/gate/A47DE151", string(body))
	if err != nil { log.Panic(err.Error()) }
	HandleResponse(req, "report.csv")
}