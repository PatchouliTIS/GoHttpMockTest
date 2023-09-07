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
	// Retrieve
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

	// Group
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
	case "DescribeGroup":
		req, ok := common.StructMap[reqName].(*yz.DescribeGroupReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.DescribeGroupRsp)
		if ok {
			rsp.SessionId = req.SessionId
			if *(req.Bitflag) == 6 {
				fmt.Println("Version And Size")
			}
			fmt.Printf("DescribeGroup BitFlag:%d\n", *(req.Bitflag))
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	case "UpdateGroup":
		req, ok := common.StructMap[reqName].(*yz.UpdateGroupReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.UpdateGroupRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	case "DeleteGroup":
		req, ok := common.StructMap[reqName].(*yz.DeleteGroupReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.DeleteGroupRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp

	// Features
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
	case "UpdateFeas":
		req, ok := common.StructMap[reqName].(*yz.UpdateFeasReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.UpdateFeasRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	case "DeleteFeas":
		req, ok := common.StructMap[reqName].(*yz.DeleteFeasReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.DeleteFeasRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp

	// Entity
	case "DescribeEntity":
		req, ok := common.StructMap[reqName].(*yz.DescribeEntityReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.DescribeEntityRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	case "UpdateEntity":
		req, ok := common.StructMap[reqName].(*yz.UpdateEntityReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.UpdateEntityRsp)
		if ok {
			rsp.SessionId = req.SessionId
		} else {
			panic("interface assert rsp failed. \n")
		}
		return req, rsp
	case "DeleteEntity":
		req, ok := common.StructMap[reqName].(*yz.DeleteEntityReq)
		if !ok {
			panic("interface assert req failed. \n")
		}

		rsp, ok := common.StructMap[rsqName].(*yz.DeleteEntityRsp)
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
