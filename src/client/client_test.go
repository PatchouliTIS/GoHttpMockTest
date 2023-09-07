package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	yz "zeus/api/youtu_zeus"
	"zeus/src/common"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	client = &http.Client{}
	config = &common.Config{}
)

func TestMain(m *testing.M) {
	// 本地服务器注册对应的 RESTful API 接口
	// http.HandleFunc("/ZeusService/CreateGroup", server.HandleRequstJSON)
	// http.HandleFunc("/ZeusService/AddFeas", server.HandleRequstJSON)
	// http.HandleFunc("/ZeusService/Retrieve", server.HandleRequstJSON)
	// http.HandleFunc("/ZeusService/DeleteGroup", server.HandleRequstJSON)
	// http.HandleFunc("/ZeusService/DescribeGroup", server.HandleRequstJSON)
	// go http.ListenAndServe(":14000", nil)
	// time.Sleep(1 * time.Second)

	// 初始化 Client 客户端
	config = &common.Config{}
	config.LoadConfig()
	client = NewClient()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateGroup(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"CreateGroup": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
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
		resp, ok := common.StructMap[name+"Rsp"].(*yz.CreateGroupRsp)
		if ok {
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got error message for u : %s", string(respBytes))
				panic(err)
			}
			printRespInfo(resp)
		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestDescribeGroup(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"DescribeGroup": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
		}

		// 通过 ProtoReflect 动态获取结构体的字段

		// 发起请求
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
		resp, ok := common.StructMap[name+"Rsp"].(*yz.DescribeGroupRsp)
		if ok {
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got error message for u : %s", string(respBytes))
				panic(err)
			}
			printRespInfo(resp)
			if resp.GroupConfig != nil {
				fmt.Printf("GroupConfig: %v\n", resp.GroupConfig)
			}
		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestUpdateGroup(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"UpdateGroup": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
		}

		// 通过 ProtoReflect 动态获取结构体的字段

		// 发起请求
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
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got error message for u : %s", string(respBytes))
				panic(err)
			}
			printRespInfo(resp)
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

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
		}

		// // 通过 ProtoReflect 动态获取结构体的字段
		// respReflect := expectResp.ProtoReflect()
		// sessionDescriber := respReflect.Descriptor().Fields().ByName("session_id")
		// sessionInterface := respReflect.Get(sessionDescriber)
		// sessionId, ok := sessionInterface.Interface().(string)
		// if ok {
		// 	fmt.Printf("Current Session Id : %v\n", sessionId)
		// } else {
		// 	panic(">>> protoreflect interface wrong <<<\n")
		// }

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
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got message for u : %s", string(respBytes))
				panic(err)
			}

			printRespInfo(resp)

		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestGetFeas(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"GetFeas": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
		}

		// // 通过 ProtoReflect 动态获取结构体的字段
		// respReflect := expectResp.ProtoReflect()
		// sessionDescriber := respReflect.Descriptor().Fields().ByName("session_id")
		// sessionInterface := respReflect.Get(sessionDescriber)
		// sessionId, ok := sessionInterface.Interface().(string)
		// if ok {
		// 	fmt.Printf("Current Session Id : %v\n", sessionId)
		// } else {
		// 	panic(">>> protoreflect interface wrong <<<\n")
		// }

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
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got message for u : %s", string(respBytes))
				panic(err)
			}

			printRespInfo(resp)

		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestUpdateFeas(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"UpdateFeas": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
		}

		// // 通过 ProtoReflect 动态获取结构体的字段
		// respReflect := expectResp.ProtoReflect()
		// sessionDescriber := respReflect.Descriptor().Fields().ByName("session_id")
		// sessionInterface := respReflect.Get(sessionDescriber)
		// sessionId, ok := sessionInterface.Interface().(string)
		// if ok {
		// 	fmt.Printf("Current Session Id : %v\n", sessionId)
		// } else {
		// 	panic(">>> protoreflect interface wrong <<<\n")
		// }

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
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got message for u : %s", string(respBytes))
				panic(err)
			}

			printRespInfo(resp)

		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestDescribeEntity(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"DescribeEntity": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
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
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got error message for u : %s", string(respBytes))
				panic(err)
			}
			printRespInfo(resp)
		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestUpdateEntity(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"UpdateEntity": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
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
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got error message for u : %s", string(respBytes))
				panic(err)
			}
			printRespInfo(resp)
		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestRetrieve(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"Retrieve": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
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
		resp, ok := common.StructMap[name+"Rsp"].(*yz.RetrieveRsp)
		if ok {
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got error message for u : %s", string(respBytes))
				panic(err)
			}
			printRespInfo(resp)
			if resp.GroupItems != nil {
				for _, item := range resp.GroupItems {
					fmt.Printf("%v\n", item)
				}
			}
		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestDeleteEntity(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"DeleteEntity": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
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
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got error message for u : %s", string(respBytes))
				panic(err)
			}
			printRespInfo(resp)
		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestDeleteFeas(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"DeleteFeas": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
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
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got error message for u : %s", string(respBytes))
				panic(err)
			}
			printRespInfo(resp)
		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

func TestDeleteGroup(t *testing.T) {
	// 发出请求，接收数据
	API_name := map[string]string{
		"DeleteGroup": "114514",
	}

	for name := range API_name {
		req, _, err := NewHTTPRequest(config, name)
		if err != nil {
			panic(err)
		}

		// 通过 ProtoReflect 动态获取结构体的字段

		// 发起请求
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
			err = json.Unmarshal(respBytes, resp)
			if err != nil {
				fmt.Printf("JSONUnmarshal got error message for u : %s", string(respBytes))
				panic(err)
			}
			printRespInfo(resp)
		} else {
			t.Errorf("deserialize data failed: %sRsp", name)
		}
	}
}

// 通过反射获取错误码
func getErrorCode(resp proto.Message, name string, t *testing.T) (value int32) {
	reflect := resp.ProtoReflect()
	describer := reflect.Descriptor().Fields().ByName("error_code")
	value, ok := reflect.Get(describer).Interface().(int32)
	if ok {
		t.Logf("deserialize data success: error_code -> %d\n", value)
		if value != 0 {
			describer := reflect.Descriptor().Fields().ByName("error_msg")
			msg, ok := reflect.Get(describer).Interface().(string)
			if ok {
				t.Errorf("We Have Error Code:%d\tMsg:%s\n", value, msg)
			} else {
				t.Errorf(">>> Error Message read failed <<<\n")
			}
		} else {
			t.Logf("%s rpc call success!\n ------------------------------------\n", name)
		}
	} else {
		t.Errorf(">>> protoreflect interface convert failed <<<\n")
	}
	return
}

func printRespInfo(resp proto.Message) {
	reflect := resp.ProtoReflect()
	reflect.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		fmt.Printf("Field Name: %s, Field Type: %s, Field Value: ", fd.Name(), fd.Kind())
		// 遍历数组元素
		if fd.IsList() {
			list := v.List()
			fmt.Printf("[ ")
			for i := 0; i < list.Len(); i++ {
				fmt.Printf("%v, ", list.Get(i).Interface())
			}
			fmt.Printf("]\n")
		} else {
			fmt.Printf("%v\n", v.Interface())
		}
		// fmt.Printf("Field Name: %s, Field Type: %s, Field Value: %v\n", fd.Name(), fd.Kind(), v.Interface())
		return true
	})

	fmt.Println("----------------------------------")

}
