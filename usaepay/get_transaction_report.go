package usaepay

import (
	"encoding/xml"
	"encoding/base64"
	"io"
	"io/ioutil"
	"github.com/hoisie/mustache"
)

type GetTransactionReportRequest struct {
	StartDate string
	EndDate string
	Report string
	Format string
	Token
}

func (req *GetTransactionReportRequest) String() string {
	return mustache.RenderFileInLayout("./usaepay/getTransactionReport.xml", "./usaepay/envelope.xml", req)
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
