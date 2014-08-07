package usaepay

import (
	"encoding/xml"
	"log"
)
/*
boolean	MatchAll	If set to “true,” only results matching all search criteria will be returned, if set to “false,” results matching any of the search criteria will be returned.
integer	Start	Sequence number to start returning on.
integer	Limit	Maximum number of transactions to return in result set.
string	FieldList	String Array of fields to return in search.
string	Format	Specify format of return data. Possible formats include: csv, tab, xml.
string	Sort	Field name to sort the results by
*/

type Field struct {
	Item string `xml:"item"`
}

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

type SearchTransactionsCustomRequest struct {
	XMLName xml.Name `xml:"ns1:searchTransactionsCustom"`
	MatchAll bool
	Start int
	Limit int
	FieldList []*Field
	Format string
	Sort string
	Search *Search
	USAePayRequest
}

type SearchTransactionsCustomResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body string `xml:"Body>searchTransactionsCustomResponse>searchTransactionsCustomReturn"`
	USAePayResponse
}
func (res *SearchTransactionsCustomResponse) GetBody() string {
	return res.Body
}

func (res *SearchTransactionsCustomResponse) Decode() ([]byte, error) {
	log.Println("OI")
	return Base64Decode(res.Body)
}

func (res *SearchTransactionsCustomResponse) DecodeString() (string, error) {
	d, err := res.Decode()
	return string(d), err
}
