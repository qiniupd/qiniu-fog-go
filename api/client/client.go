package client

import (
	"net/http"
	"time"

	"github.com/qiniupd/qiniu-fog-go/api/auth/qiniumac.v1"
	"github.com/qiniupd/qiniu-go-sdk/x/rpc.v7"
)


func NewQiniuAuthRPCClient(ak, sk string, timeout time.Duration, header http.Header) *rpc.Client {
	return &rpc.Client{
		Client: &http.Client{
			Timeout: timeout,
			Transport: qiniumac.NewTransport(
				&qiniumac.Mac{AccessKey: ak, SecretKey: []byte(sk)},
				http.DefaultTransport,
				header,
			),
		},
	}
}
