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
