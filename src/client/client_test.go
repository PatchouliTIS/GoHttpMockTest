package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
	"zeus/src/common"

	"google.golang.org/protobuf/proto"
)

var (
	client = &http.Client{}
	config = &common.Config{}
)

func TestMain(m *testing.M) {
	// 本地服务器注册对应的 RESTful API 接口
	// http.HandleFunc("/ZeusService/CreateGroup", server.HandleRequst)
	// http.HandleFunc("/ZeusService/AddFeas", server.HandleRequst)
	// http.HandleFunc("/ZeusService/Retrieve", server.HandleRequst)
	// go http.ListenAndServe(":14000", nil)
	time.Sleep(1 * time.Second)

	// 初始化 Client 客户端
	config = &common.Config{}
	config.LoadConfig()
	client = NewClient()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestRetrieve(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"Retrieve": "114514",
	}

	for name, session := range API_name {
		req, expectResp, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
		}

		// 通过 ProtoReflect 动态获取结构体的字段
		respReflect := expectResp.ProtoReflect()
		sessionDescriber := respReflect.Descriptor().Fields().ByName("session_id")
		sessionInterface := respReflect.Get(sessionDescriber)
		sessionId, ok := sessionInterface.Interface().(string)
		if ok {
			fmt.Printf("Current Session Id : %v\n", sessionId)
		} else {
			panic(">>> protoreflect interface wrong <<<\n")
		}

		respBody, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		respBytes, err := io.ReadAll(respBody.Body)
		if err != nil {
			panic(err)
		}

		defer respBody.Body.Close()

		// TODO: 通过 ProtoReflect 获取不同内容
		resp, ok := common.StructMap[name+"Rsp"].(proto.Message)
		if ok {
			err = proto.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("ProtoUnmarshal got message for u : %s", string(respBytes))
				panic(err)
			}

			// 通过反射获取值
			reflect := resp.ProtoReflect()
			describer := reflect.Descriptor().Fields().ByName("session_id")
			value, ok := reflect.Get(describer).Interface().(string)
			if ok {
				t.Logf("deserialize data success: session_id -> %s", value)
				if value != session {
					t.Errorf("We Have Wrong Response, expect %s but get %s", session, value)
				} else {
					t.Logf("%s rpc call success!\n ------------------------------------", name)
				}
			} else {
				t.Errorf(">>> protoreflect interface convert failed <<<")
			}

		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestAddFeas(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"AddFeas": "114514",
	}

	for name, session := range API_name {
		req, expectResp, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
		}

		// 通过 ProtoReflect 动态获取结构体的字段
		respReflect := expectResp.ProtoReflect()
		sessionDescriber := respReflect.Descriptor().Fields().ByName("session_id")
		sessionInterface := respReflect.Get(sessionDescriber)
		sessionId, ok := sessionInterface.Interface().(string)
		if ok {
			fmt.Printf("Current Session Id : %v\n", sessionId)
		} else {
			panic(">>> protoreflect interface wrong <<<\n")
		}

		respBody, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		respBytes, err := io.ReadAll(respBody.Body)
		if err != nil {
			panic(err)
		}

		defer respBody.Body.Close()

		// TODO: 通过 ProtoReflect 获取不同内容
		resp, ok := common.StructMap[name+"Rsp"].(proto.Message)
		if ok {
			err = proto.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("ProtoUnmarshal got message for u : %s", string(respBytes))
				panic(err)
			}

			// 通过反射获取值
			reflect := resp.ProtoReflect()
			describer := reflect.Descriptor().Fields().ByName("session_id")
			value, ok := reflect.Get(describer).Interface().(string)
			if ok {
				t.Logf("deserialize data success: session_id -> %s", value)
				if value != session {
					t.Errorf("We Have Wrong Response, expect %s but get %s", session, value)
				} else {
					t.Logf("%s rpc call success!\n ------------------------------------", name)
				}
			} else {
				t.Errorf(">>> protoreflect interface convert failed <<<")
			}

		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}

}
