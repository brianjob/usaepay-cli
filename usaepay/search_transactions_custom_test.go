package usaepay

import (
	"testing"
	"encoding/xml"
	"log"
)

func TestRequestXML(t *testing.T) {
	token := NewToken("testkey", "testpin")
	param := &SearchParam{
		Field: "amount",
		Type: "gt",
		Value: "0",
	}
	param2 := &SearchParam{
		Field: "created",
		Type: "gt",
		Value: "8/7/2014",
	}
	params := []*SearchParam{param, param2}
	search := &Search{Params: params}
	field := &Field{"Details.Amount"}
	fields := []*Field{field}
	r := &SearchTransactionsCustomRequest{
		MatchAll: true,
		Start: 10,
		Limit: 20,
		Sort: "created",
		Format: "csv",
		Search: search,
		FieldList: fields,
	}
	r.SetToken(token)
	m, _ := xml.MarshalIndent(r, "", "   ")
	log.Println(string(m))
}
