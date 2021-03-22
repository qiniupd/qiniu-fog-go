package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/qiniupd/qiniu-fog-go/api/client"
)

func main() {

	ctx := context.Background()
	cli := client.NewQiniuAuthRPCClient(
		"<ak>",
		"<sk>",
		time.Minute)

	var ret interface{}
	err := cli.CallWithJson(ctx, &ret,
		// "GET", "http://app-async-gate.qa.qiniu.io/v1/task/01z001c9tgdhmiri7s00ty373r000191",
		"POST", "http://test.local.net:19005/test/?url=http://861h53.com2.z0.glb.qiniucdn.com/upload.jpg&cmd=qhash/qhash/md5",
		nil)

	bts, _ := json.Marshal(ret)
	fmt.Println(string(bts), err)
}
