package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"zeus/api/youtu"
	yz "zeus/api/youtu_zeus"

	"github.com/golang/protobuf/proto"
)

const (
	content_type = "application/x-protobuf"
	ZeusSrv      = "ZeusService"
)

var (
	Error_code       int32             = 0
	Session_id       string            = "114514"
	Group_id         string            = "75th Ranger Regiment"
	Error_msg        string            = ""
	CPU_Platform     youtu.Platform    = 0
	GPU_Platform     youtu.Platform    = 1
	Dimension        int32             = 5
	FeatureType_INT8 youtu.FeatureType = 1
	FeatureIdx       int32             = 0
	Scale            float64           = 1.4038
	FeaId            string            = "YJSP"
	EntityId         string            = "Koumakan"
	Feature                            = yz.FeatureConfig{
		Dimension:   &Dimension,
		FeatureType: &(FeatureType_INT8),
		Scale:       &Scale,
		FeatureIdx:  &FeatureIdx,
	}

	StructMap = map[string]proto.Message{
		"AddFeasReq":       &yz.AddFeasReq{SessionId: &Session_id},
		"AddFeasRsp":       &yz.AddFeasRsp{SessionId: &Session_id},
		"RetrieveReq":      &yz.RetrieveReq{SessionId: &Session_id},
		"RetrieveRsp":      &yz.RetrieveRsp{SessionId: &Session_id},
		"TruncateGroupReq": &yz.TruncateGroupReq{SessionId: &Session_id},
		"CreateGroupReq": &yz.CreateGroupReq{
			SessionId:     &Session_id,
			GroupId:       &Group_id,
			Platform:      &CPU_Platform,
			FeatureConfig: &Feature,
		},
	}
)

// main access logic
func HandleRequst(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	// 通过 Reques 路径 获取类的信息
	path := r.URL.EscapedPath()
	pathSlice := strings.Split(path, "/")
	reqName := pathSlice[len(pathSlice)-1] + "Req"
	rsqName := pathSlice[len(pathSlice)-1] + "Rsp"

	req := StructMap[reqName]
	req.ProtoMessage()
	proto.Unmarshal(body, req)

	// sid := req.String()

	rsp := StructMap[rsqName]

	// rsp := yz.AddFeasRsp{
	// 	SessionId: &sid,
	// 	Errorcode: &Error_code,
	// 	Errormsg:  &Error_msg,
	// }
	mockRspBytes, _ := proto.Marshal(rsp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/x-protobuf")
	w.Write(mockRspBytes)

	fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\n",
		r.Method, r.URL.Path, r.URL.EscapedPath())
}

func Retrieve(api string, req *yz.RetrieveReq) (*yz.RetrieveRsp, error) {
	var err error
	// 1. 将Req体序列化
	reqProto, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(reqProto)

	// 发起 POST Request
	url := fmt.Sprintf("%s/%s/Retrieve", api, ZeusSrv)
	resp, _ := http.Post(url, content_type, body)
	// httpReq, err := http.NewRequest(http.MethodPost, url, body)
	// if err != nil {
	// 	log.Fatalf("New Request Failed:%s", err)
	// }

	if resp.StatusCode != http.StatusOK {
		return &yz.RetrieveRsp{}, fmt.Errorf("response didn't get Status 200 but %s instead", resp.Status)
	}

	// 2. 读取resp.Body()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	rsp := yz.RetrieveRsp{}

	err = proto.Unmarshal(respBody, &rsp)
	if err != nil {
		return &yz.RetrieveRsp{}, fmt.Errorf("decode resp failed")
	}

	return &rsp, nil

}

func AddFeas(api string, req *yz.AddFeasReq) (*yz.AddFeasRsp, error) {
	// 1. 将Req体序列化
	reqProto, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	body := strings.NewReader(string(reqProto))

	url := fmt.Sprintf("%s/%s/AddFeas", api, ZeusSrv)
	resp, _ := http.Post(url, content_type, body)

	if resp.StatusCode != http.StatusOK {
		return &yz.AddFeasRsp{}, fmt.Errorf("response didn't get Status 200 but %s instead", resp.Status)
	}

	// 2. 读取resp.Body()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	rsp := yz.AddFeasRsp{}

	err = proto.Unmarshal(respBody, &rsp)
	if err != nil {
		return &yz.AddFeasRsp{}, fmt.Errorf("decode resp failed")
	}

	return &rsp, nil
}

func CreateGroup(api string, req *yz.CreateGroupReq) (*yz.CreateGroupRsp, error) {
	var err error
	// 1. 将Req体序列化
	reqProto, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(reqProto)

	// 发起 POST Request
	url := fmt.Sprintf("%s/%s/CreateGroup", api, ZeusSrv)
	resp, _ := http.Post(url, content_type, body)
	// httpReq, err := http.NewRequest(http.MethodPost, url, body)
	// if err != nil {
	// 	log.Fatalf("New Request Failed:%s", err)
	// }

	if resp.StatusCode != http.StatusOK {
		return &yz.CreateGroupRsp{}, fmt.Errorf("response didn't get Status 200 but %s instead", resp.Status)
	}

	// 2. 读取resp.Body()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	rsp := yz.CreateGroupRsp{}

	err = proto.Unmarshal(respBody, &rsp)
	if err != nil {
		return &yz.CreateGroupRsp{}, fmt.Errorf("decode resp failed")
	}

	return &rsp, nil
}
