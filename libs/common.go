/*
 * Copyright 2017 DIST Author. All Rights Reserved.
 * DIST Agent Libs and Global values
 * 2017/7/25, by Tan Haipeng, create
 */

package libs

import "net/http"
import (
	"github.com/go-ozzo/ozzo-config"
	"github.com/go-ozzo/ozzo-log"
	"io"
)

func InitConfig(path string) *config.Config {
	conf := config.New()
	conf.Load(path)
	return conf
}

func InitLogger(path string) *log.Logger {
	logger := log.NewLogger()
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = path
	t2.MaxLevel = log.LevelDebug
	logger.Targets = append(logger.Targets, t1, t2)
	logger.Open()
	return logger
}

func GetRequest(req *http.Request, rtype string) map[string]string {
	var params = make(map[string]string)
	if rtype == "get" {
		query := req.URL.Query()
		for k, v := range query {
			if len(v) > 0 {
				params[k] = v[0]
			}
		}
	}
	if rtype == "post" {

	}
	return params
}

func SendResponse(rsp http.ResponseWriter, rtype string) {
	if rtype == "json" {
		io.WriteString(rsp, CombineRetData(0, "success"))
	} else {
		io.WriteString(rsp, CombineRetData(100, "type error"))
	}
}

func CombineRetData(code int, msg string) string {

	return ""
}

func CallSlaveApi(url string, params map[string]string) bool {
	return true
}