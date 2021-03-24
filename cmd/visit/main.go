package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/qiniupd/qiniu-fog-go/api"
)

func main() {
	a := api.NewApiWithQueryHost("",
		"",
		"http://app-async-gate.qa.qiniu.io/v1/task/")
	data := []byte("http://qj9mqal37.hn-bkt.clouddn.com/00c")
	ctx := context.Background()
	h := http.Header{
		"qiniu1":[]string{"1", "2"},
	}
	body := bytes.NewReader(data)
	job, err := a.SendJob(ctx, "POST", "http://filecoin.app-async-gate.qa.qiniu.io/c2u", body, len(data), h)

	fmt.Println(job, err)
	if err == nil {
		t, err := a.QueryJob(ctx, job)
		fmt.Println(t, err)
		if err == nil {
			fmt.Println(t.State)
		}
	}

}
