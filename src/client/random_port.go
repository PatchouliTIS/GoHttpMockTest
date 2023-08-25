package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	yz "zeus/api/youtu_zeus"
	"zeus/src/common"

	"google.golang.org/protobuf/proto"
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

	var mockRspBytes []byte
	// 使用 interface{} 断言来获取值
	// req := StructMap[reqName]
	switch reqName {
	case "Retrieve":
		req, ok := common.StructMap[reqName].(*yz.RetrieveReq)
		if ok {
			// 解析获取的POST body
			// req.ProtoMessage()
			proto.Unmarshal(body, req)
		} else {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.RetrieveRsp)
		if ok {
			rsp.SessionId = req.SessionId
			mockRspBytes, _ = proto.Marshal(rsp)
		} else {
			panic("interface assert rsp failed. \n")
		}

	case "AddFeas":
		req, ok := common.StructMap[reqName].(*yz.AddFeasReq)
		if ok {
			// 解析获取的POST body
			// req.ProtoMessage()
			proto.Unmarshal(body, req)
		} else {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.AddFeasRsp)
		if ok {
			rsp.SessionId = req.SessionId
			mockRspBytes, _ = proto.Marshal(rsp)
		} else {
			panic("interface assert rsp failed. \n")
		}
	case "TruncateGroup":
		req, ok := common.StructMap[reqName].(*yz.TruncateGroupReq)
		if ok {
			// 解析获取的POST body
			// req.ProtoMessage()
			proto.Unmarshal(body, req)
		} else {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.TruncateGroupRsp)
		if ok {
			rsp.SessionId = req.SessionId
			mockRspBytes, _ = proto.Marshal(rsp)
		} else {
			panic("interface assert rsp failed. \n")
		}
	case "CreateGroup":
		req, ok := common.StructMap[reqName].(*yz.CreateGroupReq)
		if ok {
			// 解析获取的POST body
			// req.ProtoMessage()
			proto.Unmarshal(body, req)
		} else {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.CreateGroupRsp)
		if ok {
			rsp.SessionId = req.SessionId
			mockRspBytes, _ = proto.Marshal(rsp)
		} else {
			panic("interface assert rsp failed. \n")
		}
	case "GetGroupDetail":
		req, ok := common.StructMap[reqName].(*yz.GetGroupDetailReq)
		if ok {
			// 解析获取的POST body
			// req.ProtoMessage()
			proto.Unmarshal(body, req)
		} else {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.GetGroupDetailRsp)
		if ok {
			rsp.SessionId = req.SessionId
			mockRspBytes, _ = proto.Marshal(rsp)
		} else {
			panic("interface assert rsp failed. \n")
		}
	default:
		panic("请求的路径无对应Handle Function!!!\n")
	}

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
	url := fmt.Sprintf("%s/%s/Retrieve", api, common.ZeusSrv)
	resp, _ := http.Post(url, common.Content_type, body)
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

	url := fmt.Sprintf("%s/%s/AddFeas", api, common.ZeusSrv)
	resp, _ := http.Post(url, common.Content_type, body)

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
	url := fmt.Sprintf("%s/%s/CreateGroup", api, common.ZeusSrv)
	resp, _ := http.Post(url, common.Content_type, body)
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
