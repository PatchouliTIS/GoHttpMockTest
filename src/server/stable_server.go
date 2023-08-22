package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	yz "zeus/api/youtu_zeus"
	"zeus/src/client"

	"google.golang.org/protobuf/proto"
)

var (
	rpc_name string

	// proto.Message只能在当前源文件内引用，跨文件引用会变成protoreflect.Message
	StructMap = map[string]proto.Message{
		"AddFeasReq":       &yz.AddFeasReq{SessionId: &client.Session_id},
		"AddFeasRsp":       &yz.AddFeasRsp{SessionId: &client.Session_id},
		"RetrieveReq":      &yz.RetrieveReq{SessionId: &client.Session_id},
		"RetrieveRsp":      &yz.RetrieveRsp{SessionId: &client.Session_id},
		"TruncateGroupReq": &yz.TruncateGroupReq{SessionId: &client.Session_id},
		"CreateGroupReq": &yz.CreateGroupReq{
			SessionId:     &client.Session_id,
			GroupId:       &client.Group_id,
			Platform:      &client.CPU_Platform,
			FeatureConfig: &client.Feature,
		},
		"GetGroupDetailReq": &yz.GetGroupDetailReq{
			SessionId:  &client.Session_id,
			GroupId:    &client.Group_id,
			FeatureIdx: &client.FeatureIdx,
		},
	}
)

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	req := yz.CreateGroupReq{}
	proto.Unmarshal(body, &req)

	rsp := yz.CreateGroupRsp{
		SessionId: req.SessionId,
		Errorcode: &client.Error_code,
		Errormsg:  &client.Error_msg,
	}
	mockRspBytes, _ := proto.Marshal(&rsp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(mockRspBytes)

	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\nNested Struct:%s\nResponse Msg:%s",
		r.Method, r.URL.Path, r.URL.EscapedPath(), &req, &rsp)
}

func AddFeas(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	req := yz.AddFeasReq{}
	proto.Unmarshal(body, &req)

	rsp := yz.AddFeasRsp{
		SessionId: req.SessionId,
		Errorcode: &client.Error_code,
		Errormsg:  &client.Error_msg,
	}
	mockRspBytes, _ := proto.Marshal(&rsp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(mockRspBytes)

	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\nResponse Msg:%s",
		r.Method, r.URL.Path, r.URL.EscapedPath(), &rsp)
}

func TruncateGroup(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	req := yz.TruncateGroupReq{}
	proto.Unmarshal(body, &req)

	rsp := yz.TruncateGroupRsp{
		SessionId: req.SessionId,
		Errorcode: &client.Error_code,
		Errormsg:  &client.Error_msg,
	}
	mockRspBytes, _ := proto.Marshal(&rsp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(mockRspBytes)

	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\nResponse Msg:%s",
		r.Method, r.URL.Path, r.URL.EscapedPath(), &rsp)
}

func GetGroupDetail(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	req := yz.GetGroupDetailReq{}
	proto.Unmarshal(body, &req)

	rsp := yz.GetGroupDetailRsp{
		SessionId: req.SessionId,
		Errorcode: &client.Error_code,
		Errormsg:  &client.Error_msg,
		Items: []*yz.FeaItem{
			{
				FeaId:     &client.FeaId,
				EntityId:  &client.EntityId,
				Feature_0: []byte("114"),
				Feature_1: []byte("514"),
			},
			{
				FeaId:     &client.FeaId,
				EntityId:  &client.EntityId,
				Feature_0: []byte("1919"),
				Feature_1: []byte("810"),
			},
		},
	}
	mockRspBytes, _ := proto.Marshal(&rsp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(mockRspBytes)

	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\nResponse Msg:%s",
		r.Method, r.URL.Path, r.URL.EscapedPath(), &rsp)
}

/*
将要发送的数据通过protobuf序列化后，输出到.txt文件下，以便后续测试框架调用
*/
func Body2Txt(rpc string) error {
	file, err := os.Create(rpc + ".txt")
	if err != nil {
		panic(err)
	}

	// 顺带新建基准测试日志文件
	_, err = os.Create(rpc + ".log")
	if err != nil {
		panic(err)
	}

	// new RequestBody
	data, err := proto.Marshal(StructMap[rpc+"Req"])
	if err != nil {
		panic(err)
	}

	// 复写原有文件的数据
	count, err := file.Write(data)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf(">>>Input %s Accomplished, %d has been written...\n", rpc, count)
	}

	return nil
}

func main() {

	flag.StringVar(&rpc_name, "rpc", "", "name of rpc called")
	flag.Parse()
	if rpc_name == "" || len(rpc_name) == 0 {
		panic("called rpc name wrong! Not exists")
	}
	rpc_name = strings.TrimSpace(rpc_name)
	names := strings.Split(rpc_name, ",")

	// 使用 flag 生成 POST body .txt 文件，以供gohttpbench调用
	for _, name := range names {
		Body2Txt(name)
	}

	// 本地服务器注册对应的 RESTful API 接口
	http.HandleFunc("/ZeusService/CreateGroup", CreateGroup)
	http.HandleFunc("/ZeusService/AddFeas", AddFeas)
	http.HandleFunc("/ZeusService/TruncateGroup", AddFeas)
	http.HandleFunc("/ZeusService/GetGroupDetail", GetGroupDetail)
	http.ListenAndServe(":14000", nil)
}
