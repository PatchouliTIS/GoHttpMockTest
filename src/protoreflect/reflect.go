package protoreflect

import (
	"fmt"
	"zeus/api/youtu_zeus"
	"zeus/src/common"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var father proto.Message

var son = common.StructMap["CreateGroupReq"]

func Reflect() (msg string, err error) {
	res, ok := son.(*youtu_zeus.CreateGroupReq)
	if ok {
		father = res
		/*
		  使用 ProtoReflect 获取 protoreflect.Message 对象
		  protoreflect.ProtoMessage -> protoreflect.Message
		*/
		messageReflect := father.ProtoReflect()

		// 遍历字段
		messageReflect.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			fmt.Printf("Field:\t%s\t Value:\t%v\n", fd.Name(), v.Interface())
			if fd.Name() == "session_id" {
				newValue := protoreflect.ValueOfString("1919810")
				messageReflect.Set(fd, newValue)
				// fmt.Printf("After >>>> Field:\t%s\t Value:\t%v\n", fd.Name(), v.Interface())
			}

			return true
		})

		fmt.Printf("After >>> \n %s", res)

		return
	} else {
		msg = ">>> interface converse failed <<<"
		err = fmt.Errorf(">>> interface converse failed <<<")
		return
	}
}
