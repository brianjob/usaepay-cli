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

func (res *GetTransactionReportResponse) Decode(respBody []byte) ([]byte, error) {
	err := xml.Unmarshal(respBody, res)
	if err != nil { return nil, err }
	return Base64Decode(res.Body)
}
