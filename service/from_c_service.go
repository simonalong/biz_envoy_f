package service

import (
	"biz-f-service/biz_config"
	"biz-f-service/pojo/domain"
	"context"
	"gorm.io/gorm/clause"
	"time"

	"github.com/isyscore/isc-gobase/logger"

	"github.com/gin-gonic/gin"
	"github.com/isyscore/isc-gobase/isc"
	"github.com/isyscore/isc-gobase/server/rsp"
	baseTime "github.com/isyscore/isc-gobase/time"
)

func Ef(c *gin.Context) {
	//logger.Info("调用 Ef 接口")

	d := c.Param("haveMysql")
	haveMysql := isc.ToBool(d)
	if haveMysql {
		biz_config.Db.Save(
			domain.BizEnvoyF{
				ServiceName: "biz-f-service",
				Times:       baseTime.TimeToStringYmdHms(time.Now()),
			},
		).Where("service_name=?", "biz-f-service")
	}
	rsp.SuccessOfStandard(c, "ok")
}

func Cf(c *gin.Context) {
	logger.Info("调用 cf 接口")

	haveRedis := isc.ToBool(c.Param("haveRedis"))
	haveMysql := isc.ToBool(c.Param("haveMysql"))
	haveEtcd := isc.ToBool(c.Param("haveEtcd"))
	logger.Info("haveRedis:%v, haveMysql:%v, haveEtcd:%v", haveRedis, haveMysql, haveEtcd)
	if haveRedis {
		rdb := biz_config.RedisDb
		ctx := context.Background()
		rs := rdb.Set(ctx, "biz-f-service.redis.key", baseTime.TimeToStringYmdHms(time.Now()), time.Hour*1)
		if rs.Err() != nil {
			logger.Error("调用redis设置失败：%v", rs.Err().Error())
		} else {
			logger.Info("调用redis设置成功")
		}
	}

	if haveMysql {
		biz_config.Db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&domain.BizEnvoyF{ServiceName: "biz-f-service", Times: baseTime.TimeToStringYmdHms(time.Now())})
	}

	if haveEtcd {
		ctx := context.Background()
		biz_config.EtcdClient.Put(ctx, "biz-f-service.etcd.key", baseTime.TimeToStringYmdHms(time.Now()))
	}

	rsp.SuccessOfStandard(c, "ok")
}

func CfOfEtcd(c *gin.Context) {
	//logger.Info("调用 Ef 接口")

	d := c.Param("haveEtcd")
	value := c.Param("value")
	haveEtcd := isc.ToBool(d)
	logger.Info("haveEtcd 字段为：%v", haveEtcd)
	if haveEtcd {
		ctx := context.Background()
		biz_config.EtcdClient.Put(ctx, "biz-f-service.key", value)
	}
	rsp.SuccessOfStandard(c, "ok")
}

func CfOkOk(c *gin.Context) {
	//logger.Info("调用 CfOkOk 接口")
	rsp.SuccessOfStandard(c, "ok")
}

func CfOkStop(c *gin.Context) {
	//logger.Info("调用 CfOkStop 接口")
	rsp.SuccessOfStandard(c, "ok")
}

func CfStopOk(c *gin.Context) {
	//logger.Info("调用 CfStopOk 接口")
	rsp.SuccessOfStandard(c, "ok")
}

func CfStopStop(c *gin.Context) {
	//logger.Info("调用 CfStopStop 接口")
	rsp.SuccessOfStandard(c, "ok")
}

func CfTrace(c *gin.Context) {
	//logger.Info("调用 CfTrace 接口")
	rsp.SuccessOfStandard(c, "ok")
}
