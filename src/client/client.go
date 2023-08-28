package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"zeus/src/common"
	"zeus/src/util"

	"google.golang.org/protobuf/proto"
)

func NewClient() *http.Client {
	return &http.Client{Transport: common.Transport}
}

func NewHTTPRequest(config *common.Config, rpcName string) (request *http.Request, respBody proto.Message, err error) {

	var body io.Reader

	if config.Method == "POST" || config.Method == "PUT" {
		// body = bytes.NewReader(config.bodyContent)
		// 根据传入的调用名读取 HTTP Request
		var reqBody proto.Message
		reqBody, respBody = util.GetReqAndRsp(rpcName)
		reqBytes, err := proto.Marshal(reqBody)
		if err != nil {
			panic(err)
		}
		body = bytes.NewReader(reqBytes)
	}
	//
	route := string(config.Host + ":" + config.Port + "/" + common.ZeusSrv + "/" + rpcName)
	// route := string("/" + common.ZeusSrv + "/" + rpcName)
	fmt.Println(route)
	request, err = http.NewRequest(config.Method, route, body)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", config.ContentType)
	// request.Header.Set("User-Agent", config.userAgent)

	// if config.keepAlive {
	request.Header.Set("Connection", "keep-alive")
	// }

	for _, header := range config.Headers {
		pair := strings.Split(header, ":")
		request.Header.Add(pair[0], pair[1])
	}

	for _, cookie := range config.Cookies {
		pair := strings.Split(cookie, "=")
		c := &http.Cookie{Name: pair[0], Value: pair[1]}
		request.AddCookie(c)
	}

	if config.BasicAuthentication != "" {
		pair := strings.Split(config.BasicAuthentication, ":")
		request.SetBasicAuth(pair[0], pair[1])
	}

	return
}
