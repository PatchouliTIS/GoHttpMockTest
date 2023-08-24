package main

import (
	"flag"
	"net/http"
	"strings"
	"zeus/src/server"
	"zeus/src/util"
)

func main() {

	flag.StringVar(&server.Rpc_name, "rpc", "", "name of rpc called")
	flag.Parse()
	if server.Rpc_name == "" || len(server.Rpc_name) == 0 {
		panic("called rpc name wrong! Not exists")
	}
	server.Rpc_name = strings.TrimSpace(server.Rpc_name)
	names := strings.Split(server.Rpc_name, ",")

	// 使用 flag 生成 POST body .txt 文件，以供gohttpbench调用
	for _, name := range names {
		util.Body2Txt(name)
	}

	// 本地服务器注册对应的 RESTful API 接口
	http.HandleFunc("/ZeusService/CreateGroup", server.HandleRequst)
	http.HandleFunc("/ZeusService/AddFeas", server.HandleRequst)
	http.HandleFunc("/ZeusService/TruncateGroup", server.HandleRequst)
	http.HandleFunc("/ZeusService/GetGroupDetail", server.HandleRequst)
	http.ListenAndServe(":14000", nil)
}
