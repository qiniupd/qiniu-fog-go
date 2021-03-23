package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/qiniupd/qiniu-fog-go/api"
)

func main() {
	a := api.NewApi("", "")
	data := []byte("http://qj9mqal37.hn-bkt.clouddn.com/00c")
	ctx := context.Background()
	body := bytes.NewReader(data)
	job, err := a.SendJob(ctx, "POST", "http://filecoin.app-async-gate.qa.qiniu.io/c2u", body, len(data))

	fmt.Println(job, err)
	if err == nil {
		t, err := a.QueryJob(ctx, job)
		fmt.Println(t, err)
		if err == nil {
			fmt.Println(t.State)
		}
	}

}
