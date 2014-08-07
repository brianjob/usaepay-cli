package usaepay

import (
	"encoding/xml"
)

const (
	DateFormat = "01/02/2006"
)

type GetTransactionReportRequest struct {
	XMLName xml.Name `xml:"ns1:getTransactionReport"`
	Token Token
	StartDate string
	EndDate string
	Report string
	Format string
	USAePayRequest
}

type GetTransactionReportResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body string `xml:"Body>getTransactionReportResponse>getTransactionReportReturn"`
	USAePayResponse
}

func (res *GetTransactionReportResponse) GetBody() string {
	return res.Body
}

func (res *GetTransactionReportResponse) Decode() ([]byte, error) {
	return Base64Decode(res.Body)
}

func (res *GetTransactionReportResponse) DecodeString() (string, error) {
	d, err := res.Decode()
	return string(d), err
}
