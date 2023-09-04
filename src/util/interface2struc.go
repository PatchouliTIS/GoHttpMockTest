package util

import (
	"fmt"
	yz "zeus/api/youtu_zeus"
	"zeus/src/common"

	"google.golang.org/protobuf/proto"
)

func GetReqAndRsp(name string) (proto.Message, proto.Message) {

	reqName := name + "Req"
	rsqName := name + "Rsp"

	// 无法通过接口类型调用实现了该接口的实体对象的自有方法！
	// var req proto.Message
	// var rsp proto.Message
	// var ok bool

	switch name {
	case "Retrieve":
		req, ok := common.StructMap[reqName].(*yz.RetrieveReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.RetrieveRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	case "AddFeas":
		req, ok := common.StructMap[reqName].(*yz.AddFeasReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.AddFeasRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp

	case "CreateGroup":
		req, ok := common.StructMap[reqName].(*yz.CreateGroupReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.CreateGroupRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp

	case "GetFeas":
		req, ok := common.StructMap[reqName].(*yz.GetFeasReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.GetFeasRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	default:
		panic(fmt.Sprintf("请求的路径 %s 无对应Handle Function!!!\n", name))
	}
}
