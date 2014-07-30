package usaepay

import (
	"time"
	"testing"
	"encoding/xml"
	"log"
)

func TestNewGetTransactionReportRequest(t *testing.T) {
	token := NewToken("testkey", "testpin")
	start := time.Now()
	end := time.Now()
	report := "some report"
	format := "tab"
	r := NewGetTransactionReportRequest(token, start, end, report, format)
	m, _ := xml.MarshalIndent(r, "", "   ")
	log.Println(string(m))
}