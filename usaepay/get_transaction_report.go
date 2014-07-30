package usaepay

import (
	"encoding/xml"
	"encoding/base64"
	"io"
	"io/ioutil"
	"time"
)

const (
	DateFormat = "01/02/2006"
)

type GetTransactionReportRequest struct {
	XMLName xml.Name `xml:"ns1:getTransactionReport"`
	Token *Token
	StartDate string
	EndDate string
	Report string
	Format string
}

func NewGetTransactionReportRequest(token *Token, start time.Time, end time.Time, report string, format string) *GetTransactionReportRequest {
	return &GetTransactionReportRequest{
		Token: token,
		StartDate: start.Format(DateFormat),
		EndDate: end.Format(DateFormat),
		Report: report,
		Format: format,
	}
}

type GetTransactionReportResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body string `xml:"Body>getTransactionReportResponse>getTransactionReportReturn"`
}

func (res *GetTransactionReportResponse) Decode() ([]byte, error) {
	return base64.URLEncoding.DecodeString(res.Body)
}

func (res *GetTransactionReportResponse) DecodeString() (string, error) {
	d, err := res.Decode()
	return string(d), err
}

func NewGetTransactionReportResponse(reqBody io.ReadCloser) (*GetTransactionReportResponse, error) {
	body, err := ioutil.ReadAll(reqBody)
	if err != nil { return nil, err }
	res := &GetTransactionReportResponse{}
	err = xml.Unmarshal(body, res)
	if err != nil { return nil, err }
	return res, nil
}
