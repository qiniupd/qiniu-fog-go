package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/qiniupd/qiniu-fog-go/api/auth/qboxmac.v1"
	"github.com/qiniupd/qiniu-fog-go/api/auth/qiniumac.v1"
	"github.com/qiniupd/qiniu-go-sdk/x/rpc.v7"
)

type QiniuStubTransport struct {
	UID   uint32
	Utype uint32
	http.RoundTripper
}

func (t QiniuStubTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(
		"Authorization",
		fmt.Sprintf("QiniuStub uid=%d&ut=%d", t.UID, t.Utype),
	)
	return t.RoundTripper.RoundTrip(req)
}

func NewQiniuStubRPCClient(uid, utype uint32, timeout time.Duration) *rpc.Client {
	return &rpc.Client{
		Client: &http.Client{
			Timeout: timeout,
			Transport: QiniuStubTransport{
				UID:          uid,
				Utype:        utype,
				RoundTripper: http.DefaultTransport,
			},
		},
	}
}

func NewQiniuAuthRPCClient(ak, sk string, timeout time.Duration) *rpc.Client {
	return &rpc.Client{
		Client: &http.Client{
			Timeout: timeout,
			Transport: qiniumac.NewTransport(
				&qiniumac.Mac{AccessKey: ak, SecretKey: []byte(sk)},
				http.DefaultTransport,
			),
		},
	}
}

func NewQiniuAdminRPCClient(adminMac *qiniumac.Mac, uid uint32, timeout time.Duration) *rpc.Client {
	return &rpc.Client{
		Client: &http.Client{
			Timeout: timeout,
			Transport: qiniumac.NewAdminTransport(
				adminMac, fmt.Sprintf("%d/0", uid), http.DefaultTransport,
			),
		},
	}
}

func NewQboxAuthRPCClient(ak, sk string, timeout time.Duration) *rpc.Client {
	return &rpc.Client{
		Client: &http.Client{
			Timeout: timeout,
			Transport: qboxmac.NewTransport(
				&qboxmac.Mac{AccessKey: ak, SecretKey: []byte(sk)},
				http.DefaultTransport,
			),
		},
	}
}
