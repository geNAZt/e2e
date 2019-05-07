package window

import (
	"github.com/BurntSushi/xgbutil"
	"gomint.io/e2e/test/logger"
)

var xgbUtil *xgbutil.XUtil

func GetConnection() *xgbutil.XUtil {
	EnsureConnection()
	return xgbUtil
}

func EnsureConnection() {
	if xgbUtil != nil {
		return
	}
	X, err := xgbutil.NewConn()
	if err != nil {
		logger.Fatal("Failed to create new XGB connection: %s", err.Error())
	}
	xgbUtil = X
}
