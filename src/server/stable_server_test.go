package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	yz "zeus/api/youtu_zeus"
	"zeus/src/common"

	"google.golang.org/protobuf/proto"
)

func TestMain(m *testing.M) {
	// 本地服务器注册对应的 RESTful API 接口
	http.HandleFunc("/ZeusService/CreateGroup", HandleRequst)
	http.HandleFunc("/ZeusService/AddFeas", HandleRequstJSON)
	http.HandleFunc("/ZeusService/TruncateGroup", HandleRequst)
	http.HandleFunc("/ZeusService/GetGroupDetail", HandleRequst)
	go http.ListenAndServe(":14000", nil)
	time.Sleep(1 * time.Second)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateGroup(t *testing.T) {
	// Create a sample RequestPost
	var data []byte
	// 直接从预设好的结构体哈希表中获取，不用自己建
	reqStruc := common.StructMap["CreateGroupReq"].(*yz.CreateGroupReq)
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

	// 获取空白 ResponseRecorder 以供 ServeHTTP使用
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequst)

	handler.ServeHTTP(rr, req)

	// 处理响应状态码
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("CreateGroupHandler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// 判断与预期是否相符
	expectedSessionId := common.Session_id

	rspData := rr.Body.Bytes()
	// rspData := io.ReadAll(rr.Body)
	rsp := &yz.CreateGroupRsp{}
	err = proto.Unmarshal(rspData, rsp)
	if err != nil {
		t.Error(err)
	}

	if *rsp.SessionId != expectedSessionId {
		t.Errorf("Handler returned unexpected body: got %v want %v", *rsp.SessionId, expectedSessionId)
	}
}

func TestGetGroupDetail(t *testing.T) {
	// Create a sample RequestPost
	var data []byte
	// 直接从预设好的结构体哈希表中获取，不用自己建
	reqStruc := common.StructMap["GetGroupDetailReq"].(*yz.GetGroupDetailReq)
	data, err := proto.Marshal(reqStruc)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", "/ZeusService/GetGroupDetail", bytes.NewReader(data))
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
		t.Errorf("GetGroupDetailHandler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// 判断与预期是否相符
	expectedSessionId := common.Session_id

	rspData := rr.Body.Bytes()
	rsp := &yz.GetGroupDetailRsp{}
	err = proto.Unmarshal(rspData, rsp)
	if err != nil {
		t.Error(err)
	}

	if *rsp.SessionId != expectedSessionId {
		t.Errorf("Handler returned unexpected body: got %v want %v", *rsp.SessionId, expectedSessionId)
	}
}

/*
json发送，proto接收是错误的，发送方和接收方必须使用同一种传输协议。
protoc编译生成的结构体具有proto和json两种格式的注释。
*/
func TestAddFeasJSON(t *testing.T) {
	// Create a sample RequestPost
	var data []byte
	// 直接从预设好的结构体哈希表中获取，不用自己建
	reqStruc := common.StructMap["GetGroupDetailReq"].(*yz.GetGroupDetailReq)
	data, err := json.Marshal(reqStruc)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", "/ZeusService/GetGroupDetail", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	// 设置数据格式
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRequstJSON)

	handler.ServeHTTP(rr, req)

	// 处理响应状态码
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetGroupDetailHandler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// 判断与预期是否相符
	expectedSessionId := common.Session_id

	rspData := rr.Body.Bytes()
	rsp := &yz.GetGroupDetailRsp{}
	fmt.Println(">>>JSON<<<")
	err = json.Unmarshal(rspData, rsp)
	fmt.Println(">>>JSON DONE<<<")
	if err != nil {
		t.Error(err)
	}

	if *rsp.SessionId != expectedSessionId {
		t.Errorf("Handler returned unexpected body: got %v want %v", *rsp.SessionId, expectedSessionId)
	}

	fmt.Printf("Response:\n%s", rsp)

}
