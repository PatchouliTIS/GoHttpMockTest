package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	yz "zeus/api/youtu_zeus"
	"zeus/src/client"

	"google.golang.org/protobuf/proto"
)

func TestMain(m *testing.M) {
	// 本地服务器注册对应的 RESTful API 接口
	http.HandleFunc("/ZeusService/CreateGroup", HandleRequst)
	http.HandleFunc("/ZeusService/AddFeas", HandleRequst)
	http.HandleFunc("/ZeusService/TruncateGroup", HandleRequst)
	http.HandleFunc("/ZeusService/GetGroupDetail", HandleRequst)
	go http.ListenAndServe(":14000", nil)
	time.Sleep(2 * time.Second)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateGroup(t *testing.T) {
	// Create a sample RequestPost
	var data []byte
	reqStruc := client.StructMap["CreateGroupReq"].(*yz.CreateGroupReq)
	data, err := proto.Marshal(reqStruc)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", "/ZeusService/CreateGroup", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	// 设置数据格式
	req.Header.Set("Content-Type", "application/x-protobuf")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequst)

	handler.ServeHTTP(rr, req)

	// 处理响应状态码
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("CreateGroupHandler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// 判断与预期是否相符
	expectedSessionId := client.Session_id

	rspData := rr.Body.Bytes()
	rsp := &yz.CreateGroupRsp{}
	err = proto.Unmarshal(rspData, rsp)
	if err != nil {
		t.Error(err)
	}

	if *rsp.SessionId != expectedSessionId {
		t.Errorf("Handler returned unexpected body: got %v want %v", *rsp.SessionId, expectedSessionId)
	}
}
