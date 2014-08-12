package usaepay

import (
	"encoding/xml"
)

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
