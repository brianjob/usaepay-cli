package usaepay

import (
	"io"
	"crypto/sha1"
	"strconv"
	"time"
	"encoding/hex"
	"github.com/hoisie/mustache"
)

type Token struct {
	ClientIP string
	SourceKey string
	Type string
	Pin string
	unixNano string
}

func (t *Token) UnixNano() string {
	if t.unixNano == "" {
		t.unixNano = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return t.unixNano
}

func (t *Token) Seed() string {
	return t.UnixNano()
}

func (t *Token) HashValue() string {
	h := sha1.New()
	io.WriteString(h, t.SourceKey)
	io.WriteString(h, t.Seed())
	io.WriteString(h, t.Pin)
	return hex.EncodeToString(h.Sum(nil))
}

func (t *Token) TokenString() string {
	body := mustache.RenderFile("./usaepay/token.xml", t)
	return body
}

