package bootstrap

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/hhxsv5/gin-x/config"
	_ "github.com/hhxsv5/gin-x/db/mysql"
	"github.com/hhxsv5/gin-x/log"
	"gitlab.kucoin.net/golang/skeleton/app/cron"
	//_ "github.com/hhxsv5/gin-x/db/postgresql"
)

func init() {
	go writePIdFile(config.BasePath + "/runtime/pid")

	if err := log.InitDailyLog("gin", config.BasePath+"/runtime/logs"); err != nil {
		panic(err)
	}

	if config.AppCron() {
		crons.Start()
	}
}

// 记录PID到文件，Reload时需要
func writePIdFile(path string) {
	p := strconv.Itoa(os.Getpid())
	b := bytes.NewBufferString(p).Bytes()
	err := ioutil.WriteFile(path, b, os.ModePerm)
	if err != nil {
		panic(errors.New(fmt.Sprintf("create pidfile fail: %s", err)))
	}
}
