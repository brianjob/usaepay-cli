package usaepay

import (
	"encoding/xml"
)

type Search struct {
	XMLName xml.Name `xml:"Search"`
	Params []*SearchParam
}

type SearchParam struct {
	XMLName xml.Name `xml:"item"`
	Field string
	Type string
	Value string
}

type SearchCustomersRequest struct {
	XMLName xml.Name `xml:"ns1:searchCustomers"`
	MatchAll bool
	Start int
	Limit int
	Sort string
	Search *Search
	USAePayRequest
}

type SearchCustomersResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body string
	USAePayResponse
}

func (res *SearchCustomersResponse) GetBody() string {
	return res.Body
}

func (res *SearchCustomersResponse) Decode(respBody []byte) ([]byte, error) {
	return respBody, nil
}
