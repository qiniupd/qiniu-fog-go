package qboxmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"io"
	"net/http"

	"github.com/qiniupd/qiniu-go-sdk/x/bytes.v7/seekable"
)

// ---------------------------------------------------------------------------------------

func incBody(req *http.Request) bool {

	if req.Body == nil || req.ContentLength == 0 {
		return false
	}
	if ct, ok := req.Header["Content-Type"]; ok {
		switch ct[0] {
		case "application/x-www-form-urlencoded":
			return true
		}
	}
	return false
}

func SignRequest(sk []byte, req *http.Request) ([]byte, error) {

	h := hmac.New(sha1.New, sk)

	u := req.URL
	data := u.Path
	if u.RawQuery != "" {
		data += "?" + u.RawQuery
	}
	io.WriteString(h, data+"\n")

	if incBody(req) {
		s2, err2 := seekable.New(req)
		if err2 != nil {
			return nil, err2
		}
		h.Write(s2.Bytes())
	}

	return h.Sum(nil), nil
}

// ---------------------------------------------------------------------------------------

type RequestSigner struct {
}

var (
	DefaultRequestSigner RequestSigner
)

func (p RequestSigner) Sign(sk []byte, req *http.Request) ([]byte, error) {

	return SignRequest(sk, req)
}

// ---------------------------------------------------------------------------------------
