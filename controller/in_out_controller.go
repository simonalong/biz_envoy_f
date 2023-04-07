package controller

import (
	"biz-f-service/service"

	"github.com/isyscore/isc-gobase/server"
)

func CfController() {
	// 测试入口和出口的拦截情况
	server.Get("ef/:haveMysql", service.Ef)
	server.Get("cf/:haveRedis/:haveMysql/:haveEtcd", service.Cf)
	// server.Get("cfOfEtcd/:haveEtcd/:value", service.CfOfEtcd)

	server.Get("cf/ok/ok", service.CfOkOk)
	server.Get("cf/ok/stop", service.CfOkStop)
	server.Get("cf/stop/ok", service.CfStopOk)
	server.Get("cf/stop/stop", service.CfStopStop)

	// 测试：链路
	server.Get("cf/trace", service.CfTrace)
}
