package common

import (
	"crypto/tls"
	"net/http"
	"time"
	"zeus/api/youtu"
	yz "zeus/api/youtu_zeus"
)

const (
	ContentJSONType = "application/x-protobuf"
	ZeusSrv         = "ZeusService"
	GET_SIZE        = "BITFLAG_GET_GROUP_SIZE"
	GET_CONFIG      = "BITFLAG_GET_GROUP_CONFIG"
	GET_VERSION     = "BITFLAG_GET_GROUP_VERSION"
)

var (
	Error_code   int32          = 0
	Session_id   string         = "1919813"
	Group_id     string         = "patchychen_zeus_function_test"
	Error_msg    string         = ""
	TopN         uint32         = 10
	Threshold    float32        = 0.74
	CPU_Platform youtu.Platform = 0
	GPU_Platform youtu.Platform = 1
	Dimension0   int32          = int32(len(FeaFace0Bytes))
	Dimension1   int32          = int32(len(FeaFace1Bytes))
	UpdDi0       int32          = int32(len(UpdateFeaFace0Bytes))
	UpdDi1       int32          = int32(len(UpdateFeaFace1Bytes))
	TRUE         bool           = true
	FALSE        bool           = false

	DescGroupConfig    yz.BitFlag = yz.BitFlag(yz.BitFlag_value[GET_CONFIG])
	DescGroupSize      yz.BitFlag = yz.BitFlag(yz.BitFlag_value[GET_SIZE])
	DescGroupVersion   yz.BitFlag = yz.BitFlag(yz.BitFlag_value[GET_VERSION])
	DescVersionAndSize yz.BitFlag = (DescGroupSize | DescGroupVersion)

	ADD_FEATURE_CONFIG    yz.UpdateGroupReq_UpdateGroupType = 1
	DELETE_FEATURE_CONFIG yz.UpdateGroupReq_UpdateGroupType = 2

	Feature0Idx         int32   = 0 // 0->feature0   1->feature1   2->双特征
	Feature1Idx         int32   = 1
	FeaFace0Bytes       []byte  = []byte("this is face_fea_0")
	FeaFace1Bytes       []byte  = []byte("this is face_fea_1")
	UpdateFeaFace0Bytes []byte  = []byte("updated face_fea_0")
	UpdateFeaFace1Bytes []byte  = []byte("updated face_fea_1")
	Scale               float64 = 1.4038

	Capacity         int32             = 2048
	FeatureType_INT8 youtu.FeatureType = youtu.FeatureType_INT8
	FeaFace0Id       string            = "fea_face_0"
	FeaFace1Id       string            = "fea_face_1"
	FeaIds                             = []string{"fea_face_0", "fea_face_1"}
	EntityId         string            = "entity_koumakan"
	UpdateEntityId   string            = "entity_ranger"

	Feature = yz.FeatureConfig{
		Dimension:   &Dimension0,
		FeatureType: &(FeatureType_INT8),
		Scale:       &Scale,
		FeatureIdx:  &Feature0Idx,
	}
	UpdateFeature = yz.FeatureConfig{
		Dimension:   &Dimension0,
		FeatureType: &(FeatureType_INT8),
		Scale:       &Scale,
		FeatureIdx:  &Feature1Idx,
	}
	FeaItem = []*yz.FeaItem{
		{
			FeaId:      &FeaFace0Id,
			EntityId:   &EntityId,
			Feature_0:  FeaFace0Bytes,
			Feature_1:  FeaFace1Bytes,
			ExtendInfo: FeaFace0Bytes,
		},
		{
			FeaId:      &FeaFace1Id,
			EntityId:   &EntityId,
			Feature_0:  FeaFace0Bytes,
			Feature_1:  FeaFace1Bytes,
			ExtendInfo: FeaFace1Bytes,
		},
		// {
		// 	FeaId:      &FeaChestId,
		// 	EntityId:   &EntityId,
		// 	Feature_0:  FeaChestBytes,
		// 	Feature_1:  FeaChestBytes,
		// 	ExtendInfo: FeaChestBytes,
		// },
	}

	UpdateFeaItem = []*yz.FeaItem{
		{
			FeaId:      &FeaFace0Id,
			EntityId:   &EntityId,
			Feature_0:  UpdateFeaFace0Bytes,
			Feature_1:  UpdateFeaFace1Bytes,
			ExtendInfo: UpdateFeaFace0Bytes,
		},
		{
			FeaId:      &FeaFace1Id,
			EntityId:   &EntityId,
			Feature_0:  UpdateFeaFace0Bytes,
			Feature_1:  UpdateFeaFace1Bytes,
			ExtendInfo: UpdateFeaFace1Bytes,
		},
		// {
		// 	FeaId:      &FeaChestId,
		// 	EntityId:   &EntityId,
		// 	Feature_0:  UpdateFeaChestBytes,
		// 	Feature_1:  UpdateFeaChestBytes,
		// 	ExtendInfo: UpdateFeaChestBytes,
		// },
	}

	// StructMap = map[string]proto.Message{
	StructMap = map[string]interface{}{
		// Features
		"AddFeasReq": &yz.AddFeasReq{
			SessionId:  &Session_id,
			GroupId:    &Group_id,
			Items:      FeaItem,
			Force:      &FALSE,
			FeatureIdx: &Feature0Idx,
		},
		"AddFeasRsp": &yz.AddFeasRsp{
			FeaIds: FeaIds,
		},
		"GetFeasReq": &yz.GetFeasReq{
			SessionId:  &Session_id,
			GroupId:    &Group_id,
			FeaIds:     FeaIds,
			FeatureIdx: &Feature0Idx,
		},
		"GetFeasRsp": &yz.GetFeasRsp{},
		"UpdateFeasReq": &yz.UpdateFeasReq{
			SessionId:  &Session_id,
			GroupId:    &Group_id,
			Items:      UpdateFeaItem,
			FeatureIdx: &Feature0Idx,
		},
		"UpdateFeasRsp": &yz.UpdateFeasRsp{},
		"DeleteFeasReq": &yz.DeleteFeasReq{
			SessionId: &Session_id,
			GroupId:   &Group_id,
			FeaIds:    FeaIds,
		},
		"DeleteFeasRsp": &yz.DeleteFeasRsp{},

		// Group
		"CreateGroupReq": &yz.CreateGroupReq{
			SessionId:     &Session_id,
			GroupId:       &Group_id,
			Capacity:      &Capacity,
			Platform:      &CPU_Platform,
			FeatureConfig: &Feature,
		},
		"CreateGroupRsp": &yz.CreateGroupRsp{},
		"DescribeGroupReq": &yz.DescribeGroupReq{
			SessionId: &Session_id,
			GroupId:   &Group_id,
			Bitflag:   &DescVersionAndSize,
		},
		"DescribeGroupRsp": &yz.DescribeGroupRsp{},
		"UpdateGroupReq": &yz.UpdateGroupReq{
			SessionId:       &Session_id,
			GroupId:         &Group_id,
			UpdateGroupType: &ADD_FEATURE_CONFIG,
			FeatureConfig:   &UpdateFeature,
		},
		"UpdateGroupRsp": &yz.UpdateGroupRsp{},
		"DeleteGroupReq": &yz.DeleteGroupReq{
			SessionId: &Session_id,
			GroupId:   &Group_id,
		},
		"DeleteGroupRsp": &yz.DeleteGroupRsp{},

		// Entity
		// 更新什么？
		// 更新指定的FeaIds的EntityId
		"UpdateEntityReq": &yz.UpdateEntityReq{
			SessionId: &Session_id,
			GroupId:   &Group_id,
			FeaIds:    FeaIds,
			EntityId:  &UpdateEntityId,
		},
		"UpdateEntityRsp": &yz.UpdateEntityRsp{},
		"DescribeEntityReq": &yz.DescribeEntityReq{
			SessionId:  &Session_id,
			GroupId:    &Group_id,
			FeatureIdx: &Feature0Idx,
			EntityId:   &EntityId,
		},
		"DescribeEntityRsp": &yz.DescribeEntityRsp{},
		"DeleteEntityReq": &yz.DeleteEntityReq{
			SessionId: &Session_id,
			GroupId:   &Group_id,
			EntityId:  &EntityId,
		},
		"DeleteEntityRsp": &yz.DeleteEntityRsp{},

		// Retrieve
		"RetrieveReq": &yz.RetrieveReq{
			SessionId: &Session_id,
			GroupIds:  []string{Group_id},
			Features: [][]byte{
				FeaFace0Bytes,
				FeaFace1Bytes,
			},
			Topn:        &TopN,
			Threshold:   &Threshold,
			FeatureIdxs: []int32{Feature0Idx},
		},
		"RetrieveRsp": &yz.RetrieveRsp{},
	}

	// skip certification check for self-signed certificates
	Tlsconfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// Custom Transport
	Transport = &http.Transport{
		DisableCompression: true,
		DisableKeepAlives:  false,
		TLSClientConfig:    Tlsconfig,
	}
)

type Config struct {
	Requests         int // 全部HTTP请求数
	Concurrency      int // 并发用户数
	Timelimit        int
	ExecutionTimeout time.Duration

	Method              string
	BodyContent         []byte
	ContentProtoType    string
	ContentJSONType     string
	Headers             []string
	Cookies             []string
	Gzip                bool
	KeepAlive           bool
	BasicAuthentication string
	// userAgent           string

	// Url string
	Host string
	Port string
}

func (c *Config) LoadConfig() error {

	c.Requests = 1
	c.Concurrency = 1
	c.Method = http.MethodPost
	c.Gzip = true
	c.KeepAlive = true
	c.Host = "http://127.0.0.1"
	c.Port = "12100" // zeus_access12100
	c.ContentProtoType = "application/x-protobuf"
	c.ContentJSONType = "application/json"

	return nil
}
