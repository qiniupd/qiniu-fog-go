package api

import (
	"context"
	"github.com/qiniupd/qiniu-fog-go/api/client"
	"github.com/qiniupd/qiniu-go-sdk/x/errors.v1"
	"io"
	"net/http"
	"time"
)

type Api struct {
	ak string
	sk string
}

func NewApi(ak, sk string) *Api{
	return &Api{
		ak: ak,
		sk: sk,
	}
}

type JobResponse struct {
	ID    string `json:"id,omitempty"`
	Error *string `json:"error,omitempty"`
}

func (a *Api)SendJob(ctx context.Context, method string, url string, body io.Reader, bodyLength int) (id string, err error){
	cli := client.NewQiniuAuthRPCClient(a.ak, a.sk, time.Minute)
	j := JobResponse{}
	err = cli.CallWith(ctx, &j, method, url, "application/octet-stream", body, bodyLength)
	if err != nil {
		return "", err
	}
	if j.Error != nil {
		return "", errors.New(*j.Error)
	}
	return j.ID, nil
}

type State int

type TaskResult struct {
	StatusCode int         `json:"status_code"`
	Header     http.Header `json:"header"`
	Body       string      `json:"body"`
}

type TaskBody struct {
	Method   string      `json:"method"`
	Path     string      `json:"path"`
	RawQuery string      `json:"raw_query"`
	Header   http.Header `json:"header"`
	Body     string      `json:"body"`
}

type TaskInfo struct {
	// 固定属性
	ID        string    `json:"id"`
	ReqID     string    `json:"req_id"`
	Request   *TaskBody `json:"request"`
	CreatedAt int64     `json:"created_at"`

	// 可变属性
	State    State       `json:"state"`
	Response *TaskResult `json:"response,omitempty"`
	ExecAt   int64       `json:"exec_at,omitempty"`
}

func (a *Api)QueryJob(ctx context.Context, id string) (resp *TaskInfo, err error){
	cli := client.NewQiniuAuthRPCClient(a.ak, a.sk, time.Minute)
	t := TaskInfo{}
	err = cli.CallWithJson(ctx, &t, "GET", "https://async.qiniuapp.com/v1/task/" + id, nil)
	if err != nil {
		return nil, err
	}
	return &t, nil
}