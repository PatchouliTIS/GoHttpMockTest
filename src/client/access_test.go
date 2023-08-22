package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	yz "zeus/api/youtu_zeus"

	"github.com/golang/protobuf/proto"
)

func TestRetrieve(t *testing.T) {

	retrieveReq := yz.RetrieveReq{
		SessionId: &Session_id,
	}

	// 准备httptest
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body := make([]byte, r.ContentLength)

		r.Body.Read(body)
		req := yz.RetrieveReq{}
		proto.Unmarshal(body, &req)

		rsp := yz.RetrieveRsp{
			SessionId: req.SessionId,
			Errorcode: &Error_code,
			Errormsg:  &Error_msg,
		}
		mockRspBytes, _ := proto.Marshal(&rsp)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.Write(mockRspBytes)

		fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\n",
			r.Method, r.URL.Path, r.URL.EscapedPath())
	}))

	defer ts.Close()

	api := ts.URL
	fmt.Println("API route:", api)

	resp, err := Retrieve(api, &retrieveReq)
	if err != nil {
		t.Fatalf("Something wrong in Retrieve:%s", err)
	} else {
		t.Logf(">>>>>Retrieve Done: %v", resp)
	}

}

func TestAddFeas(t *testing.T) {
	addFeasReq := yz.AddFeasReq{
		SessionId: &Session_id,
	}

	// 准备httptest
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		req := yz.AddFeasReq{}
		proto.Unmarshal(body, &req)

		rsp := yz.AddFeasRsp{
			SessionId: req.SessionId,
			Errorcode: &Error_code,
			Errormsg:  &Error_msg,
		}
		mockRspBytes, _ := proto.Marshal(&rsp)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.Write(mockRspBytes)

		fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\n",
			r.Method, r.URL.Path, r.URL.EscapedPath())
	}))

	defer ts.Close()

	api := ts.URL
	fmt.Println("API route:", api)

	resp, err := AddFeas(api, &addFeasReq)
	if err != nil {
		t.Fatalf("Something wrong in AddFeas:%s", err)
	} else {
		t.Logf(">>>>>>AddFeas Done: %v", resp)
	}
}

func TestCreateGroup(t *testing.T) {
	createGroupReq := yz.CreateGroupReq{
		SessionId:     &Session_id,
		GroupId:       &Group_id,
		Platform:      &CPU_Platform,
		FeatureConfig: &Feature,
	}

	// 准备httptest
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		req := yz.CreateGroupReq{}
		proto.Unmarshal(body, &req)

		rsp := yz.CreateGroupRsp{
			SessionId: req.SessionId,
			Errorcode: &Error_code,
			Errormsg:  &Error_msg,
		}
		mockRspBytes, _ := proto.Marshal(&rsp)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/x-protobuf")
		w.Write(mockRspBytes)

		fmt.Printf("Capture Request Method: %s \t Request Path: %s\nRequest Escape Path:%s\nNested Struct:%s",
			r.Method, r.URL.Path, r.URL.EscapedPath(), &req)
	}))

	defer ts.Close()

	api := ts.URL
	fmt.Println("API route:", api)

	resp, err := CreateGroup(api, &createGroupReq)
	if err != nil {
		t.Fatalf("Something wrong in CreateGroup:%s", err)
	} else {
		t.Logf(">>>>>>CreateGroup Done: %v", resp)
	}
}
