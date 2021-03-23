package qiniumac

import (
	"encoding/base64"
	"net/http"

	. "github.com/qiniupd/qiniu-go-sdk/api.v8/conf"
)

// ---------------------------------------------------------------------------------------

type Mac struct {
	AccessKey string
	SecretKey []byte
}

type Transport struct {
	mac       Mac
	Transport http.RoundTripper
}

func (t *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {

	sign, err := SignRequest(t.mac.SecretKey, req)
	if err != nil {
		return
	}

	auth := "Qiniu " + t.mac.AccessKey + ":" + base64.URLEncoding.EncodeToString(sign)
	req.Header.Set("Authorization", auth)
	return t.Transport.RoundTrip(req)
}

func (t *Transport) NestedObject() interface{} {

	return t.Transport
}

func NewTransport(mac *Mac, transport http.RoundTripper) *Transport {

	if transport == nil {
		transport = http.DefaultTransport
	}
	t := &Transport{Transport: transport}
	if mac == nil {
		t.mac.AccessKey = ACCESS_KEY
		t.mac.SecretKey = []byte(SECRET_KEY)
	} else {
		t.mac = *mac
	}
	return t
}

func NewClient(mac *Mac, transport http.RoundTripper) *http.Client {

	t := NewTransport(mac, transport)
	return &http.Client{Transport: t}
}
