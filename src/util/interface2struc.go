package util

import (
	yz "zeus/api/youtu_zeus"
	"zeus/src/client"

	"google.golang.org/protobuf/proto"
)

func GetReqAndRsp(name string) (proto.Message, proto.Message) {

	reqName := name + "Req"
	rsqName := name + "Rsp"

	switch name {
	case "Retrieve":
		req, ok := client.StructMap[reqName].(*yz.RetrieveReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := client.StructMap[rsqName].(*yz.RetrieveRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	case "AddFeas":
		req, ok := client.StructMap[reqName].(*yz.AddFeasReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := client.StructMap[rsqName].(*yz.AddFeasRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	case "TruncateGroup":
		req, ok := client.StructMap[reqName].(*yz.TruncateGroupReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := client.StructMap[rsqName].(*yz.TruncateGroupRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	case "CreateGroup":
		req, ok := client.StructMap[reqName].(*yz.CreateGroupReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := client.StructMap[rsqName].(*yz.CreateGroupRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	case "GetGroupDetail":
		req, ok := client.StructMap[reqName].(*yz.GetGroupDetailReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := client.StructMap[rsqName].(*yz.GetGroupDetailRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	default:
		panic("请求的路径无对应Handle Function!!!\n")
	}
}
