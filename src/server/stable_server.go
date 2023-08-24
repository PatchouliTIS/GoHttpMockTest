package server

import (
	"fmt"
	"net/http"
	"strings"
	"zeus/src/util"

	"google.golang.org/protobuf/proto"
)

var (
	Rpc_name string

	// proto.Message只能在当前源文件内引用，跨文件引用会变成protoreflect.Message；
	// 尝试使用 interface{} 和断言的方式，跨文件获取。
	// StructMap = map[string]proto.Message{
	// 	"AddFeasReq":       &yz.AddFeasReq{SessionId: &client.Session_id},
	// 	"AddFeasRsp":       &yz.AddFeasRsp{SessionId: &client.Session_id},
	// 	"RetrieveReq":      &yz.RetrieveReq{SessionId: &client.Session_id},
	// 	"RetrieveRsp":      &yz.RetrieveRsp{SessionId: &client.Session_id},
	// 	"TruncateGroupReq": &yz.TruncateGroupReq{SessionId: &client.Session_id},
	// 	"CreateGroupReq": &yz.CreateGroupReq{
	// 		SessionId:     &client.Session_id,
	// 		GroupId:       &client.Group_id,
	// 		Platform:      &client.CPU_Platform,
	// 		FeatureConfig: &client.Feature,
	// 	},
	// 	"GetGroupDetailReq": &yz.GetGroupDetailReq{
	// 		SessionId:  &client.Session_id,
	// 		GroupId:    &client.Group_id,
	// 		FeatureIdx: &client.FeatureIdx,
	// 	},
	// }
)

// func CreateGroup(w http.ResponseWriter, r *http.Request) {
// 	body := make([]byte, r.ContentLength)
// 	r.Body.Read(body)
// 	req := yz.CreateGroupReq{}
// 	proto.Unmarshal(body, &req)

// 	rsp := yz.CreateGroupRsp{
// 		SessionId: req.SessionId,
// 		Errorcode: &client.Error_code,
// 		Errormsg:  &client.Error_msg,
// 	}
// 	mockRspBytes, _ := proto.Marshal(&rsp)

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/x-protobuf")
// 	w.Write(mockRspBytes)

// 	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\nNested Struct:%s\nResponse Msg:%s",
// 		r.Method, r.URL.Path, r.URL.EscapedPath(), &req, &rsp)
// }

// func AddFeas(w http.ResponseWriter, r *http.Request) {
// 	body := make([]byte, r.ContentLength)
// 	r.Body.Read(body)
// 	req := yz.AddFeasReq{}
// 	proto.Unmarshal(body, &req)

// 	rsp := yz.AddFeasRsp{
// 		SessionId: req.SessionId,
// 		Errorcode: &client.Error_code,
// 		Errormsg:  &client.Error_msg,
// 	}
// 	mockRspBytes, _ := proto.Marshal(&rsp)

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/x-protobuf")
// 	w.Write(mockRspBytes)

// 	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\nResponse Msg:%s",
// 		r.Method, r.URL.Path, r.URL.EscapedPath(), &rsp)
// }

// func TruncateGroup(w http.ResponseWriter, r *http.Request) {
// 	body := make([]byte, r.ContentLength)
// 	r.Body.Read(body)
// 	req := yz.TruncateGroupReq{}
// 	proto.Unmarshal(body, &req)

// 	rsp := yz.TruncateGroupRsp{
// 		SessionId: req.SessionId,
// 		Errorcode: &client.Error_code,
// 		Errormsg:  &client.Error_msg,
// 	}
// 	mockRspBytes, _ := proto.Marshal(&rsp)

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/x-protobuf")
// 	w.Write(mockRspBytes)

// 	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\nResponse Msg:%s",
// 		r.Method, r.URL.Path, r.URL.EscapedPath(), &rsp)
// }

// func GetGroupDetail(w http.ResponseWriter, r *http.Request) {
// 	body := make([]byte, r.ContentLength)
// 	r.Body.Read(body)
// 	req := yz.GetGroupDetailReq{}
// 	proto.Unmarshal(body, &req)

// 	rsp := yz.GetGroupDetailRsp{
// 		SessionId: req.SessionId,
// 		Errorcode: &client.Error_code,
// 		Errormsg:  &client.Error_msg,
// 		Items: []*yz.FeaItem{
// 			{
// 				FeaId:     &client.FeaId,
// 				EntityId:  &client.EntityId,
// 				Feature_0: []byte("114"),
// 				Feature_1: []byte("514"),
// 			},
// 			{
// 				FeaId:     &client.FeaId,
// 				EntityId:  &client.EntityId,
// 				Feature_0: []byte("1919"),
// 				Feature_1: []byte("810"),
// 			},
// 		},
// 	}
// 	mockRspBytes, _ := proto.Marshal(&rsp)

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/x-protobuf")
// 	w.Write(mockRspBytes)

// 	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\nResponse Msg:%s",
// 		r.Method, r.URL.Path, r.URL.EscapedPath(), &rsp)
// }

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
