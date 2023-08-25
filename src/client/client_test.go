package client

import (
	"io"
	"net/http"
	"os"
	"testing"
	"time"
	yz "zeus/api/youtu_zeus"
	"zeus/src/common"

	"google.golang.org/protobuf/proto"
)

var (
	client = &http.Client{}
	config = &common.Config{}
)

func TestMain(m *testing.M) {
	// 本地服务器注册对应的 RESTful API 接口
	http.HandleFunc("http://127.0.0.1:14000/ZeusService/CreateGroup", HandleRequst)
	http.HandleFunc("http://127.0.0.1:14000/ZeusService/AddFeas", HandleRequst)
	http.HandleFunc("http://127.0.0.1:14000/ZeusService/TruncateGroup", HandleRequst)
	http.HandleFunc("http://127.0.0.1:14000/ZeusService/GetGroupDetail", HandleRequst)
	go http.ListenAndServe(":14000", nil)
	time.Sleep(1 * time.Second)

	// 初始化 Client 客户端
	config = &common.Config{}
	config.LoadConfig()
	client = NewClient()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestClientAddFeas(t *testing.T) {
	// 发出请求，接收数据
	addFeasReq, err := NewHTTPRequest(config, "AddFeas")
	if err != nil {
		panic(err)
	}

	respBody, err := client.Do(addFeasReq)
	if err != nil {
		panic(err)
	}

	respBytes, err := io.ReadAll(respBody.Body)
	if err != nil {
		panic(err)
	}

	// TODO: 使用Client的方法没法自定义响应处理逻辑
	resp, ok := common.StructMap["AddFeasRsp"].(*yz.AddFeasRsp)
	if ok {
		proto.Unmarshal(respBytes, resp)
		t.Logf("deserialize data success: session_id -> %s", *resp.SessionId)
	} else {
		t.Errorf("deserialize data failed: AddFeasRsp")
	}

}
