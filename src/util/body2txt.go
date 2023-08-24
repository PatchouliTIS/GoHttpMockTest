package util

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

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
	req, _ := GetReqAndRsp(rpc)
	data, err := proto.Marshal(req)
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
