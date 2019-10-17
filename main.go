package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/selinplus/data-encrption/models"
	"github.com/selinplus/data-encrption/pkg/gredis"
	"github.com/selinplus/data-encrption/pkg/logging"
	"github.com/selinplus/data-encrption/pkg/setting"
	"github.com/selinplus/data-encrption/pkg/util"
	"github.com/selinplus/data-encrption/routers"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)
	log.Printf("[info] start http server listening %s", endPoint)

	err = server.ListenAndServe()
	if err != nil {
		log.Printf("init listen server fail:%v", err)
	}
}
