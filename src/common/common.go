package common

import (
	"crypto/tls"
	"net/http"
	"time"
	"zeus/api/youtu"
	yz "zeus/api/youtu_zeus"
)

const (
	Content_type = "application/x-protobuf"
	ZeusSrv      = "ZeusService"
)

var (
	Error_code       int32             = 0
	Session_id       string            = "114514"
	Group_id         string            = "75th Ranger Regiment"
	Error_msg        string            = ""
	CPU_Platform     youtu.Platform    = 0
	GPU_Platform     youtu.Platform    = 1
	Dimension        int32             = 5
	FeatureType_INT8 youtu.FeatureType = 1
	FeatureIdx       int32             = 0
	Scale            float64           = 1.4038
	FeaId            string            = "YJSP"
	EntityId         string            = "Koumakan"
	Feature                            = yz.FeatureConfig{
		Dimension:   &Dimension,
		FeatureType: &(FeatureType_INT8),
		Scale:       &Scale,
		FeatureIdx:  &FeatureIdx,
	}

	// StructMap = map[string]proto.Message{
	StructMap = map[string]interface{}{
		"AddFeasReq":       &yz.AddFeasReq{SessionId: &Session_id},
		"AddFeasRsp":       &yz.AddFeasRsp{SessionId: &Session_id},
		"RetrieveReq":      &yz.RetrieveReq{SessionId: &Session_id},
		"RetrieveRsp":      &yz.RetrieveRsp{SessionId: &Session_id},
		"TruncateGroupReq": &yz.TruncateGroupReq{SessionId: &Session_id},
		"TruncateGroupRsp": &yz.TruncateGroupRsp{},
		"CreateGroupReq": &yz.CreateGroupReq{
			SessionId:     &Session_id,
			GroupId:       &Group_id,
			Platform:      &CPU_Platform,
			FeatureConfig: &Feature,
		},
		"CreateGroupRsp": &yz.CreateGroupRsp{},
		"GetGroupDetailReq": &yz.GetGroupDetailReq{
			SessionId:  &Session_id,
			GroupId:    &Group_id,
			FeatureIdx: &FeatureIdx,
		},
		"GetGroupDetailRsp": &yz.GetGroupDetailRsp{},
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
	ContentType         string
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
	c.Port = "14000"
	c.ContentType = "application/x-protobuf"

	return nil
}
