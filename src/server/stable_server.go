package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"zeus/src/util"

	"google.golang.org/protobuf/proto"
)

var (
	Rpc_name string
)

// main access logic
func HandleRequst(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	// 通过 Reques 路径 获取类的信息
	path := r.URL.EscapedPath()
	pathSlice := strings.Split(path, "/")
	name := pathSlice[len(pathSlice)-1]

	/*
	 插入 GetReqAndRsp
	 使用 interface{} 断言来获取值
	 req := StructMap[reqName]
	*/
	_, rsp := util.GetReqAndRsp(name)
	mockRspBytes, err := proto.Marshal(rsp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(mockRspBytes)

	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\n",
		r.Method, r.URL.Path, r.URL.EscapedPath())
}

// main access logic
func HandleRequstJSON(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	// 通过 Request 路径 获取类的信息
	path := r.URL.EscapedPath()
	pathSlice := strings.Split(path, "/")
	name := pathSlice[len(pathSlice)-1]

	/*
	 插入 GetReqAndRsp
	 使用 interface{} 断言来获取值
	 req := StructMap[reqName]
	*/
	_, rsp := util.GetReqAndRsp(name)
	mockRspBytes, err := json.Marshal(rsp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(mockRspBytes)

	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\n",
		r.Method, r.URL.Path, r.URL.EscapedPath())
}
